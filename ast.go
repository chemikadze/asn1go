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

type Assignment interface {
	Reference() Reference
}

type ValueAssignment struct {
	ValueReference ValueReference
	Type           Type
	Value          Value
}

func (v ValueAssignment) Reference() Reference {
	return v.ValueReference
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

// value reference
type ValueReference string

func (r ValueReference) Name() string {
	return string(r)
}

// identifier type
type Identifier string

// number lexem
type Number int

func (x Number) IntValue() int {
	return int(x)
}

// real lexem
type Real float64

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// types

type Type interface {
	Zero() interface{}
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

// boolean
type BooleanType struct{}

func (BooleanType) Zero() interface{} {
	return false
}

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
