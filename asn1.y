// header
%{
package asn1go
%}
////////////////////////////
//  declarations section
//

// extra SymType fields
%union{
    name       string
    number     Number
    real       Real
    numberRepr string

    TagDefault int
    ExtensionDefault bool
    ModuleIdentifier ModuleIdentifier
}

%token WHITESPACE
%token NEWLINE
%token <name> TYPEORMODULEREFERENCE
%token <name> VALUEIDENTIFIER
%token <number> NUMBER
%token <real> REALNUMBER
%token <bstring> BSTRING          // TODO not implemented in lexer
%token <bstring> XMLBSTRING       // TODO not implemented in lexer
%token <hstring> HSTRING          // TODO not implemented in lexer
%token <hstring> XMLHSTRING       // TODO not implemented in lexer
%token <cstring> CSTRING          // TODO not implemented in lexer
%token <cstring> XMLCSTRING       // TODO not implemented in lexer
%token ASSIGNMENT
%token RANGE_SEPARATOR
%token ELLIPSIS
%token LEFT_VERSION_BRACKETS
%token RIGHT_VERSION_BRACKETS
%token XML_END_TAG_START    // TODO not implemented in lexer
%token XML_SINGLE_START_END // TODO not implemented in lexer
%token XML_BOOLEAN_TRUE     // TODO not implemented in lexer
%token XML_BOOLEAN_FALSE    // TODO not implemented in lexer
%token XMLASN1TYPENAME      // TODO not implemented in lexer

// single-symbol tokens
%token OPEN_CURLY  // "{"
%token CLOSE_CURLY  // "}"
%token LESS  // "<"
%token GREATER  // ">"
%token COMMA  // ","
%token DOT  // "."
%token OPEN_ROUND  // "("
%token CLOSE_ROUND  // ")"
%token OPEN_SQUARE  // "["
%token CLOSE_SQUARE  // "]"
%token MINUS  // "-" (HYPEN-MINUS)
%token COLON  // ":"
%token EQUALS  // "="
%token QUOTATION_MARK  // """ (QUOTATION MARK)
%token APOSTROPHE  // "'" (APOSTROPHE)
%token SPACE  // " " (SPACE)  // TODO won't be parsed probably
%token SEMICOLON  // ";"
%token AT  // "@"
%token PIPE  // "|"
%token EXCLAMATION  // "!"
%token CARET  // "^"

// reserved words
%token ABSENT
%token ENCODED
%token INTEGER
%token RELATIVE_OID
%token ABSTRACT_SYNTAX
%token END
%token INTERSECTION
%token SEQUENCE
%token ALL
%token ENUMERATED
%token ISO646String
%token SET
%token APPLICATION
%token EXCEPT
%token MAX
%token SIZE
%token AUTOMATIC
%token EXPLICIT
%token MIN
%token STRING
%token BEGIN
%token EXPORTS
%token MINUS_INFINITY
%token SYNTAX
%token BIT
%token EXTENSIBILITY
%token NULL
%token T61String
%token BMPString
%token EXTERNAL
%token NumericString
%token TAGS
%token BOOLEAN
%token FALSE
%token OBJECT
%token TeletexString
%token BY
%token FROM
%token ObjectDescriptor
%token TRUE
%token CHARACTER
%token GeneralizedTime
%token OCTET
%token TYPE_IDENTIFIER
%token CHOICE
%token GeneralString
%token OF
%token UNION
%token CLASS
%token GraphicString
%token OPTIONAL
%token UNIQUE
%token COMPONENT
%token IA5String
%token PATTERN
%token UNIVERSAL
%token COMPONENTS
%token IDENTIFIER
%token PDV
%token UniversalString
%token CONSTRAINED
%token IMPLICIT
%token PLUS_INFINITY
%token UTCTime
%token CONTAINING
%token IMPLIED
%token PRESENT
%token UTF8String
%token DEFAULT
%token IMPORTS
%token PrintableString
%token VideotexString
%token DEFINITIONS
%token INCLUDES
%token PRIVATE
%token VisibleString
%token EMBEDDED
%token INSTANCE
%token REAL
%token WITH

%type <name> modulereference
%type <ExtensionDefault> ExtensionDefault
%type <TagDefault> TagDefault
%type <ModuleIdentifier> ModuleIdentifier

//
// end declarations
////////////////////////////

%%

////////////////////////////
// grammar/rules section
//

// Code inside the grammar actions may refer to the variable yylex,
// which holds the yyLexer passed to yyParse.

//start : anytoken
//    {
//        yylex.(*MyLexer).result = $1
//    }
//;

//anytoken : TYPEREFERENCE
//            { $$ = $1 }
//         | VALUEIDENTIFIER
//            { $$ = $1 }
//         | NUMBER
//            { $$ = $1 }
//         | REALNUMBER
//            { $$ = $1 }
//;

ModuleDefinition :
    ModuleIdentifier
    DEFINITIONS
    TagDefault
    ExtensionDefault
    ASSIGNMENT
    BEGIN
    ModuleBody
    END
    { yylex.(*MyLexer).result = &ModuleDefinition{ModuleIdentifier: $1, TagDefault: $3, ExtensibilityImplied: $4} }
;

typereference: TYPEORMODULEREFERENCE;

modulereference: TYPEORMODULEREFERENCE;

valuereference: VALUEIDENTIFIER;

identifier: VALUEIDENTIFIER;

ModuleIdentifier :
                   modulereference
                   DefinitiveIdentifier
                   { $$ = ModuleIdentifier{Reference: $1} }
;

DefinitiveIdentifier : OPEN_CURLY DefinitiveObjIdComponentList CLOSE_CURLY
                     | /*empty*/
;

DefinitiveObjIdComponentList :  DefinitiveObjIdComponent
                             | DefinitiveObjIdComponent DefinitiveObjIdComponentList
;

DefinitiveObjIdComponent : NameForm
                         | DefinitiveNumberForm
                         | DefinitiveNameAndNumberForm
;

DefinitiveNumberForm : NUMBER
;

DefinitiveNameAndNumberForm : identifier OPEN_ROUND DefinitiveNumberForm CLOSE_ROUND
;

TagDefault : EXPLICIT TAGS   { $$ = TAGS_EXPLICIT }
           | IMPLICIT TAGS   { $$ = TAGS_IMPLICIT }
           | AUTOMATIC TAGS  { $$ = TAGS_AUTOMATIC }
           | /*empty*/       { $$ = TAGS_EXPLICIT }
;

ExtensionDefault : EXTENSIBILITY IMPLIED { $$ = true }
                 | /*empty*/             { $$ = false }
;

ModuleBody : Exports Imports AssignmentList
           | /*empty*/
;


Exports : EXPORTS SymbolsExported SEMICOLON
        | EXPORTS ALL SEMICOLON
        | /*empty*/
;

SymbolsExported : SymbolList
                | /*empty*/
;

Imports : IMPORTS SymbolsImported SEMICOLON
        | /*empty*/
;

SymbolsImported : SymbolsFromModuleList
                | /*empty*/
;

SymbolsFromModuleList : SymbolsFromModule
                      | SymbolsFromModuleList SymbolsFromModule
;

SymbolsFromModule : SymbolList FROM GlobalModuleReference
;

GlobalModuleReference : modulereference AssignedIdentifier
;

AssignedIdentifier : "t" "o" "d" "o"
//                     ObjectIdentifierValue
//                   | DefinedValue
//                   | /*empty*/
;

SymbolList : Symbol
           | SymbolList COMMA Symbol
;

Symbol : Reference
//       | ParameterizedReference
;

Reference : modulereference // modulereference
          | valuereference       // valuereference
//          | objectclassreference
//          | objectreference
//          | objectsetreference
;

AssignmentList : Assignment
               | AssignmentList Assignment
;

Assignment : TypeAssignment
           | ValueAssignment
//           | XMLValueAssignment
//           | ValueSetTypeAssignment
//           | ObjectClassAssignment
//           | ObjectAssignment
//           | ObjectSetAssignment
//           | ParameterizedAssignment
;

// 15.1

TypeAssignment : typereference ASSIGNMENT Type
;

ValueAssignment : valuereference Type ASSIGNMENT Value
;

// 16.1

Type : BuiltinType
//     | ReferencedType
//     | ConstrainedType
;

// 16.2

BuiltinType : //BitStringType
            | BooleanType
//            | CharacterStringType
//            | ChoiceType
//            | EmbeddedPDVType
//            | EnumeratedType
//            | ExternalType
//            | InstanceOfType
            | IntegerType
//            | NullType
//            | ObjectClassFieldType
//            | ObjectIdentifierType
//            | OctetStringType
//            | RealType
//            | RelativeOIDType
//            | SequenceType
//            | SequenceOfType
//            | SetType
//            | SetOfType
//            | TaggedType
;

// 16.7

Value : BuiltinValue
//      | ReferencedValue
//      | ObjectClassFieldValue
;

// 16.8

// TODO
BuiltinValue : // BitStringValue
             | BooleanValue
//             | CharacterStringValue
//             | ChoiceValue
//             | EmbeddedPDVValue
//             | EnumeratedValue
//             | ExternalValue
//             | InstanceOfValue
             | IntegerValue
//             | NullValue
//             | ObjectIdentifierValue
//             | OctetStringValue
//             | RealValue
//             | RelativeOIDValue
//             | SequenceValue
//             | SequenceOfValue
//             | SetValue
//             | SetOfValue
//             | TaggedValue
;

// 17.3

BooleanType : BOOLEAN
;

BooleanValue : TRUE | FALSE
;

// 18.1

IntegerType : INTEGER
            | INTEGER OPEN_CURLY NamedNumberList CLOSE_CURLY
;

NamedNumberList : NamedNumber
                | NamedNumberList COMMA NamedNumber
;

NamedNumber : identifier OPEN_ROUND SignedNumber CLOSE_ROUND
//          | identifier OPEN_ROUND DefinedValue CLOSE_ROUND
;

SignedNumber : NUMBER
             | "-" NUMBER
;

// 18.9

IntegerValue : SignedNumber
             | identifier
;

// 31.3

NameForm : identifier
;

//
// end grammar
////////////////////////////

%%