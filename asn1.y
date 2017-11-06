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
%token <typeref> TYPEREFERENCE    // note - also used for MODULEREFERENCE (semantics depends on context)
%token <identifier> IDENTIFIER    // note - also used for VALUEREFERENCE (semantics depends on context)
%token <number> NUMBER
%token <real> REAL

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
         | IDENTIFIER
            { $$ = $1 }
         | NUMBER
            { $$ = $1 }
         | REAL
            { $$ = $1 }
;

//
// end grammar
////////////////////////////

%%