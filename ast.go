package asn1go

type AstNode interface{}

type ModuleDefinition struct {
	ModuleIdentifier     ModuleIdentifier
	TagDefault           int
	ExtensibilityImplied bool
	Assignments          map[Reference]AssignableItem
}

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

type AssignableItem interface{}

type Reference interface {
	Name() string
}

type TypeReference string

func (r TypeReference) Name() string {
	return string(r)
}

type ValueReference string

func (r ValueReference) Name() string {
	return string(r)
}

type Identifier string

type Number int

func (x Number) IntValue() int {
	return int(x)
}

type Real float64
