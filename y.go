//line asn1.y:3
package asn1go

import __yyfmt__ "fmt"

//line asn1.y:3
import (
	"fmt"
	"math"
)

//line asn1.y:15
type yySymType struct {
	yys        int
	name       string
	numberRepr string

	Number                       Number
	Real                         Real
	TagDefault                   int
	ExtensionDefault             bool
	ModuleIdentifier             ModuleIdentifier
	DefinitiveObjIdComponent     DefinitiveObjIdComponent
	DefinitiveObjIdComponentList []DefinitiveObjIdComponent
	DefinitiveIdentifier         DefinitiveIdentifier
	Type                         Type
	ObjIdComponents              ObjIdComponents
	DefinedValue                 DefinedValue
	ObjectIdentifierValue        ObjectIdentifierValue
	Value                        Value
	Assignment                   Assignment
	AssignmentList               AssignmentList
	ModuleBody                   ModuleBody
	ValueReference               ValueReference
	TypeReference                TypeReference
	Constraint                   Constraint
	ConstraintSpec               ConstraintSpec
	ElementSetSpec               ElementSetSpec
	Unions                       Unions
	Intersections                Intersections
	IntersectionElements         IntersectionElements
	Exclusions                   Exclusions
	Elements                     Elements
	SubtypeConstraint            SubtypeConstraint
	RangeEndpoint                RangeEndpoint
	NamedType                    NamedType
	ComponentType                ComponentType
	ComponentTypeList            ComponentTypeList
	SequenceType                 SequenceType
	Tag                          Tag
	Class                        int
	SequenceOfType               SequenceOfType
	NamedBitList                 []NamedBit
	NamedBit                     NamedBit
	Imports                      []SymbolsFromModule
	SymbolsFromModule            SymbolsFromModule
	SymbolList                   []Symbol
	Symbol                       Symbol
	GlobalModuleReference        GlobalModuleReference
	AlternativeTypeList          []NamedType
	ChoiceType                   ChoiceType
}

const WHITESPACE = 57346
const NEWLINE = 57347
const TYPEORMODULEREFERENCE = 57348
const VALUEIDENTIFIER = 57349
const NUMBER = 57350
const BSTRING = 57351
const XMLBSTRING = 57352
const HSTRING = 57353
const XMLHSTRING = 57354
const CSTRING = 57355
const XMLCSTRING = 57356
const ASSIGNMENT = 57357
const RANGE_SEPARATOR = 57358
const ELLIPSIS = 57359
const LEFT_VERSION_BRACKETS = 57360
const RIGHT_VERSION_BRACKETS = 57361
const XML_END_TAG_START = 57362
const XML_SINGLE_START_END = 57363
const XML_BOOLEAN_TRUE = 57364
const XML_BOOLEAN_FALSE = 57365
const XMLASN1TYPENAME = 57366
const EXPONENT = 57367
const OPEN_CURLY = 57368
const CLOSE_CURLY = 57369
const LESS = 57370
const GREATER = 57371
const COMMA = 57372
const DOT = 57373
const OPEN_ROUND = 57374
const CLOSE_ROUND = 57375
const OPEN_SQUARE = 57376
const CLOSE_SQUARE = 57377
const MINUS = 57378
const COLON = 57379
const EQUALS = 57380
const QUOTATION_MARK = 57381
const APOSTROPHE = 57382
const SPACE = 57383
const SEMICOLON = 57384
const AT = 57385
const PIPE = 57386
const EXCLAMATION = 57387
const CARET = 57388
const ABSENT = 57389
const ENCODED = 57390
const INTEGER = 57391
const RELATIVE_OID = 57392
const ABSTRACT_SYNTAX = 57393
const END = 57394
const INTERSECTION = 57395
const SEQUENCE = 57396
const ALL = 57397
const ENUMERATED = 57398
const ISO646String = 57399
const SET = 57400
const APPLICATION = 57401
const EXCEPT = 57402
const MAX = 57403
const SIZE = 57404
const AUTOMATIC = 57405
const EXPLICIT = 57406
const MIN = 57407
const STRING = 57408
const BEGIN = 57409
const EXPORTS = 57410
const MINUS_INFINITY = 57411
const SYNTAX = 57412
const BIT = 57413
const EXTENSIBILITY = 57414
const NULL = 57415
const T61String = 57416
const BMPString = 57417
const EXTERNAL = 57418
const NumericString = 57419
const TAGS = 57420
const BOOLEAN = 57421
const FALSE = 57422
const OBJECT = 57423
const TeletexString = 57424
const BY = 57425
const FROM = 57426
const ObjectDescriptor = 57427
const TRUE = 57428
const CHARACTER = 57429
const GeneralizedTime = 57430
const OCTET = 57431
const TYPE_IDENTIFIER = 57432
const CHOICE = 57433
const GeneralString = 57434
const OF = 57435
const UNION = 57436
const CLASS = 57437
const GraphicString = 57438
const OPTIONAL = 57439
const UNIQUE = 57440
const COMPONENT = 57441
const IA5String = 57442
const PATTERN = 57443
const UNIVERSAL = 57444
const COMPONENTS = 57445
const IDENTIFIER = 57446
const PDV = 57447
const UniversalString = 57448
const CONSTRAINED = 57449
const IMPLICIT = 57450
const PLUS_INFINITY = 57451
const UTCTime = 57452
const CONTAINING = 57453
const IMPLIED = 57454
const PRESENT = 57455
const UTF8String = 57456
const DEFAULT = 57457
const IMPORTS = 57458
const PrintableString = 57459
const VideotexString = 57460
const DEFINITIONS = 57461
const INCLUDES = 57462
const PRIVATE = 57463
const VisibleString = 57464
const EMBEDDED = 57465
const INSTANCE = 57466
const REAL = 57467
const WITH = 57468

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"WHITESPACE",
	"NEWLINE",
	"TYPEORMODULEREFERENCE",
	"VALUEIDENTIFIER",
	"NUMBER",
	"BSTRING",
	"XMLBSTRING",
	"HSTRING",
	"XMLHSTRING",
	"CSTRING",
	"XMLCSTRING",
	"ASSIGNMENT",
	"RANGE_SEPARATOR",
	"ELLIPSIS",
	"LEFT_VERSION_BRACKETS",
	"RIGHT_VERSION_BRACKETS",
	"XML_END_TAG_START",
	"XML_SINGLE_START_END",
	"XML_BOOLEAN_TRUE",
	"XML_BOOLEAN_FALSE",
	"XMLASN1TYPENAME",
	"EXPONENT",
	"OPEN_CURLY",
	"CLOSE_CURLY",
	"LESS",
	"GREATER",
	"COMMA",
	"DOT",
	"OPEN_ROUND",
	"CLOSE_ROUND",
	"OPEN_SQUARE",
	"CLOSE_SQUARE",
	"MINUS",
	"COLON",
	"EQUALS",
	"QUOTATION_MARK",
	"APOSTROPHE",
	"SPACE",
	"SEMICOLON",
	"AT",
	"PIPE",
	"EXCLAMATION",
	"CARET",
	"ABSENT",
	"ENCODED",
	"INTEGER",
	"RELATIVE_OID",
	"ABSTRACT_SYNTAX",
	"END",
	"INTERSECTION",
	"SEQUENCE",
	"ALL",
	"ENUMERATED",
	"ISO646String",
	"SET",
	"APPLICATION",
	"EXCEPT",
	"MAX",
	"SIZE",
	"AUTOMATIC",
	"EXPLICIT",
	"MIN",
	"STRING",
	"BEGIN",
	"EXPORTS",
	"MINUS_INFINITY",
	"SYNTAX",
	"BIT",
	"EXTENSIBILITY",
	"NULL",
	"T61String",
	"BMPString",
	"EXTERNAL",
	"NumericString",
	"TAGS",
	"BOOLEAN",
	"FALSE",
	"OBJECT",
	"TeletexString",
	"BY",
	"FROM",
	"ObjectDescriptor",
	"TRUE",
	"CHARACTER",
	"GeneralizedTime",
	"OCTET",
	"TYPE_IDENTIFIER",
	"CHOICE",
	"GeneralString",
	"OF",
	"UNION",
	"CLASS",
	"GraphicString",
	"OPTIONAL",
	"UNIQUE",
	"COMPONENT",
	"IA5String",
	"PATTERN",
	"UNIVERSAL",
	"COMPONENTS",
	"IDENTIFIER",
	"PDV",
	"UniversalString",
	"CONSTRAINED",
	"IMPLICIT",
	"PLUS_INFINITY",
	"UTCTime",
	"CONTAINING",
	"IMPLIED",
	"PRESENT",
	"UTF8String",
	"DEFAULT",
	"IMPORTS",
	"PrintableString",
	"VideotexString",
	"DEFINITIONS",
	"INCLUDES",
	"PRIVATE",
	"VisibleString",
	"EMBEDDED",
	"INSTANCE",
	"REAL",
	"WITH",
	"\"t\"",
	"\"o\"",
	"\"d\"",
	"\",\"",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line asn1.y:964

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 31,
	52, 24,
	-2, 27,
	-1, 164,
	44, 210,
	94, 210,
	-2, 206,
	-1, 166,
	46, 213,
	53, 213,
	-2, 208,
	-1, 170,
	60, 216,
	-2, 214,
	-1, 178,
	16, 234,
	28, 234,
	-2, 228,
	-1, 285,
	46, 213,
	53, 213,
	-2, 209,
}

const yyPrivate = 57344

const yyLast = 801

var yyAct = [...]int{

	150, 178, 187, 186, 253, 333, 273, 163, 17, 269,
	209, 289, 244, 220, 17, 198, 149, 193, 190, 168,
	166, 228, 170, 180, 231, 145, 216, 156, 293, 306,
	281, 119, 323, 19, 223, 211, 223, 266, 151, 137,
	5, 258, 19, 154, 38, 29, 124, 45, 3, 11,
	9, 19, 19, 260, 204, 234, 203, 24, 282, 259,
	23, 151, 195, 290, 176, 22, 21, 126, 62, 134,
	35, 155, 192, 120, 31, 125, 121, 42, 65, 232,
	58, 36, 136, 237, 47, 48, 50, 229, 117, 113,
	238, 61, 60, 120, 10, 342, 292, 265, 312, 330,
	329, 138, 324, 130, 158, 235, 19, 154, 46, 274,
	322, 337, 321, 19, 154, 148, 287, 276, 32, 131,
	141, 147, 115, 188, 191, 151, 128, 188, 188, 199,
	202, 294, 151, 40, 127, 155, 63, 275, 114, 212,
	116, 59, 155, 120, 157, 210, 54, 200, 200, 212,
	215, 201, 222, 212, 212, 212, 206, 207, 54, 129,
	292, 217, 205, 140, 295, 118, 214, 225, 158, 271,
	252, 27, 15, 224, 254, 158, 343, 41, 320, 148,
	242, 240, 120, 227, 245, 147, 148, 248, 250, 314,
	257, 251, 147, 248, 247, 230, 239, 62, 307, 301,
	30, 267, 255, 246, 188, 188, 262, 264, 157, 25,
	183, 123, 249, 122, 7, 157, 310, 256, 222, 222,
	19, 302, 195, 283, 241, 64, 28, 261, 263, 19,
	18, 336, 339, 219, 12, 313, 96, 311, 272, 279,
	226, 14, 18, 291, 268, 270, 278, 14, 26, 188,
	19, 297, 191, 280, 285, 284, 286, 4, 188, 277,
	199, 304, 300, 47, 48, 296, 55, 48, 236, 299,
	298, 233, 44, 303, 53, 44, 334, 332, 319, 338,
	189, 309, 39, 34, 305, 308, 53, 1, 185, 184,
	72, 139, 291, 213, 44, 245, 43, 57, 56, 44,
	37, 243, 315, 69, 76, 318, 317, 316, 82, 135,
	208, 95, 80, 194, 326, 196, 197, 325, 79, 78,
	67, 188, 328, 335, 331, 55, 19, 154, 81, 87,
	86, 71, 181, 288, 179, 177, 172, 175, 188, 188,
	335, 174, 341, 340, 171, 151, 169, 167, 164, 327,
	162, 173, 161, 112, 160, 155, 159, 83, 68, 33,
	49, 51, 52, 153, 152, 143, 146, 77, 89, 144,
	142, 221, 218, 94, 165, 74, 102, 66, 70, 73,
	75, 130, 6, 16, 182, 13, 2, 8, 158, 20,
	84, 0, 90, 106, 98, 0, 103, 0, 85, 148,
	91, 105, 0, 0, 0, 147, 111, 97, 92, 0,
	88, 99, 55, 19, 154, 100, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 107, 0, 0, 157, 0,
	0, 0, 151, 108, 0, 0, 104, 109, 173, 0,
	112, 110, 155, 0, 93, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 89, 0, 0, 0, 0,
	94, 0, 0, 102, 0, 0, 0, 0, 130, 0,
	0, 182, 0, 0, 0, 158, 0, 84, 0, 90,
	106, 98, 0, 103, 0, 85, 148, 91, 105, 55,
	0, 281, 147, 111, 97, 92, 0, 88, 99, 0,
	0, 0, 100, 0, 0, 0, 101, 0, 0, 0,
	0, 0, 107, 0, 0, 157, 0, 112, 0, 282,
	108, 0, 0, 104, 109, 0, 0, 0, 110, 0,
	0, 93, 89, 0, 0, 0, 0, 94, 0, 0,
	102, 0, 0, 0, 0, 0, 0, 0, 55, 0,
	0, 0, 0, 0, 84, 0, 90, 106, 98, 0,
	103, 0, 85, 0, 91, 105, 0, 0, 0, 0,
	111, 97, 92, 0, 88, 99, 112, 0, 0, 100,
	0, 0, 0, 101, 0, 0, 0, 0, 0, 107,
	0, 89, 0, 0, 0, 0, 94, 108, 0, 102,
	104, 109, 0, 0, 0, 110, 133, 0, 93, 0,
	212, 0, 0, 84, 0, 90, 106, 98, 0, 103,
	0, 85, 0, 91, 105, 55, 19, 0, 0, 111,
	97, 92, 0, 88, 99, 0, 0, 0, 100, 0,
	0, 0, 101, 0, 0, 0, 0, 0, 107, 0,
	132, 0, 0, 112, 0, 0, 108, 0, 0, 104,
	109, 0, 0, 0, 110, 0, 0, 93, 89, 0,
	0, 0, 0, 94, 0, 0, 102, 0, 0, 0,
	0, 55, 0, 0, 0, 0, 0, 0, 0, 0,
	84, 0, 90, 106, 98, 0, 103, 0, 85, 0,
	91, 105, 0, 0, 0, 0, 111, 97, 92, 112,
	88, 99, 0, 0, 0, 100, 0, 0, 0, 101,
	0, 0, 0, 0, 89, 107, 0, 0, 0, 94,
	0, 0, 102, 108, 0, 0, 104, 109, 0, 0,
	0, 110, 0, 0, 93, 0, 84, 0, 90, 106,
	98, 0, 103, 0, 85, 0, 91, 105, 0, 0,
	0, 0, 111, 97, 92, 0, 88, 99, 0, 0,
	0, 100, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 107, 0, 0, 0, 0, 0, 0, 0, 108,
	0, 0, 104, 109, 0, 0, 0, 110, 0, 0,
	93,
}
var yyPact = [...]int{

	251, -1000, -79, 188, -1000, -14, -1000, 222, -6, -13,
	-18, -21, 182, 222, -1000, -1000, -1000, 139, -1000, -1000,
	211, -67, -1000, -1000, -1000, -1000, -1000, 234, 7, -1000,
	85, 2, -1000, 29, -72, 78, -1000, 260, 257, 50,
	49, 167, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 260,
	-1000, -1000, -1000, 210, 675, -1000, 47, 257, -1000, 38,
	-1000, -1000, 257, -1000, 675, 150, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 10, -1000, -1000, -1000, 187, 185,
	-1000, -58, 9, -1000, 41, 542, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 3, -20, -1000, -1000, 251, -1000, 111, 106, -1000,
	319, 184, 243, 243, -1000, -1000, 45, 619, -37, -39,
	111, 111, 675, 675, -1000, 27, -1000, -1000, -1000, -1000,
	12, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 26, -1000, -1000, 142, 232, -1000, -1000, -1000, 42,
	-1000, -1000, 165, -1000, -1000, 19, -1000, 11, -1000, 37,
	-1000, 19, -1000, 319, -1000, -1000, -1000, -1000, -1000, 208,
	111, 152, -1000, 243, 176, 164, 163, -1000, 675, 161,
	-1000, 138, -1000, 144, 175, 42, -1000, 160, -1000, -56,
	-40, 111, -1000, 619, 619, -1000, 111, 111, 62, -1000,
	-1000, -1000, -91, -1000, -1000, -1000, 174, 26, 26, -1000,
	-1000, -1000, 137, -1000, 230, 101, 142, -1000, 84, 483,
	206, -1000, 406, 406, -1000, -1000, 406, -1000, -1000, -1000,
	83, 35, -1000, 1, -1000, 132, -1000, 205, 243, 111,
	-1000, 243, 22, 172, 204, -1000, -1000, 44, -1000, 106,
	675, 111, -1000, 111, -1000, -1000, -100, -1000, 171, -1000,
	-1000, 28, 191, -1000, -1000, 229, -1000, -1000, -1000, -1000,
	61, -1000, 227, 159, -1000, -1000, -1000, -1000, -1000, -1000,
	99, -1000, -1000, -1000, 243, 27, 148, -1000, -1000, 79,
	77, -1000, -1000, -1000, -1000, 111, -96, -1000, 69, -1000,
	101, -1000, 106, -1000, 319, -1000, -1000, 67, 66, 144,
	213, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 81, -1000, -1000, -1000, 224, 213, 243, 58,
	-1000, 157, -1000, -1000,
}
var yyPgo = [...]int{

	0, 27, 6, 47, 236, 0, 389, 387, 386, 385,
	172, 234, 383, 382, 233, 9, 380, 379, 378, 377,
	23, 375, 2, 372, 13, 371, 26, 25, 370, 1,
	369, 367, 366, 365, 364, 363, 16, 10, 86, 362,
	361, 360, 359, 108, 358, 357, 31, 356, 354, 352,
	350, 349, 7, 348, 347, 20, 346, 19, 24, 22,
	344, 341, 337, 336, 335, 64, 334, 333, 332, 11,
	331, 330, 329, 328, 320, 319, 318, 15, 316, 315,
	313, 312, 311, 310, 309, 308, 304, 303, 301, 12,
	300, 298, 297, 80, 141, 77, 296, 293, 291, 290,
	289, 3, 288, 287, 283, 282, 280, 18, 17, 4,
	21, 280, 280, 280, 280, 280, 279, 278, 277, 5,
	276, 271, 268, 259,
}
var yyR1 = [...]int{

	0, 103, 4, 3, 43, 37, 5, 8, 13, 13,
	11, 11, 9, 9, 9, 10, 12, 7, 7, 7,
	7, 6, 6, 42, 42, 104, 104, 104, 105, 105,
	90, 90, 91, 91, 92, 92, 93, 98, 97, 97,
	97, 94, 94, 95, 96, 96, 96, 41, 41, 38,
	38, 73, 15, 40, 39, 20, 20, 20, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	74, 74, 22, 29, 28, 28, 28, 28, 18, 33,
	33, 17, 17, 106, 106, 107, 107, 36, 36, 30,
	30, 31, 32, 32, 34, 34, 35, 35, 1, 1,
	1, 1, 2, 2, 87, 87, 88, 88, 89, 89,
	86, 21, 75, 75, 75, 108, 108, 109, 109, 80,
	79, 111, 112, 112, 113, 113, 114, 114, 115, 116,
	116, 78, 78, 77, 77, 77, 77, 99, 100, 100,
	102, 117, 117, 118, 118, 119, 119, 120, 101, 101,
	81, 81, 81, 82, 83, 83, 84, 84, 84, 84,
	76, 76, 16, 27, 27, 26, 26, 23, 23, 23,
	23, 24, 24, 25, 14, 70, 70, 71, 71, 71,
	71, 71, 71, 71, 71, 71, 71, 71, 71, 71,
	72, 85, 44, 44, 45, 45, 45, 45, 46, 47,
	48, 49, 49, 49, 50, 51, 52, 52, 53, 53,
	54, 55, 55, 56, 57, 57, 60, 58, 121, 121,
	122, 122, 59, 59, 63, 63, 63, 63, 61, 62,
	66, 66, 67, 67, 68, 68, 69, 69, 65, 64,
	110, 110, 123, 123, 123,
}
var yyR2 = [...]int{

	0, 8, 1, 1, 1, 1, 1, 2, 3, 0,
	1, 2, 1, 1, 1, 1, 4, 2, 2, 2,
	0, 2, 0, 3, 0, 3, 3, 0, 1, 0,
	3, 0, 1, 0, 1, 2, 3, 2, 1, 1,
	0, 1, 3, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 4, 3, 4, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 4, 1, 3, 4, 4, 1, 2, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 1, 3,
	5, 3, 1, 2, 2, 5, 1, 3, 4, 4,
	2, 1, 3, 5, 4, 1, 2, 2, 0, 1,
	1, 2, 2, 0, 1, 3, 1, 1, 4, 0,
	2, 1, 3, 1, 2, 3, 3, 4, 1, 5,
	1, 2, 0, 1, 3, 1, 1, 4, 1, 3,
	2, 3, 3, 4, 1, 1, 1, 1, 1, 0,
	3, 3, 2, 3, 4, 1, 2, 1, 1, 1,
	1, 1, 1, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 1, 2, 1, 4, 4, 4, 4, 4, 1,
	1, 1, 3, 5, 1, 1, 1, 2, 1, 3,
	1, 1, 3, 1, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 3, 1, 1, 1, 1, 1, 3,
	1, 2, 1, 2, 1, 1, 1, 1, 2, 1,
	2, 0, 1, 1, 3,
}
var yyChk = [...]int{

	-1000, -103, -8, -3, 6, 119, -13, 26, -7, 64,
	108, 63, -11, -9, -14, -10, -12, -5, 8, 7,
	-6, 72, 78, 78, 78, 27, -11, 32, 15, 112,
	-10, 67, 33, -42, -104, 68, 52, -90, 116, -105,
	55, -94, -95, -96, -4, -3, -43, 6, 7, -41,
	-38, -40, -39, -4, -43, 6, -91, -92, -93, -94,
	42, 42, 30, -38, 15, -20, -19, -74, -44, -87,
	-18, -70, -99, -17, -21, -16, -86, -31, -75, -76,
	-81, -73, -85, -45, 71, 79, -71, -72, 91, 49,
	73, 81, 89, 125, 54, -82, -4, 88, 75, 92,
	96, 100, 57, 77, 117, 82, 74, 106, 114, 118,
	122, 87, 34, 42, -93, 84, -95, -20, 15, -46,
	32, 66, 26, 26, 104, 66, 26, 93, -46, -65,
	62, -20, 108, 64, 66, -84, 102, 59, 121, -98,
	-3, -29, -28, -33, -30, -27, -32, 86, 80, -36,
	-5, 26, -34, -35, 8, 36, -1, 109, 69, -47,
	-48, -49, -50, -52, -53, 55, -55, -54, -57, -56,
	-59, -60, -63, 32, -61, -62, -65, -64, -29, -66,
	-20, -68, 65, 26, -100, -102, -101, -22, -5, -106,
	-107, -5, 27, -108, -80, 17, -79, -78, -77, -22,
	103, -20, -22, 93, 93, -46, -20, -20, -83, -37,
	-15, 8, 127, -97, -27, -15, -26, -15, -23, -14,
	-24, -25, -5, 8, 31, 25, 8, -1, -110, 45,
	30, -58, 60, -121, 44, 94, -122, 46, 53, -58,
	-52, 16, 28, -88, -89, -5, 27, 30, 30, -20,
	27, 30, 32, -109, 30, 27, -110, 30, 97, 115,
	93, -20, -22, -20, -22, 35, 128, 27, -26, -15,
	-26, 32, 8, -2, 8, 36, 33, -123, -36, -15,
	-20, 8, 36, 17, -59, -55, -57, 33, -67, -69,
	28, -29, 61, 27, 130, 32, -108, -22, -107, -36,
	-15, 27, 17, -77, -29, -20, 129, 27, -24, -15,
	25, 8, 37, 8, 30, -69, -89, -37, -15, -117,
	30, 33, 33, 128, 33, -2, -29, -51, -52, 33,
	33, -109, -118, -119, -120, -22, 18, 30, -116, 8,
	-119, -101, 37, 19,
}
var yyDef = [...]int{

	0, -2, 0, 9, 3, 20, 7, 0, 22, 0,
	0, 0, 0, 10, 12, 13, 14, 174, 15, 6,
	0, 0, 17, 18, 19, 8, 11, 0, 0, 21,
	0, -2, 16, 0, 31, 29, 1, 0, 33, 0,
	0, 28, 41, 43, 44, 45, 46, 2, 4, 23,
	47, 49, 50, 0, 0, 2, 0, 32, 34, 0,
	25, 26, 0, 48, 0, 0, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 193, 0, 78, 175, 176, 0, 81,
	111, 0, 0, 91, 0, 0, 51, 191, 177, 178,
	179, 180, 181, 182, 183, 184, 185, 186, 187, 188,
	189, 0, 159, 30, 35, 0, 42, 53, 0, 192,
	0, 104, 0, 0, 162, 110, 0, 0, 0, 0,
	0, 150, 0, 0, 190, 0, 156, 157, 158, 36,
	40, 54, 73, 74, 75, 76, 77, 79, 80, 89,
	90, 0, 92, 93, 87, 0, 94, 96, 97, 241,
	199, 200, 201, 204, -2, 0, -2, 0, 211, 0,
	-2, 0, 222, 0, 224, 225, 226, 227, -2, 0,
	239, 230, 235, 0, 0, 138, 140, 148, 0, 0,
	83, 0, 112, 118, 0, 115, 119, 120, 131, 133,
	0, 160, 161, 0, 0, 238, 151, 152, 0, 154,
	155, 5, 0, 37, 38, 39, 0, 170, 165, 167,
	168, 169, 174, 171, 0, 0, 88, 95, 0, 0,
	0, 207, 0, 0, 218, 219, 0, 220, 221, 215,
	0, 0, 231, 0, 106, 0, 137, 0, 0, 72,
	82, 0, 0, 0, 0, 114, 116, 0, 134, 0,
	0, 194, 196, 195, 197, 153, 0, 163, 0, 170,
	166, 0, 99, 101, 102, 0, 198, 240, 242, 243,
	0, 87, 0, 202, 217, -2, 212, 223, 229, 232,
	0, 236, 237, 105, 0, 0, 142, 149, 84, 0,
	0, 113, 117, 132, 135, 136, 0, 164, 0, 172,
	0, 103, 0, 88, 0, 233, 107, 0, 0, 118,
	0, 85, 86, 52, 173, 100, 244, 203, 205, 108,
	109, 139, 141, 143, 145, 146, 129, 0, 0, 0,
	144, 0, 130, 147,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 130, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	129, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 128, 3, 3, 3, 3, 127,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line asn1.y:316
		{
			yylex.(*MyLexer).result = &ModuleDefinition{ModuleIdentifier: yyDollar[1].ModuleIdentifier, TagDefault: yyDollar[3].TagDefault, ExtensibilityImplied: yyDollar[4].ExtensionDefault, ModuleBody: yyDollar[7].ModuleBody}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:319
		{
			yyVAL.TypeReference = TypeReference(yyDollar[1].name)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:324
		{
			yyVAL.ValueReference = ValueReference(yyDollar[1].name)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:335
		{
			yyVAL.ModuleIdentifier = ModuleIdentifier{Reference: yyDollar[1].name, DefinitiveIdentifier: yyDollar[2].DefinitiveIdentifier}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:338
		{
			yyVAL.DefinitiveIdentifier = DefinitiveIdentifier(yyDollar[2].DefinitiveObjIdComponentList)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:339
		{
			yyVAL.DefinitiveIdentifier = DefinitiveIdentifier(make([]DefinitiveObjIdComponent, 0))
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:342
		{
			yyVAL.DefinitiveObjIdComponentList = append(make([]DefinitiveObjIdComponent, 0), yyDollar[1].DefinitiveObjIdComponent)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:343
		{
			yyVAL.DefinitiveObjIdComponentList = append(append(make([]DefinitiveObjIdComponent, 0), yyDollar[1].DefinitiveObjIdComponent), yyDollar[2].DefinitiveObjIdComponentList...)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:346
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Name: yyDollar[1].name}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:347
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Id: yyDollar[1].Number.IntValue()}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:348
		{
			yyVAL.DefinitiveObjIdComponent = yyDollar[1].DefinitiveObjIdComponent
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:351
		{
			yyVAL.Number = yyDollar[1].Number
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:355
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Name: yyDollar[1].name, Id: yyDollar[3].Number.IntValue()}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:358
		{
			yyVAL.TagDefault = TAGS_EXPLICIT
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:359
		{
			yyVAL.TagDefault = TAGS_IMPLICIT
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:360
		{
			yyVAL.TagDefault = TAGS_AUTOMATIC
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:361
		{
			yyVAL.TagDefault = TAGS_EXPLICIT
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:364
		{
			yyVAL.ExtensionDefault = true
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:365
		{
			yyVAL.ExtensionDefault = false
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:368
		{
			yyVAL.ModuleBody = ModuleBody{Imports: yyDollar[2].Imports, AssignmentList: yyDollar[3].AssignmentList}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:369
		{
			yyVAL.ModuleBody = ModuleBody{}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:382
		{
			yyVAL.Imports = yyDollar[2].Imports
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:383
		{
			yyVAL.Imports = make([]SymbolsFromModule, 0)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:386
		{
			yyVAL.Imports = yyDollar[1].Imports
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:387
		{
			yyVAL.Imports = make([]SymbolsFromModule, 0)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:390
		{
			yyVAL.Imports = append(make([]SymbolsFromModule, 0), yyDollar[1].SymbolsFromModule)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:391
		{
			yyVAL.Imports = append(yyDollar[1].Imports, yyDollar[2].SymbolsFromModule)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:394
		{
			yyVAL.SymbolsFromModule = SymbolsFromModule{yyDollar[1].SymbolList, yyDollar[3].GlobalModuleReference}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:397
		{
			yyVAL.GlobalModuleReference = GlobalModuleReference{yyDollar[1].name, yyDollar[2].Value}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:400
		{
			yyVAL.Value = yyDollar[1].ObjectIdentifierValue
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:401
		{
			yyVAL.Value = yyDollar[1].DefinedValue
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:402
		{
			yyVAL.Value = nil
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:405
		{
			yyVAL.SymbolList = append(make([]Symbol, 0), yyDollar[1].Symbol)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:406
		{
			yyVAL.SymbolList = append(yyDollar[1].SymbolList, yyDollar[3].Symbol)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:413
		{
			yyVAL.Symbol = TypeReference(yyDollar[1].TypeReference)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:414
		{
			yyVAL.Symbol = ModuleReference(yyDollar[1].name)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:415
		{
			yyVAL.Symbol = ValueReference(yyDollar[1].ValueReference)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:421
		{
			yyVAL.AssignmentList = NewAssignmentList(yyDollar[1].Assignment)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:422
		{
			yyVAL.AssignmentList = yyDollar[1].AssignmentList.Append(yyDollar[2].Assignment)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:438
		{
			yyVAL.Type = yyDollar[1].TypeReference
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:445
		{
			yyVAL.DefinedValue = DefinedValue{}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:453
		{
			yyVAL.Assignment = TypeAssignment{yyDollar[1].TypeReference, yyDollar[3].Type}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:456
		{
			yyVAL.Assignment = ValueAssignment{yyDollar[1].ValueReference, yyDollar[2].Type, yyDollar[4].Value}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:501
		{
			yyVAL.NamedType = NamedType{Identifier: Identifier(yyDollar[1].name), Type: yyDollar[2].Type}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:524
		{
			yyVAL.Value = yyDollar[1].ObjectIdentifierValue
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:537
		{
			yyVAL.Type = BooleanType{}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:540
		{
			yyVAL.Value = Boolean(true)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:541
		{
			yyVAL.Value = Boolean(false)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:546
		{
			yyVAL.Type = IntegerType{}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:547
		{
			yyVAL.Type = IntegerType{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:558
		{
			yyVAL.Number = yyDollar[1].Number
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:559
		{
			yyVAL.Number = yyDollar[2].Number.UnaryMinus()
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:564
		{
			yyVAL.Value = yyDollar[1].Number
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:565
		{
			yyVAL.Value = IdentifiedIntegerValue{Name: yyDollar[1].name}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:570
		{
			yyVAL.Type = RealType{}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:579
		{
			yyVAL.Value = yyDollar[1].Real
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:580
		{
			yyVAL.Value = yyDollar[2].Real.UnaryMinus()
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:584
		{
			yyVAL.Value = Real(math.Inf(1))
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:585
		{
			yyVAL.Value = Real(math.Inf(-1))
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:589
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, 0, 0)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:590
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, yyDollar[3].Number, 0)
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:591
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, yyDollar[3].Number, yyDollar[5].Number)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:592
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, 0, yyDollar[3].Number)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:596
		{
			yyVAL.Number = Number(-int(yyDollar[2].Number))
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:601
		{
			yyVAL.Type = BitStringType{}
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:602
		{
			yyVAL.Type = BitStringType{NamedBits: yyDollar[4].NamedBitList}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:605
		{
			yyVAL.NamedBitList = append(make([]NamedBit, 0), yyDollar[1].NamedBit)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:606
		{
			yyVAL.NamedBitList = append(yyDollar[1].NamedBitList, yyDollar[3].NamedBit)
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:609
		{
			yyVAL.NamedBit = NamedBit{Name: Identifier(yyDollar[1].name), Index: yyDollar[3].Number}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:610
		{
			yyVAL.NamedBit = NamedBit{Name: Identifier(yyDollar[1].name), Index: yyDollar[3].DefinedValue}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:615
		{
			yyVAL.Type = OctetStringType{}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:620
		{
			yyVAL.Type = NullType{}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:625
		{
			yyVAL.Type = SequenceType{}
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:626
		{
			yyVAL.Type = SequenceType{}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:627
		{
			yyVAL.Type = SequenceType{Components: yyDollar[3].ComponentTypeList}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:669
		{
			yyVAL.ComponentTypeList = append(make(ComponentTypeList, 0), yyDollar[1].ComponentType)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:670
		{
			yyVAL.ComponentTypeList = append(yyDollar[1].ComponentTypeList, yyDollar[3].ComponentType)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:673
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:674
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType, IsOptional: true}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:675
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType, Default: &yyDollar[3].Value}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:676
		{
			yyVAL.ComponentType = ComponentsOfComponentType{Type: yyDollar[3].Type}
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:682
		{
			yyVAL.Type = yyDollar[3].ChoiceType
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:686
		{
			yyVAL.ChoiceType = ChoiceType{yyDollar[1].AlternativeTypeList}
		}
	case 139:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:688
		{
			yyVAL.ChoiceType = ChoiceType{yyDollar[1].AlternativeTypeList}
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:709
		{
			yyVAL.AlternativeTypeList = append(make([]NamedType, 0), yyDollar[1].NamedType)
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:710
		{
			yyVAL.AlternativeTypeList = append(yyDollar[1].AlternativeTypeList, yyDollar[3].NamedType)
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:715
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[2].Type}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:716
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[3].Type, TagType: TAGS_IMPLICIT, HasTagType: true}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:717
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[3].Type, TagType: TAGS_EXPLICIT, HasTagType: true}
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:720
		{
			yyVAL.Tag = Tag{Class: yyDollar[2].Class, ClassNumber: yyDollar[3].Value}
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:723
		{
			yyVAL.Value = yyDollar[1].Number
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:724
		{
			yyVAL.Value = yyDollar[1].DefinedValue
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:727
		{
			yyVAL.Class = CLASS_UNIVERSAL
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:728
		{
			yyVAL.Class = CLASS_APPLICATION
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:729
		{
			yyVAL.Class = CLASS_PRIVATE
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:730
		{
			yyVAL.Class = CLASS_CONTEXT_SPECIFIC
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:735
		{
			yyVAL.Type = SequenceOfType{yyDollar[3].Type}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:736
		{
			yyVAL.Type = SequenceOfType{yyDollar[3].NamedType}
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:741
		{
			yyVAL.Type = ObjectIdentifierType{}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:746
		{
			yyVAL.ObjectIdentifierValue = yyDollar[2].ObjectIdentifierValue
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:747
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[2].DefinedValue).Append(yyDollar[3].ObjectIdentifierValue...)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:750
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[1].ObjIdComponents)
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:751
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[1].ObjIdComponents).Append(yyDollar[2].ObjectIdentifierValue...)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:754
		{
			yyVAL.ObjIdComponents = ObjectIdElement{Name: yyDollar[1].name}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:757
		{
			yyVAL.ObjIdComponents = yyDollar[1].DefinedValue
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:760
		{
			yyVAL.ObjIdComponents = ObjectIdElement{Id: yyDollar[1].Number.IntValue()}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:761
		{
			yyVAL.ObjIdComponents = yyDollar[1].DefinedValue
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:765
		{
			switch v := yyDollar[3].ObjIdComponents.(type) {
			case DefinedValue:
				yyVAL.ObjIdComponents = ObjectIdElement{Name: yyDollar[1].name, Reference: &v}
			case ObjectIdElement:
				yyVAL.ObjIdComponents = ObjectIdElement{Name: yyDollar[1].name, Id: v.Id}
			default:
				panic(fmt.Sprintf("Expected DefinedValue or ObjectIdElement from NumberForm, got %v", yyDollar[3].ObjIdComponents))
			}
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:786
		{
			yyVAL.Type = RestrictedStringType{LexType: BMPString}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:787
		{
			yyVAL.Type = RestrictedStringType{LexType: GeneralString}
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:788
		{
			yyVAL.Type = RestrictedStringType{LexType: GraphicString}
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:789
		{
			yyVAL.Type = RestrictedStringType{LexType: IA5String}
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:790
		{
			yyVAL.Type = RestrictedStringType{LexType: ISO646String}
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:791
		{
			yyVAL.Type = RestrictedStringType{LexType: NumericString}
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:792
		{
			yyVAL.Type = RestrictedStringType{LexType: PrintableString}
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:793
		{
			yyVAL.Type = RestrictedStringType{LexType: TeletexString}
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:794
		{
			yyVAL.Type = RestrictedStringType{LexType: T61String}
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:795
		{
			yyVAL.Type = RestrictedStringType{LexType: UniversalString}
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:796
		{
			yyVAL.Type = RestrictedStringType{LexType: UTF8String}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:797
		{
			yyVAL.Type = RestrictedStringType{LexType: VideotexString}
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:798
		{
			yyVAL.Type = RestrictedStringType{LexType: VisibleString}
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:803
		{
			yyVAL.Type = CharacterStringType{}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:808
		{
			yyVAL.Type = TypeReference("GeneralizedTime")
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:813
		{
			yyVAL.Type = ConstraintedType{yyDollar[1].Type, yyDollar[2].Constraint}
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:821
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].Type}, yyDollar[2].Constraint}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:822
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].Type}, SingleElementConstraint(yyDollar[2].Elements)}
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:825
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].NamedType}, yyDollar[2].Constraint}
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:826
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].NamedType}, SingleElementConstraint(yyDollar[2].Elements)}
		}
	case 198:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:831
		{
			yyVAL.Constraint = Constraint{ConstraintSpec: yyDollar[2].ConstraintSpec}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:834
		{
			yyVAL.ConstraintSpec = yyDollar[1].SubtypeConstraint
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:844
		{
			yyVAL.SubtypeConstraint = yyDollar[1].SubtypeConstraint
		}
	case 203:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:845
		{
			yyVAL.SubtypeConstraint = append(yyDollar[1].SubtypeConstraint, yyDollar[5].ElementSetSpec)
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:848
		{
			yyVAL.SubtypeConstraint = SubtypeConstraint{yyDollar[1].ElementSetSpec}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:854
		{
			yyVAL.ElementSetSpec = yyDollar[1].Unions
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:855
		{
			yyVAL.ElementSetSpec = yyDollar[2].Exclusions
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:858
		{
			yyVAL.Unions = Unions{yyDollar[1].Intersections}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:859
		{
			yyVAL.Unions = append(yyDollar[1].Unions, yyDollar[3].Intersections)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:865
		{
			yyVAL.Intersections = Intersections{yyDollar[1].IntersectionElements}
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:866
		{
			yyVAL.Intersections = append(yyDollar[1].Intersections, yyDollar[3].IntersectionElements)
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:872
		{
			yyVAL.IntersectionElements = IntersectionElements{Elements: yyDollar[1].Elements}
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:873
		{
			yyVAL.IntersectionElements = IntersectionElements{Elements: yyDollar[1].Elements, Exclusions: yyDollar[2].Exclusions}
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:879
		{
			yyVAL.Exclusions = Exclusions{yyDollar[2].Elements}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:888
		{
			yyVAL.Elements = yyDollar[1].Elements
		}
	case 223:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:890
		{
			yyVAL.Elements = yyDollar[2].ElementSetSpec
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:905
		{
			yyVAL.Elements = SingleValue{yyDollar[1].Value}
		}
	case 229:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:910
		{
			yyVAL.Elements = ValueRange{yyDollar[1].RangeEndpoint, yyDollar[3].RangeEndpoint}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:913
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:914
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value, IsOpen: true}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:917
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value}
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:918
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[2].Value, IsOpen: true}
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:922
		{
			yyVAL.Value = nil
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:926
		{
			yyVAL.Value = nil
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:931
		{
			yyVAL.Elements = SizeConstraint{yyDollar[2].Constraint}
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:936
		{
			yyVAL.Elements = TypeConstraint{yyDollar[1].Type}
		}
	}
	goto yystack /* stack new state and value */
}
