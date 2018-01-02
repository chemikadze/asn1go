package asn1go

type AstNode interface{}

type ModuleDefinition struct {
	ModuleIdentifier     ModuleIdentifier
	TagDefault           int
	ExtensibilityImplied bool
	ModuleBody           ModuleBody
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Header information

type ModuleIdentifier struct {
	Reference            string
	DefinitiveIdentifier DefinitiveIdentifier
}

type DefinitiveIdentifier []DefinitiveObjIdComponent

type DefinitiveObjIdComponent struct {
	Name string
	Id   int
}

const (
	TAGS_EXPLICIT = iota
	TAGS_IMPLICIT
	TAGS_AUTOMATIC
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Module body

type ModuleBody struct {
	AssignmentList AssignmentList
	Imports        []SymbolsFromModule
}

type SymbolsFromModule struct {
	SymbolList []Symbol
	Module     GlobalModuleReference
}

type Symbol interface {
	IsSymbol()
}

type GlobalModuleReference struct {
	Reference          string
	AssignedIdentifier Value
}

type AssignmentList []Assignment

func NewAssignmentList(other ...Assignment) AssignmentList {
	return make(AssignmentList, 0).Append(other...)
}

func (al AssignmentList) Append(other ...Assignment) AssignmentList {
	return append(al, other...)
}

func (l AssignmentList) Get(name string) Assignment {
	for _, assignment := range l {
		if assignment.Reference().Name() == name {
			return assignment
		}
	}
	return nil
}

func (l AssignmentList) GetValue(name string) *ValueAssignment {
	a := l.Get(name)
	if a == nil {
		return nil
	}
	switch r := a.(type) {
	case ValueAssignment:
		return &r
	default:
		return nil
	}
}

func (l AssignmentList) GetType(name string) *TypeAssignment {
	a := l.Get(name)
	if a == nil {
		return nil
	}
	switch r := a.(type) {
	case TypeAssignment:
		return &r
	default:
		return nil
	}
}

type Assignment interface {
	Reference() Reference
}

// assigns Value of Type to ValueReference
type ValueAssignment struct {
	ValueReference ValueReference
	Type           Type
	Value          Value
}

func (v ValueAssignment) Reference() Reference {
	return v.ValueReference
}

// assigns Type to TypeReference
type TypeAssignment struct {
	TypeReference TypeReference
	Type          Type
}

func (v TypeAssignment) Reference() Reference {
	return v.TypeReference
}

type NamedType struct {
	Identifier Identifier
	Type       Type
}

func (t NamedType) Zero() interface{} {
	return t.Type.Zero()
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References

type Reference interface {
	Name() string
}

// type reference
type TypeReference string

func (r TypeReference) Name() string {
	return string(r)
}

func (r TypeReference) Zero() interface{} {
	return nil
}

func (TypeReference) IsSymbol() {}

// value reference
type ValueReference string

func (r ValueReference) Name() string {
	return string(r)
}

func (ValueReference) IsSymbol() {}

// module reference
type ModuleReference string

func (ModuleReference) IsSymbol() {}

// identifier type
type Identifier string

func (id Identifier) Name() string {
	return string(id)
}

// number lexem, implements Value
type Number int

func (x Number) IntValue() int {
	return int(x)
}

func (Number) Type() Type {
	return IntegerType{}
}

func (x Number) UnaryMinus() Number {
	return Number(-int(x))
}

// real lexem
type Real float64

func (x Real) Type() Type {
	return RealType{}
}

func (x Real) UnaryMinus() Real {
	return Real(-float64(x))
}

type Boolean bool

func (Boolean) Type() Type {
	return BooleanType{}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// types

type Type interface {
	Zero() interface{}
}

type NullType struct{}

func (NullType) Zero() interface{} {
	return nil
}

// object identifier
type ObjectIdentifierType struct{}

func (ObjectIdentifierType) Zero() interface{} {
	return make(DefinitiveIdentifier, 0)
}

// integer
type IntegerType struct {
	NamedNumberList map[string]int
}

func (IntegerType) Zero() interface{} {
	return 0
}

// real
type RealType struct{}

func (RealType) Zero() interface{} {
	return 0.0
}

// boolean
type BooleanType struct{}

func (BooleanType) Zero() interface{} {
	return false
}

type ChoiceType struct {
	AlternativeTypeList []NamedType
	// TODO ExtensionAndException
}

func (ChoiceType) Zero() interface{} {
	return nil
}

////////////////////////////////////////////////
// String types

type RestrictedStringType struct {
	LexType int
}

func (RestrictedStringType) Zero() interface{} {
	return ""
}

type CharacterStringType struct{}

func (CharacterStringType) Zero() interface{} {
	return ""
}

type OctetStringType struct{}

func (OctetStringType) Zero() interface{} {
	return make([]byte, 0)
}

////////////////////////////////////////////////
// sequence type

// TODO Extensions are not supported
type SequenceType struct {
	Components ComponentTypeList
}

func (SequenceType) Zero() interface{} {
	return nil
}

type ComponentTypeList []ComponentType

type ComponentType interface {
	IsComponentType()
}

// "regular" named element of SEQUENCE
type NamedComponentType struct {
	NamedType  NamedType
	IsOptional bool
	Default    *Value
}

func (NamedComponentType) IsComponentType() {}

// reference to other SEQUENCE type to be expanded
type ComponentsOfComponentType struct {
	Type Type
}

func (ComponentsOfComponentType) IsComponentType() {}

// tagged types
type TaggedType struct {
	Tag        Tag
	Type       Type
	TagType    int  // one of TAGS_*
	HasTagType bool // true if explicitly set
}

func (t TaggedType) Zero() interface{} {
	return t.Type.Zero()
}

type Tag struct {
	Class       int
	ClassNumber Value // either DefinedValue or Number
}

const (
	CLASS_CONTEXT_SPECIFIC = iota // when not specified
	CLASS_UNIVERSAL
	CLASS_APPLICATION
	CLASS_PRIVATE
)

type SequenceOfType struct {
	Type Type
}

func (SequenceOfType) Zero() interface{} {
	return make([]interface{}, 0)
}

// BIT STRING with optional named bits
type BitStringType struct {
	NamedBits []NamedBit
}

func (BitStringType) Zero() interface{} {
	return make([]bool, 0)
}

type NamedBit struct {
	Name  Identifier
	Index Value // Number or DefinedValue
}

////////////////////////////////////////////////
// type with constraints
type ConstraintedType struct {
	Type       Type
	Constraint Constraint
}

func (t ConstraintedType) Zero() interface{} {
	return t.Type.Zero()
}

type Constraint struct {
	ConstraintSpec ConstraintSpec
	//ExceptionSpec ExceptionSpec
}

// ConstraintSpec can be SubtypeConstraint or GeneralConstraint
type ConstraintSpec interface {
	IsConstraintSpec()
}

func SingleElementConstraint(elem Elements) Constraint {
	return Constraint{ConstraintSpec: SubtypeConstraint{
		Unions{Intersections{IntersectionElements{Elements: elem}}},
	}}
}

// SubtypeConstraint describes list of element sets that can be used in constainted type
type SubtypeConstraint []ElementSetSpec

func (SubtypeConstraint) IsConstraintSpec() {}

// Union or Exclusion
type ElementSetSpec interface {
	Elements
	IsElementSpec()
}

// Unions is ElementSetSpec and Elements
type Unions []Intersections

func (Unions) IsElementSpec() {}

func (Unions) IsElements() {}

// part of the Union
type Intersections []IntersectionElements

type IntersectionElements struct {
	Elements   Elements
	Exclusions Exclusions
}

// Exclusion is ElementSetSpec and Elements
type Exclusions struct {
	Elements Elements
}

func (Exclusions) IsElementSpec() {}

func (Exclusions) IsElements() {}

// Describe elements of Intersections or Exclusions
type Elements interface {
	IsElements()
}

// subtype elements

// SingleValue is Elements
type SingleValue struct {
	Value
}

func (SingleValue) IsElements() {}

// ValueRange is Elements
type ValueRange struct {
	LowerEndpoint RangeEndpoint
	UpperEndpoint RangeEndpoint
}

func (ValueRange) IsElements() {}

type RangeEndpoint struct {
	Value  Value
	IsOpen bool // X<..<X
}

// IsUnspecified corresponds to MIN or MAX
func (e RangeEndpoint) IsUnspecified() bool {
	return e.Value == nil
}

type TypeConstraint struct {
	Type Type
}

func (TypeConstraint) IsElements() {}

type SizeConstraint struct {
	Constraint Constraint
}

func (SizeConstraint) IsElements() {}

// TODO
type GeneralConstraint struct{}

func (GeneralConstraint) IsConstraintSpec() {}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// values

type Value interface {
	Type() Type
}

// TODO
// represents value reference
type DefinedValue struct{}

func (DefinedValue) Type() Type {
	return nil
}

func (DefinedValue) IsObjectIdComponent() bool {
	return true
}

// element of integer type referenced by name
type IdentifiedIntegerValue struct {
	valueType Type
	Name      string
}

func (x IdentifiedIntegerValue) Type() Type {
	return x.valueType
}

//////////////////////////////
// OID
type ObjectIdentifierValue []ObjIdComponents

func NewObjectIdentifierValue(initial ...ObjIdComponents) ObjectIdentifierValue {
	return append(make(ObjectIdentifierValue, 0), initial...)
}

func (oid ObjectIdentifierValue) Append(other ...ObjIdComponents) ObjectIdentifierValue {
	return ObjectIdentifierValue(append(oid, other...))
}

func (ObjectIdentifierValue) Type() Type {
	return ObjectIdentifierType{}
}

func (ObjectIdentifierValue) IsObjectIdComponent() bool {
	return true
}

// implemented by ObjectIdElement, DefinedValue and ObjectIdentifierValue itself
type ObjIdComponents interface {
	IsObjectIdComponent() bool // fake method for grouping
}

type ObjectIdElement struct {
	Name      string
	Id        int
	Reference *DefinedValue // nil if Id is set explicitly
}

func (ObjectIdElement) IsObjectIdComponent() bool {
	return true
}

// end OID
//////////////////////////////

const (
	GeneralizedTimeName = "GeneralizedTime"
	UTCTimeName         = "UTCTime"
)

var (
	USEFUL_TYPES map[string]Type = map[string]Type{
		GeneralizedTimeName: TaggedType{ // [UNIVERSAL 24] IMPLICIT VisibleString
			Tag:  Tag{Class: CLASS_UNIVERSAL, ClassNumber: Number(24)},
			Type: RestrictedStringType{VisibleString}},
	}
)
