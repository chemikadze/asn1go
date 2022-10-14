package asn1go

import "encoding/asn1"

// ModuleDefinition represents ASN.1 Module.
// This and all other AST types are named according to their BNF in X.680 document,
// if not specified otherwise.
// See: X.680, section 12.
type ModuleDefinition struct {
	ModuleIdentifier ModuleIdentifier
	// TagDefault is default tagging behavior, one of TAGS_ constants.
	TagDefault           int
	ExtensibilityImplied bool
	ModuleBody           ModuleBody
}

// ModuleIdentifier is root of ASN.1 module.
type ModuleIdentifier struct {
	Reference            string
	DefinitiveIdentifier DefinitiveIdentifier
}

// DefinitiveIdentifier is fully qualified name of the module.
type DefinitiveIdentifier []DefinitiveObjIdComponent

// DefinitiveObjIdComponent is part of DefinitiveIdentifier.
type DefinitiveObjIdComponent struct {
	Name string
	Id   int
}

// Consts for ModuleDefinition.TagDefault.
const (
	TAGS_EXPLICIT = iota
	TAGS_IMPLICIT
	TAGS_AUTOMATIC
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Module body

// ModuleBody holds module body.
// TODO: implement Exports.
type ModuleBody struct {
	AssignmentList AssignmentList
	Imports        []SymbolsFromModule
}

// SymbolsFromModule holds imports from particular module.
type SymbolsFromModule struct {
	SymbolList []Symbol
	Module     GlobalModuleReference
}

// Symbol is exported or imported symbol.
// Only References are supported, ParameterizedReference is not implemented.
// TODO: make IsSymbol private.
type Symbol interface {
	IsSymbol()
}

// GlobalModuleReference fully qualifies module from which symbols are imported.
type GlobalModuleReference struct {
	Reference          string
	AssignedIdentifier Value
}

// AssignmentList holds assignments in module body.
type AssignmentList []Assignment

// Get finds Assignment by defined reference name, or returns nil if not found.
func (l AssignmentList) Get(name string) Assignment {
	for _, assignment := range l {
		if assignment.Reference().Name() == name {
			return assignment
		}
	}
	return nil
}

// GetValue returns ValueAssignment by name, or nil if not found.
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

// GetType returns TypeAssignment by name, or nil if not found.
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

// Assignment is interface for Assignment nodes.
// Only TypeAssignment and ValueAssignment are supported.
// Other assignment types (value sets, xml values, objects) are not implemented.
type Assignment interface {
	Reference() Reference
}

// ValueAssignment defines ValueReference of Type with given Value.
type ValueAssignment struct {
	ValueReference ValueReference
	Type           Type
	Value          Value
}

// Reference implements Assignment.
func (v ValueAssignment) Reference() Reference {
	return v.ValueReference
}

// TypeAssignment defines TypeReference of specified Type.
type TypeAssignment struct {
	TypeReference TypeReference
	Type          Type
}

// Reference implements Assignment.
func (v TypeAssignment) Reference() Reference {
	return v.TypeReference
}

// NamedType is a identifier-type tuple.
// It's used as element in SequenceType, SetType, ChoiceType and some other types.
type NamedType struct {
	Identifier Identifier
	Type       Type
}

// Zero implements Type.
func (t NamedType) Zero() interface{} {
	return t.Type.Zero()
}

// isChoiceExtension implements ChoiceExtension.
func (t NamedType) isChoiceExtension() {
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// References

// Reference is an interface for a reference.
// There are different reference types depending on a referent.
type Reference interface {
	Name() string
}

// TypeReference refers to a type defined in same module or imported from different module.
// This is lexical construct, named `typereference` in the doc.
// See X.680, section 11.2.
type TypeReference string

// Name implements Reference.
func (r TypeReference) Name() string {
	return string(r)
}

// Zero implements Type.
func (r TypeReference) Zero() interface{} {
	return nil
}

// IsSymbol implements Symbol.
func (TypeReference) IsSymbol() {}

// ValueReference refers to a value defined in same module or imported from different module.
// This is lexical construct, named `valuereference` in the doc.
// See X.680, section 11.2.
type ValueReference string

// Name implements Reference.
func (r ValueReference) Name() string {
	return string(r)
}

// IsSymbol implements Symbol.
func (ValueReference) IsSymbol() {}

// ModuleReference refers to a module.
// This is lexical construct, named `modulereference` in the doc.
// See X.680, section 11.5.
type ModuleReference string

// IsSymbol implements Symbol.
func (ModuleReference) IsSymbol() {}

// Identifier is a non-referential identifier.
// This is a lexical construct, named `identifier` in the doc.
// See X.680, section 11.3.
type Identifier string

// Name implements Reference.
// TODO: it should not implement Reference according to the BNF.
func (id Identifier) Name() string {
	return string(id)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Values

// Number is an integer value.
// This is a lexical construct, named `number` in the doc.
// See X.680, section 11.8.
type Number int

// IntValue returns value of the number.
func (x Number) IntValue() int {
	return int(x)
}

// Type implements Value.
func (Number) Type() Type {
	return IntegerType{}
}

// UnaryMinus returns negated Number.
func (x Number) UnaryMinus() Number {
	return Number(-int(x))
}

// Real is a floating point value.
// This is a lexical construct, named `realnumber` in the doc.
// See X.680, section 11.9.
type Real float64

// Type implements Value.
func (x Real) Type() Type {
	return RealType{}
}

// UnaryMinus returns negated value.
func (x Real) UnaryMinus() Real {
	return Real(-float64(x))
}

// Boolean is a bool value.
// It is named BooleanValue in BNF.
// See X.680, section 17.3.
type Boolean bool

// Type implements Value.
func (Boolean) Type() Type {
	return BooleanType{}
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// types

// Type is a builtin, referenced or constrained type.
// TODO: replace Zero with isType.
type Type interface {
	Zero() interface{}
}

// NullType is an ast representation of NULL type.
type NullType struct{}

// Zero implements Type.
func (NullType) Zero() interface{} {
	return nil
}

// ObjectIdentifierType is an ast representation of OBJECT IDENTIFIER type.
// TODO: implement these properly.
type ObjectIdentifierType struct{}

// Zero implements Type.
func (ObjectIdentifierType) Zero() interface{} {
	return make(DefinitiveIdentifier, 0)
}

// IntegerType is an ast representation of INTEGER type.
type IntegerType struct {
	NamedNumberList map[string]int
}

// Zero implelents Type.
func (IntegerType) Zero() interface{} {
	return 0
}

// EnumeratedType is an ast representation of ENUMERATED type.
// TODO: implement enumerations properly.
type EnumeratedType struct {
}

// Zero implements Type.
func (EnumeratedType) Zero() interface{} {
	return asn1.Enumerated(0)
}

// RealType is an ast representation of REAL type.
type RealType struct{}

// Zero implements Type.
func (RealType) Zero() interface{} {
	return 0.0
}

// BooleanType is an ast representation of BOOLEAN type.
type BooleanType struct{}

// Zero implements Type.
func (BooleanType) Zero() interface{} {
	return false
}

// ChoiceType is an ast representation of CHOICE type.
// It is partially implemented, exceptions are ignored.
type ChoiceType struct {
	AlternativeTypeList []NamedType
	ExtensionTypes      []ChoiceExtension
	// TODO ExtensionAndException
}

// Zero implements Type.
func (ChoiceType) Zero() interface{} {
	return nil
}

// ChoiceExtension is a type for choice extensions.
// Only NamedType is implemented, ExtensionAdditionAlternativesGroup is not supported.
type ChoiceExtension interface {
	isChoiceExtension()
}

////////////////////////////////////////////////
// String types

// RestrictedStringType is a type for strings with restricted character set.
// It is defined as RestrictedCharacterStringType in BNF.
// See X.680, section 36.1.
type RestrictedStringType struct {
	// LexType is a lexem value for restricted string (e.g. IA5String).
	LexType int
}

// Zero implements Value.
func (RestrictedStringType) Zero() interface{} {
	return ""
}

// CharacterStringType is an ast representation of CHARACTER STRING type.
// It is defined as UnrestrictedCharacterStringType in BNF.
type CharacterStringType struct{}

// Zero implements Type.
func (CharacterStringType) Zero() interface{} {
	return ""
}

// OctetStringType is an ast representation of OCTET STRING type.
type OctetStringType struct{}

// Zero implements Type.
func (OctetStringType) Zero() interface{} {
	return make([]byte, 0)
}

////////////////////////////////////////////////
// sequence type

// SequenceType is an ast representation of SEQUENCE type.
// TODO: Extensions are not supported.
type SequenceType struct {
	Components         ComponentTypeList
	ExtensionAdditions ExtensionAdditions
}

// Zero implements Type.
func (SequenceType) Zero() interface{} {
	return nil
}

// ExtensionAdditions is a list of extension additions in SET or SEQUENCE.
type ExtensionAdditions []ExtensionAddition

// ExtensionAddition is a single element of extension addition.
type ExtensionAddition interface {
	isExtensionAddition()
}

// ComponentTypeLists is not used in AST directly but is used in parser for intermediate representation.
type ComponentTypeLists struct {
	Components         ComponentTypeList
	ExtensionAdditions ExtensionAdditions
	TrailingComponents ComponentTypeList
}

// ComponentTypeList is a list of ComponentType.
type ComponentTypeList []ComponentType

// ComponentType is a component type of SEQUENCE or SET.
// It can be used in ExtensionAddition context, so types implementing it must implement both.
// TODO: rename IsComponentType to isComponentType.
type ComponentType interface {
	ExtensionAddition
	IsComponentType()
}

// NamedComponentType is an entry in a SEQUENCE definition.
type NamedComponentType struct {
	NamedType  NamedType
	IsOptional bool
	Default    *Value
}

// IsComponentType implements ComponentType.
func (NamedComponentType) IsComponentType() {}

// isExtensionAddition implements ExtensionAddition.
func (NamedComponentType) isExtensionAddition() {}

// ComponentsOfComponentType is content of COMPONENTS OF clause.
type ComponentsOfComponentType struct {
	Type Type
}

// IsComponentType implements ComponentType.
func (ComponentsOfComponentType) IsComponentType() {}

// isExtensionAddition implements ExtensionAddition.
func (ComponentsOfComponentType) isExtensionAddition() {}

// SetType is an ast representation of SEQUENCE type.
// TODO: Extensions are not supported.
type SetType struct {
	Components         ComponentTypeList
	ExtensionAdditions ExtensionAdditions
}

// Zero implements Type.
func (SetType) Zero() interface{} {
	return nil
}

// TaggedType is a tagged type.
type TaggedType struct {
	// Tag assigned to a type.
	Tag Tag
	// Type that is being tagged.
	Type Type
	// TagType is one of TAGS_ constants.
	// E.g. IMPLICIT, EXPLICIT.
	TagType int
	// HasTagType is set to true if TagType was explicitly specified in module syntax.
	// Otherwise, TagType would hold module default.
	HasTagType bool // true if explicitly set
}

// Zero implements Type.
func (t TaggedType) Zero() interface{} {
	return t.Type.Zero()
}

// Tag is a tag value.
type Tag struct {
	// Class is a tag class, one of CLASS_ constants.
	// E.g. UNIVERSAL, APPLICATION.
	Class int
	// ClassNumber is a tag value.
	// Will hold DefinedValue or Number.
	ClassNumber Value
}

// Tag class constants.
const (
	CLASS_CONTEXT_SPECIFIC = iota // when not specified
	CLASS_UNIVERSAL
	CLASS_APPLICATION
	CLASS_PRIVATE
)

// SequenceOfType is an ast representation of SEQUENCE OF type.
type SequenceOfType struct {
	Type Type
}

// Zero implements Type.
func (SequenceOfType) Zero() interface{} {
	return make([]interface{}, 0)
}

// SetOfType is an ast representation of SET OF type.
type SetOfType struct {
	Type Type
}

// Zero implements Type.
func (SetOfType) Zero() interface{} {
	return make([]interface{}, 0)
}

// AnyType is an ast representation of ANY type.
// It is NOT defined in X.680, but added for compatibility with older ASN definitions, e.g. X.509.
// See X.208, section 27.
type AnyType struct {
	// Identifier is set if IDENTIFIED BY is provided.
	Identifier Identifier
}

// Zero implements Type.
func (AnyType) Zero() interface{} {
	return nil
}

// BitStringType is an ast representation of BIT STRING type.
type BitStringType struct {
	NamedBits []NamedBit
}

// Zero implements Type.
func (BitStringType) Zero() interface{} {
	return make([]bool, 0)
}

// NamedBit is a named bit in BIT STRING.
type NamedBit struct {
	Name Identifier
	// Index is index of the bit.
	// Will have Number or DefinedValue type.
	Index Value
}

////////////////////////////////////////////////
// type with constraints

// ConstraintedType is a type with constraints applied.
// General constraints are not implemented in parser, and exception spec is not preserved in AST.
type ConstraintedType struct {
	Type       Type
	Constraint Constraint
}

// Zero implements Type.
func (t ConstraintedType) Zero() interface{} {
	return t.Type.Zero()
}

// Constraint is a constraint applied to a type.
type Constraint struct {
	ConstraintSpec ConstraintSpec
	//ExceptionSpec ExceptionSpec
}

// ConstraintSpec can be SubtypeConstraint or GeneralConstraint.
// GeneralConstraint is not implemented.
// TODO: rename IsConstraintSpec to isConstraintSpec.
type ConstraintSpec interface {
	IsConstraintSpec()
}

// SingleElementConstraint is a Constraint of single intersection elements.
func SingleElementConstraint(elem Elements) Constraint {
	return Constraint{ConstraintSpec: SubtypeConstraint{
		Unions{Intersections{IntersectionElements{Elements: elem}}},
	}}
}

// SubtypeConstraint describes list of element sets that can be used in constainted type
type SubtypeConstraint []ElementSetSpec

// IsConstraintSpec implements ConstraintSpec.
func (SubtypeConstraint) IsConstraintSpec() {}

// ElementSetSpec is element of the SubtypeConstraint.
// TODO: rename IsElementSpec to isElementSpec.
type ElementSetSpec interface {
	Elements
	IsElementSpec()
}

// Unions is ElementSetSpec and Elements
type Unions []Intersections

// IsElementSpec implements ElementSpec.
func (Unions) IsElementSpec() {}

// IsElements implements Elements.
func (Unions) IsElements() {}

// Intersections is a part of SubtypeConstraint.
type Intersections []IntersectionElements

// IntersectionElements is part of Intersections.
type IntersectionElements struct {
	Elements   Elements
	Exclusions Exclusions
}

// Exclusions are Elements excluded from IntersectionElements.
type Exclusions struct {
	Elements Elements
}

// IsElementSpec implements ElementSpec.
func (Exclusions) IsElementSpec() {}

// IsElements implements Elements.
func (Exclusions) IsElements() {}

// Elements is one of subtype elements (values, type constraints, etc).
// TODO: rename IsElements to Elements.
type Elements interface {
	IsElements()
}

// subtype elements

// SingleValue is included or excluded value.
type SingleValue struct {
	Value
}

// IsElements implements Elements.
func (SingleValue) IsElements() {}

// ValueRange is included or excluded range.
type ValueRange struct {
	LowerEndpoint RangeEndpoint
	UpperEndpoint RangeEndpoint
}

// IsElements implements Elements.
func (ValueRange) IsElements() {}

// RangeEndpoint is left or right side of the ValueRange.
type RangeEndpoint struct {
	Value  Value
	IsOpen bool // X<..<X
}

// IsUnspecified corresponds to MIN or MAX of the range.
func (e RangeEndpoint) IsUnspecified() bool {
	return e.Value == nil
}

// TypeConstraint is a type constraint.
type TypeConstraint struct {
	Type Type
}

// IsElements implements Elements.
func (TypeConstraint) IsElements() {}

// SizeConstraint is a SIZE constraint expressed by another Constraint.
type SizeConstraint struct {
	Constraint Constraint
}

// IsElements implements Elements.
func (SizeConstraint) IsElements() {}

// InnerTypeConstraint is WITH COMPONENT constraint.
// Contents are not represented in parsed AST and are ignored.
type InnerTypeConstraint struct{}

// IsElements implements Elements.
func (InnerTypeConstraint) IsElements() {}

// GeneralConstraint is not implemented.
// It is defined by X.682.
// TODO: implement or remove.
type GeneralConstraint struct{}

// IsConstraintSpec implements ConstraintSpec.
func (GeneralConstraint) IsConstraintSpec() {}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// values

// Value is interface for values.
type Value interface {
	Type() Type
}

// DefinedValue represents value reference.
// TODO: implement properly.
type DefinedValue struct{}

// Type implements Value.
func (DefinedValue) Type() Type {
	return nil
}

// IsObjectIdComponent implements ObjectIdComponent.
func (DefinedValue) IsObjectIdComponent() bool {
	return true
}

// IdentifiedIntegerValue is named value defined for the type.
// TODO: use of these in assignments is not implemented.
type IdentifiedIntegerValue struct {
	valueType Type
	Name      string
}

// Type implements Value.
func (x IdentifiedIntegerValue) Type() Type {
	return x.valueType
}

//////////////////////////////
// OID

// ObjectIdentifierValue is a value of OBJECT IDENTIFIER type.
type ObjectIdentifierValue []ObjIdComponents

// NewObjectIdentifierValue creates ObjectIdentifierValue from components.
// TODO: remove.
func NewObjectIdentifierValue(initial ...ObjIdComponents) ObjectIdentifierValue {
	return append(make(ObjectIdentifierValue, 0), initial...)
}

// Append adds elements to the value.
// TODO: remove.
func (oid ObjectIdentifierValue) Append(other ...ObjIdComponents) ObjectIdentifierValue {
	return ObjectIdentifierValue(append(oid, other...))
}

// Type implements Value.
func (ObjectIdentifierValue) Type() Type {
	return ObjectIdentifierType{}
}

// IsObjectIdComponent implements ObjectIdComponent.
// TODO: clarify why is this needed.
func (ObjectIdentifierValue) IsObjectIdComponent() bool {
	return true
}

// ObjIdComponents is interface for components of ObjectIdentifierValues.
// TODO: .y seems a bit convoluted, explore if this can be simplified.
// TODO: rename IsObjectIdComponent as isObjectIdComponent.
type ObjIdComponents interface {
	IsObjectIdComponent() bool // fake method for grouping
}

// ObjectIdElement is object id element in name, number or name and number form.
type ObjectIdElement struct {
	// Name is non-empty in name-and-number form.
	Name string
	// Id is set in Number or NameAndNumber form, if number is specified as number literal.
	// TODO: rename to ID.
	Id int
	// Reference is set in NameAndNumber form, when number is provided as DefinedValue.
	Reference *DefinedValue // nil if Id is set explicitly
}

// IsObjectIdComponent implements ObjectIdComponent.
func (ObjectIdElement) IsObjectIdComponent() bool {
	return true
}

// end OID
//////////////////////////////

// Names for useful types.
const (
	GeneralizedTimeName = "GeneralizedTime"
	UTCTimeName         = "UTCTime"
)

var (
	// USEFUL_TYPES are defined in X.680, section 41.
	// These are built-in types that behave like type assignments that are always in scope.
	// TODO: clarify why UTCTimeName is missing here.
	USEFUL_TYPES map[string]Type = map[string]Type{
		GeneralizedTimeName: TaggedType{ // [UNIVERSAL 24] IMPLICIT VisibleString
			Tag:  Tag{Class: CLASS_UNIVERSAL, ClassNumber: Number(24)},
			Type: RestrictedStringType{VisibleString}},
	}
)
