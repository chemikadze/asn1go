// header
%{
package asn1go
%}
////////////////////////////
//  declarations section
//

// extra SymType fields
%union{
    typeref    TypeReference
    identifier Identifier
    number     Number
    real       Real
    numberRepr string
    anyval     interface{}
}

%token WHITESPACE
%token NEWLINE
%token <typeref> TYPEREFERENCE      // note - also used for MODULEREFERENCE (semantics depends on context)
%token <identifier> VALUEIDENTIFIER // note - also used for VALUEREFERENCE (semantics depends on context)
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

%type <anyval> anytoken

//
// end declarations
////////////////////////////

%%

////////////////////////////
// grammar/rules section
//

// Code inside the grammar actions may refer to the variable yylex,
// which holds the yyLexer passed to yyParse.

start : anytoken
    {
        yylex.(*MyLexer).result = $1
    }
;

anytoken : TYPEREFERENCE
            { $$ = $1 }
         | VALUEIDENTIFIER
            { $$ = $1 }
         | NUMBER
            { $$ = $1 }
         | REALNUMBER
            { $$ = $1 }
;

//
// end grammar
////////////////////////////

%%