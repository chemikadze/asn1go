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

	Number                            Number
	Real                              Real
	TagDefault                        int
	ExtensionDefault                  bool
	ModuleIdentifier                  ModuleIdentifier
	DefinitiveObjIdComponent          DefinitiveObjIdComponent
	DefinitiveObjIdComponentList      []DefinitiveObjIdComponent
	DefinitiveIdentifier              DefinitiveIdentifier
	Type                              Type
	ObjIdComponents                   ObjIdComponents
	DefinedValue                      DefinedValue
	ObjectIdentifierValue             ObjectIdentifierValue
	Value                             Value
	Assignment                        Assignment
	AssignmentList                    AssignmentList
	ModuleBody                        ModuleBody
	ValueReference                    ValueReference
	TypeReference                     TypeReference
	Constraint                        Constraint
	ConstraintSpec                    ConstraintSpec
	ElementSetSpec                    ElementSetSpec
	Unions                            Unions
	Intersections                     Intersections
	IntersectionElements              IntersectionElements
	Exclusions                        Exclusions
	Elements                          Elements
	SubtypeConstraint                 SubtypeConstraint
	RangeEndpoint                     RangeEndpoint
	NamedType                         NamedType
	ComponentType                     ComponentType
	ComponentTypeList                 ComponentTypeList
	SequenceType                      SequenceType
	Tag                               Tag
	Class                             int
	SequenceOfType                    SequenceOfType
	NamedBitList                      []NamedBit
	NamedBit                          NamedBit
	Imports                           []SymbolsFromModule
	SymbolsFromModule                 SymbolsFromModule
	SymbolList                        []Symbol
	Symbol                            Symbol
	GlobalModuleReference             GlobalModuleReference
	AlternativeTypeList               []NamedType
	ChoiceType                        ChoiceType
	ExtensionAdditionAlternative      ChoiceExtension
	ExtensionAdditionAlternativesList []ChoiceExtension
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

//line asn1.y:969

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 31,
	52, 24,
	-2, 27,
	-1, 164,
	44, 209,
	94, 209,
	-2, 205,
	-1, 166,
	46, 212,
	53, 212,
	-2, 207,
	-1, 170,
	60, 215,
	-2, 213,
	-1, 178,
	16, 233,
	28, 233,
	-2, 227,
	-1, 283,
	46, 212,
	53, 212,
	-2, 208,
}

const yyPrivate = 57344

const yyLast = 787

var yyAct = [...]int{

	150, 331, 198, 178, 251, 163, 271, 267, 17, 208,
	287, 243, 219, 197, 17, 189, 149, 192, 168, 166,
	215, 170, 180, 230, 227, 156, 145, 304, 291, 321,
	279, 119, 210, 19, 222, 222, 264, 151, 137, 5,
	19, 154, 45, 3, 256, 38, 29, 11, 9, 19,
	124, 258, 176, 233, 203, 202, 24, 35, 280, 151,
	62, 288, 257, 23, 19, 126, 22, 19, 154, 155,
	21, 120, 31, 134, 194, 231, 125, 65, 121, 58,
	36, 136, 236, 228, 191, 50, 151, 117, 113, 237,
	61, 60, 10, 263, 290, 328, 155, 327, 42, 120,
	138, 130, 158, 234, 310, 59, 272, 322, 47, 48,
	320, 319, 285, 148, 115, 274, 32, 120, 131, 147,
	293, 290, 141, 187, 190, 186, 128, 187, 187, 158,
	201, 292, 127, 46, 273, 63, 269, 114, 211, 96,
	148, 41, 157, 209, 250, 199, 147, 129, 214, 211,
	200, 211, 221, 211, 211, 205, 206, 40, 140, 216,
	199, 116, 204, 118, 224, 27, 333, 213, 252, 157,
	223, 54, 318, 248, 312, 44, 249, 53, 44, 239,
	120, 226, 255, 54, 244, 246, 229, 62, 241, 53,
	305, 299, 15, 19, 154, 238, 265, 44, 253, 245,
	25, 183, 44, 187, 187, 260, 262, 123, 122, 7,
	247, 19, 151, 308, 300, 281, 240, 221, 221, 254,
	30, 194, 155, 64, 28, 259, 261, 218, 12, 19,
	18, 311, 309, 270, 225, 14, 277, 266, 268, 47,
	48, 14, 26, 18, 289, 276, 19, 187, 4, 295,
	190, 278, 283, 282, 284, 158, 187, 275, 298, 55,
	48, 302, 235, 232, 294, 296, 148, 297, 188, 301,
	39, 34, 147, 1, 330, 317, 185, 307, 184, 72,
	139, 303, 306, 212, 43, 57, 56, 37, 242, 69,
	76, 82, 289, 244, 135, 157, 207, 95, 80, 313,
	193, 316, 195, 315, 314, 196, 79, 78, 67, 81,
	87, 55, 19, 154, 324, 323, 86, 71, 326, 187,
	181, 332, 329, 286, 179, 177, 172, 175, 174, 171,
	169, 151, 167, 164, 187, 334, 332, 173, 325, 112,
	162, 155, 161, 160, 159, 83, 68, 33, 49, 51,
	52, 153, 152, 143, 89, 146, 77, 144, 142, 94,
	165, 220, 102, 217, 74, 66, 70, 130, 73, 75,
	182, 6, 16, 13, 158, 2, 84, 8, 90, 106,
	98, 20, 103, 0, 85, 148, 91, 105, 0, 0,
	0, 147, 111, 97, 92, 0, 88, 99, 55, 19,
	154, 100, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 107, 0, 0, 157, 0, 0, 0, 151, 108,
	0, 0, 104, 109, 173, 0, 112, 110, 155, 0,
	93, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 89, 0, 0, 0, 0, 94, 0, 0, 102,
	0, 0, 0, 0, 130, 0, 0, 182, 0, 0,
	0, 158, 0, 84, 0, 90, 106, 98, 0, 103,
	0, 85, 148, 91, 105, 55, 0, 279, 147, 111,
	97, 92, 0, 88, 99, 0, 0, 0, 100, 0,
	0, 0, 101, 0, 0, 0, 0, 0, 107, 0,
	0, 157, 0, 112, 0, 280, 108, 0, 0, 104,
	109, 0, 0, 0, 110, 0, 0, 93, 89, 0,
	0, 0, 0, 94, 0, 0, 102, 0, 0, 0,
	0, 0, 0, 0, 55, 0, 0, 0, 0, 0,
	84, 0, 90, 106, 98, 0, 103, 0, 85, 0,
	91, 105, 0, 0, 0, 0, 111, 97, 92, 0,
	88, 99, 112, 0, 0, 100, 0, 0, 0, 101,
	0, 0, 0, 0, 0, 107, 0, 89, 0, 0,
	0, 0, 94, 108, 0, 102, 104, 109, 0, 0,
	0, 110, 133, 0, 93, 0, 211, 0, 0, 84,
	0, 90, 106, 98, 0, 103, 0, 85, 0, 91,
	105, 55, 19, 0, 0, 111, 97, 92, 0, 88,
	99, 0, 0, 0, 100, 0, 0, 0, 101, 0,
	0, 0, 0, 0, 107, 0, 132, 0, 0, 112,
	0, 0, 108, 0, 0, 104, 109, 0, 0, 0,
	110, 0, 0, 93, 89, 0, 0, 0, 0, 94,
	0, 0, 102, 0, 0, 0, 0, 55, 0, 0,
	0, 0, 0, 0, 0, 0, 84, 0, 90, 106,
	98, 0, 103, 0, 85, 0, 91, 105, 0, 0,
	0, 0, 111, 97, 92, 112, 88, 99, 0, 0,
	0, 100, 0, 0, 0, 101, 0, 0, 0, 0,
	89, 107, 0, 0, 0, 94, 0, 0, 102, 108,
	0, 0, 104, 109, 0, 0, 0, 110, 0, 0,
	93, 0, 84, 0, 90, 106, 98, 0, 103, 0,
	85, 0, 91, 105, 0, 0, 0, 0, 111, 97,
	92, 0, 88, 99, 0, 0, 0, 100, 0, 0,
	0, 101, 0, 0, 0, 0, 0, 107, 0, 0,
	0, 0, 0, 0, 0, 108, 0, 0, 104, 109,
	0, 0, 0, 110, 0, 0, 93,
}
var yyPact = [...]int{

	242, -1000, -80, 183, -1000, -16, -1000, 222, -2, -12,
	-15, -22, 173, 222, -1000, -1000, -1000, 133, -1000, -1000,
	209, -66, -1000, -1000, -1000, -1000, -1000, 235, 5, -1000,
	83, -11, -1000, 28, -71, 102, -1000, 253, 233, 49,
	48, 157, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 253,
	-1000, -1000, -1000, 208, 661, -1000, 46, 233, -1000, 30,
	-1000, -1000, 233, -1000, 661, 148, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 12, -1000, -1000, -1000, 182, 181,
	-1000, -54, 10, -1000, 39, 528, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 7, -21, -1000, -1000, 242, -1000, 85, 186, -1000,
	305, 175, 239, 239, -1000, -1000, 57, 605, -38, -39,
	85, 85, 661, 661, -1000, 24, -1000, -1000, -1000, -1000,
	11, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 26, -1000, -1000, 139, 226, -1000, -1000, -1000, 38,
	-1000, -1000, 156, -1000, -1000, 15, -1000, 9, -1000, 36,
	-1000, 15, -1000, 305, -1000, -1000, -1000, -1000, -1000, 200,
	85, 160, -1000, 239, 172, 155, -1000, 661, 146, -1000,
	112, -1000, 138, 171, 38, -1000, 152, -1000, -53, -42,
	85, -1000, 605, 605, -1000, 85, 85, 58, -1000, -1000,
	-1000, -92, -1000, -1000, -1000, 169, 26, 26, -1000, -1000,
	-1000, 104, -1000, 225, 98, 139, -1000, 82, 469, 198,
	-1000, 392, 392, -1000, -1000, 392, -1000, -1000, -1000, 79,
	33, -1000, 1, -1000, 88, -1000, 204, 85, -1000, 239,
	22, 164, 197, -1000, -1000, 42, -1000, 186, 661, 85,
	-1000, 85, -1000, -1000, -102, -1000, 163, -1000, -1000, 27,
	188, -1000, -1000, 224, -1000, -1000, -1000, -1000, 67, -1000,
	223, 144, -1000, -1000, -1000, -1000, -1000, -1000, 60, -1000,
	-1000, -1000, 239, 24, 142, -1000, -1000, 78, 77, -1000,
	-1000, -1000, -1000, 85, -99, -1000, 74, -1000, 98, -1000,
	186, -1000, 305, -1000, -1000, 64, 62, 138, 239, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	136, -1000, -1000, 239, -1000,
}
var yyPgo = [...]int{

	0, 25, 6, 42, 139, 0, 381, 377, 375, 373,
	192, 228, 372, 371, 227, 7, 369, 368, 366, 365,
	22, 364, 2, 363, 12, 361, 20, 26, 358, 3,
	357, 356, 355, 353, 352, 351, 16, 9, 85, 350,
	349, 348, 347, 133, 346, 345, 31, 344, 343, 342,
	340, 338, 5, 333, 332, 19, 330, 18, 23, 21,
	329, 328, 327, 326, 325, 52, 324, 323, 320, 10,
	317, 316, 310, 309, 308, 307, 306, 13, 305, 302,
	300, 298, 297, 296, 294, 291, 290, 289, 288, 11,
	287, 286, 285, 79, 105, 98, 284, 283, 280, 279,
	278, 276, 276, 1, 275, 274, 273, 271, 270, 268,
	15, 17, 4, 24, 268, 268, 268, 268, 268, 268,
	268, 263, 262, 257,
}
var yyR1 = [...]int{

	0, 106, 4, 3, 43, 37, 5, 8, 13, 13,
	11, 11, 9, 9, 9, 10, 12, 7, 7, 7,
	7, 6, 6, 42, 42, 107, 107, 107, 108, 108,
	90, 90, 91, 91, 92, 92, 93, 98, 97, 97,
	97, 94, 94, 95, 96, 96, 96, 41, 41, 38,
	38, 73, 15, 40, 39, 20, 20, 20, 19, 19,
	19, 19, 19, 19, 19, 19, 19, 19, 19, 19,
	74, 74, 22, 29, 28, 28, 28, 28, 18, 33,
	33, 17, 17, 109, 109, 110, 110, 36, 36, 30,
	30, 31, 32, 32, 34, 34, 35, 35, 1, 1,
	1, 1, 2, 2, 87, 87, 88, 88, 89, 89,
	86, 21, 75, 75, 75, 111, 111, 112, 112, 80,
	79, 114, 115, 115, 116, 116, 117, 117, 118, 119,
	119, 78, 78, 77, 77, 77, 77, 99, 100, 100,
	102, 104, 104, 105, 105, 103, 120, 101, 101, 81,
	81, 81, 82, 83, 83, 84, 84, 84, 84, 76,
	76, 16, 27, 27, 26, 26, 23, 23, 23, 23,
	24, 24, 25, 14, 70, 70, 71, 71, 71, 71,
	71, 71, 71, 71, 71, 71, 71, 71, 71, 72,
	85, 44, 44, 45, 45, 45, 45, 46, 47, 48,
	49, 49, 49, 50, 51, 52, 52, 53, 53, 54,
	55, 55, 56, 57, 57, 60, 58, 121, 121, 122,
	122, 59, 59, 63, 63, 63, 63, 61, 62, 66,
	66, 67, 67, 68, 68, 69, 69, 65, 64, 113,
	113, 123, 123, 123,
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
	2, 1, 3, 1, 2, 3, 3, 4, 5, 1,
	1, 2, 0, 1, 3, 1, 4, 1, 3, 2,
	3, 3, 4, 1, 1, 1, 1, 1, 0, 3,
	3, 2, 3, 4, 1, 2, 1, 1, 1, 1,
	1, 1, 4, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 2, 1, 4, 4, 4, 4, 4, 1, 1,
	1, 3, 5, 1, 1, 1, 2, 1, 3, 1,
	1, 3, 1, 1, 2, 1, 2, 1, 1, 1,
	1, 1, 3, 1, 1, 1, 1, 1, 3, 1,
	2, 1, 2, 1, 1, 1, 1, 2, 1, 2,
	0, 1, 1, 3,
}
var yyChk = [...]int{

	-1000, -106, -8, -3, 6, 119, -13, 26, -7, 64,
	108, 63, -11, -9, -14, -10, -12, -5, 8, 7,
	-6, 72, 78, 78, 78, 27, -11, 32, 15, 112,
	-10, 67, 33, -42, -107, 68, 52, -90, 116, -108,
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
	-20, -68, 65, 26, -100, -101, -22, -5, -109, -110,
	-5, 27, -111, -80, 17, -79, -78, -77, -22, 103,
	-20, -22, 93, 93, -46, -20, -20, -83, -37, -15,
	8, 127, -97, -27, -15, -26, -15, -23, -14, -24,
	-25, -5, 8, 31, 25, 8, -1, -113, 45, 30,
	-58, 60, -121, 44, 94, -122, 46, 53, -58, -52,
	16, 28, -88, -89, -5, 27, 30, -20, 27, 30,
	32, -112, 30, 27, -113, 30, 97, 115, 93, -20,
	-22, -20, -22, 35, 128, 27, -26, -15, -26, 32,
	8, -2, 8, 36, 33, -123, -36, -15, -20, 8,
	36, 17, -59, -55, -57, 33, -67, -69, 28, -29,
	61, 27, 130, 32, -111, -22, -110, -36, -15, 27,
	17, -77, -29, -20, 129, 27, -24, -15, 25, 8,
	37, 8, 30, -69, -89, -37, -15, -104, 30, 33,
	33, 128, 33, -2, -29, -51, -52, 33, 33, -112,
	-105, -103, -22, 30, -103,
}
var yyDef = [...]int{

	0, -2, 0, 9, 3, 20, 7, 0, 22, 0,
	0, 0, 0, 10, 12, 13, 14, 173, 15, 6,
	0, 0, 17, 18, 19, 8, 11, 0, 0, 21,
	0, -2, 16, 0, 31, 29, 1, 0, 33, 0,
	0, 28, 41, 43, 44, 45, 46, 2, 4, 23,
	47, 49, 50, 0, 0, 2, 0, 32, 34, 0,
	25, 26, 0, 48, 0, 0, 55, 56, 57, 58,
	59, 60, 61, 62, 63, 64, 65, 66, 67, 68,
	69, 70, 71, 192, 0, 78, 174, 175, 0, 81,
	111, 0, 0, 91, 0, 0, 51, 190, 176, 177,
	178, 179, 180, 181, 182, 183, 184, 185, 186, 187,
	188, 0, 158, 30, 35, 0, 42, 53, 0, 191,
	0, 104, 0, 0, 161, 110, 0, 0, 0, 0,
	0, 149, 0, 0, 189, 0, 155, 156, 157, 36,
	40, 54, 73, 74, 75, 76, 77, 79, 80, 89,
	90, 0, 92, 93, 87, 0, 94, 96, 97, 240,
	198, 199, 200, 203, -2, 0, -2, 0, 210, 0,
	-2, 0, 221, 0, 223, 224, 225, 226, -2, 0,
	238, 229, 234, 0, 0, 139, 147, 0, 0, 83,
	0, 112, 118, 0, 115, 119, 120, 131, 133, 0,
	159, 160, 0, 0, 237, 150, 151, 0, 153, 154,
	5, 0, 37, 38, 39, 0, 169, 164, 166, 167,
	168, 173, 170, 0, 0, 88, 95, 0, 0, 0,
	206, 0, 0, 217, 218, 0, 219, 220, 214, 0,
	0, 230, 0, 106, 0, 137, 0, 72, 82, 0,
	0, 0, 0, 114, 116, 0, 134, 0, 0, 193,
	195, 194, 196, 152, 0, 162, 0, 169, 165, 0,
	99, 101, 102, 0, 197, 239, 241, 242, 0, 87,
	0, 201, 216, -2, 211, 222, 228, 231, 0, 235,
	236, 105, 0, 0, 142, 148, 84, 0, 0, 113,
	117, 132, 135, 136, 0, 163, 0, 171, 0, 103,
	0, 88, 0, 232, 107, 0, 0, 118, 0, 85,
	86, 52, 172, 100, 243, 202, 204, 108, 109, 138,
	141, 143, 145, 0, 144,
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
		//line asn1.y:321
		{
			yylex.(*MyLexer).result = &ModuleDefinition{ModuleIdentifier: yyDollar[1].ModuleIdentifier, TagDefault: yyDollar[3].TagDefault, ExtensibilityImplied: yyDollar[4].ExtensionDefault, ModuleBody: yyDollar[7].ModuleBody}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:324
		{
			yyVAL.TypeReference = TypeReference(yyDollar[1].name)
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:329
		{
			yyVAL.ValueReference = ValueReference(yyDollar[1].name)
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:340
		{
			yyVAL.ModuleIdentifier = ModuleIdentifier{Reference: yyDollar[1].name, DefinitiveIdentifier: yyDollar[2].DefinitiveIdentifier}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:343
		{
			yyVAL.DefinitiveIdentifier = DefinitiveIdentifier(yyDollar[2].DefinitiveObjIdComponentList)
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:344
		{
			yyVAL.DefinitiveIdentifier = DefinitiveIdentifier(make([]DefinitiveObjIdComponent, 0))
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:347
		{
			yyVAL.DefinitiveObjIdComponentList = append(make([]DefinitiveObjIdComponent, 0), yyDollar[1].DefinitiveObjIdComponent)
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:348
		{
			yyVAL.DefinitiveObjIdComponentList = append(append(make([]DefinitiveObjIdComponent, 0), yyDollar[1].DefinitiveObjIdComponent), yyDollar[2].DefinitiveObjIdComponentList...)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:351
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Name: yyDollar[1].name}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:352
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Id: yyDollar[1].Number.IntValue()}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:353
		{
			yyVAL.DefinitiveObjIdComponent = yyDollar[1].DefinitiveObjIdComponent
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:356
		{
			yyVAL.Number = yyDollar[1].Number
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:360
		{
			yyVAL.DefinitiveObjIdComponent = DefinitiveObjIdComponent{Name: yyDollar[1].name, Id: yyDollar[3].Number.IntValue()}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:363
		{
			yyVAL.TagDefault = TAGS_EXPLICIT
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:364
		{
			yyVAL.TagDefault = TAGS_IMPLICIT
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:365
		{
			yyVAL.TagDefault = TAGS_AUTOMATIC
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:366
		{
			yyVAL.TagDefault = TAGS_EXPLICIT
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:369
		{
			yyVAL.ExtensionDefault = true
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:370
		{
			yyVAL.ExtensionDefault = false
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:373
		{
			yyVAL.ModuleBody = ModuleBody{Imports: yyDollar[2].Imports, AssignmentList: yyDollar[3].AssignmentList}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:374
		{
			yyVAL.ModuleBody = ModuleBody{}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:387
		{
			yyVAL.Imports = yyDollar[2].Imports
		}
	case 31:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:388
		{
			yyVAL.Imports = make([]SymbolsFromModule, 0)
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:391
		{
			yyVAL.Imports = yyDollar[1].Imports
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:392
		{
			yyVAL.Imports = make([]SymbolsFromModule, 0)
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:395
		{
			yyVAL.Imports = append(make([]SymbolsFromModule, 0), yyDollar[1].SymbolsFromModule)
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:396
		{
			yyVAL.Imports = append(yyDollar[1].Imports, yyDollar[2].SymbolsFromModule)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:399
		{
			yyVAL.SymbolsFromModule = SymbolsFromModule{yyDollar[1].SymbolList, yyDollar[3].GlobalModuleReference}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:402
		{
			yyVAL.GlobalModuleReference = GlobalModuleReference{yyDollar[1].name, yyDollar[2].Value}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:405
		{
			yyVAL.Value = yyDollar[1].ObjectIdentifierValue
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:406
		{
			yyVAL.Value = yyDollar[1].DefinedValue
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:407
		{
			yyVAL.Value = nil
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:410
		{
			yyVAL.SymbolList = append(make([]Symbol, 0), yyDollar[1].Symbol)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:411
		{
			yyVAL.SymbolList = append(yyDollar[1].SymbolList, yyDollar[3].Symbol)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:418
		{
			yyVAL.Symbol = TypeReference(yyDollar[1].TypeReference)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:419
		{
			yyVAL.Symbol = ModuleReference(yyDollar[1].name)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:420
		{
			yyVAL.Symbol = ValueReference(yyDollar[1].ValueReference)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:426
		{
			yyVAL.AssignmentList = NewAssignmentList(yyDollar[1].Assignment)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:427
		{
			yyVAL.AssignmentList = yyDollar[1].AssignmentList.Append(yyDollar[2].Assignment)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:443
		{
			yyVAL.Type = yyDollar[1].TypeReference
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:450
		{
			yyVAL.DefinedValue = DefinedValue{}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:458
		{
			yyVAL.Assignment = TypeAssignment{yyDollar[1].TypeReference, yyDollar[3].Type}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:461
		{
			yyVAL.Assignment = ValueAssignment{yyDollar[1].ValueReference, yyDollar[2].Type, yyDollar[4].Value}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:506
		{
			yyVAL.NamedType = NamedType{Identifier: Identifier(yyDollar[1].name), Type: yyDollar[2].Type}
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:529
		{
			yyVAL.Value = yyDollar[1].ObjectIdentifierValue
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:542
		{
			yyVAL.Type = BooleanType{}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:545
		{
			yyVAL.Value = Boolean(true)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:546
		{
			yyVAL.Value = Boolean(false)
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:551
		{
			yyVAL.Type = IntegerType{}
		}
	case 82:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:552
		{
			yyVAL.Type = IntegerType{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:563
		{
			yyVAL.Number = yyDollar[1].Number
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:564
		{
			yyVAL.Number = yyDollar[2].Number.UnaryMinus()
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:569
		{
			yyVAL.Value = yyDollar[1].Number
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:570
		{
			yyVAL.Value = IdentifiedIntegerValue{Name: yyDollar[1].name}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:575
		{
			yyVAL.Type = RealType{}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:584
		{
			yyVAL.Value = yyDollar[1].Real
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:585
		{
			yyVAL.Value = yyDollar[2].Real.UnaryMinus()
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:589
		{
			yyVAL.Value = Real(math.Inf(1))
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:590
		{
			yyVAL.Value = Real(math.Inf(-1))
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:594
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, 0, 0)
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:595
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, yyDollar[3].Number, 0)
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:596
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, yyDollar[3].Number, yyDollar[5].Number)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:597
		{
			yyVAL.Real = parseRealNumber(yyDollar[1].Number, 0, yyDollar[3].Number)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:601
		{
			yyVAL.Number = Number(-int(yyDollar[2].Number))
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:606
		{
			yyVAL.Type = BitStringType{}
		}
	case 105:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:607
		{
			yyVAL.Type = BitStringType{NamedBits: yyDollar[4].NamedBitList}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:610
		{
			yyVAL.NamedBitList = append(make([]NamedBit, 0), yyDollar[1].NamedBit)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:611
		{
			yyVAL.NamedBitList = append(yyDollar[1].NamedBitList, yyDollar[3].NamedBit)
		}
	case 108:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:614
		{
			yyVAL.NamedBit = NamedBit{Name: Identifier(yyDollar[1].name), Index: yyDollar[3].Number}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:615
		{
			yyVAL.NamedBit = NamedBit{Name: Identifier(yyDollar[1].name), Index: yyDollar[3].DefinedValue}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:620
		{
			yyVAL.Type = OctetStringType{}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:625
		{
			yyVAL.Type = NullType{}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:630
		{
			yyVAL.Type = SequenceType{}
		}
	case 113:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:631
		{
			yyVAL.Type = SequenceType{}
		}
	case 114:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:632
		{
			yyVAL.Type = SequenceType{Components: yyDollar[3].ComponentTypeList}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:674
		{
			yyVAL.ComponentTypeList = append(make(ComponentTypeList, 0), yyDollar[1].ComponentType)
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:675
		{
			yyVAL.ComponentTypeList = append(yyDollar[1].ComponentTypeList, yyDollar[3].ComponentType)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:678
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:679
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType, IsOptional: true}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:680
		{
			yyVAL.ComponentType = NamedComponentType{NamedType: yyDollar[1].NamedType, Default: &yyDollar[3].Value}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:681
		{
			yyVAL.ComponentType = ComponentsOfComponentType{Type: yyDollar[3].Type}
		}
	case 137:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:687
		{
			yyVAL.Type = yyDollar[3].ChoiceType
		}
	case 138:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:691
		{
			yyVAL.ChoiceType = ChoiceType{yyDollar[1].AlternativeTypeList, yyDollar[4].ExtensionAdditionAlternativesList}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:692
		{
			yyVAL.ChoiceType = ChoiceType{AlternativeTypeList: yyDollar[1].AlternativeTypeList}
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:699
		{
			yyVAL.ExtensionAdditionAlternativesList = yyDollar[2].ExtensionAdditionAlternativesList
		}
	case 142:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:700
		{
			yyVAL.ExtensionAdditionAlternativesList = make([]ChoiceExtension, 0)
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:703
		{
			yyVAL.ExtensionAdditionAlternativesList = append(make([]ChoiceExtension, 0), yyDollar[1].ExtensionAdditionAlternative)
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:704
		{
			yyVAL.ExtensionAdditionAlternativesList = append(yyDollar[1].ExtensionAdditionAlternativesList, yyDollar[3].ExtensionAdditionAlternative)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:708
		{
			yyVAL.ExtensionAdditionAlternative = yyDollar[1].NamedType
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:714
		{
			yyVAL.AlternativeTypeList = append(make([]NamedType, 0), yyDollar[1].NamedType)
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:715
		{
			yyVAL.AlternativeTypeList = append(yyDollar[1].AlternativeTypeList, yyDollar[3].NamedType)
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:720
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[2].Type}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:721
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[3].Type, TagType: TAGS_IMPLICIT, HasTagType: true}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:722
		{
			yyVAL.Type = TaggedType{Tag: yyDollar[1].Tag, Type: yyDollar[3].Type, TagType: TAGS_EXPLICIT, HasTagType: true}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:725
		{
			yyVAL.Tag = Tag{Class: yyDollar[2].Class, ClassNumber: yyDollar[3].Value}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:728
		{
			yyVAL.Value = yyDollar[1].Number
		}
	case 154:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:729
		{
			yyVAL.Value = yyDollar[1].DefinedValue
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:732
		{
			yyVAL.Class = CLASS_UNIVERSAL
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:733
		{
			yyVAL.Class = CLASS_APPLICATION
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:734
		{
			yyVAL.Class = CLASS_PRIVATE
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line asn1.y:735
		{
			yyVAL.Class = CLASS_CONTEXT_SPECIFIC
		}
	case 159:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:740
		{
			yyVAL.Type = SequenceOfType{yyDollar[3].Type}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:741
		{
			yyVAL.Type = SequenceOfType{yyDollar[3].NamedType}
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:746
		{
			yyVAL.Type = ObjectIdentifierType{}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:751
		{
			yyVAL.ObjectIdentifierValue = yyDollar[2].ObjectIdentifierValue
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:752
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[2].DefinedValue).Append(yyDollar[3].ObjectIdentifierValue...)
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:755
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[1].ObjIdComponents)
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:756
		{
			yyVAL.ObjectIdentifierValue = NewObjectIdentifierValue(yyDollar[1].ObjIdComponents).Append(yyDollar[2].ObjectIdentifierValue...)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:759
		{
			yyVAL.ObjIdComponents = ObjectIdElement{Name: yyDollar[1].name}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:762
		{
			yyVAL.ObjIdComponents = yyDollar[1].DefinedValue
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:765
		{
			yyVAL.ObjIdComponents = ObjectIdElement{Id: yyDollar[1].Number.IntValue()}
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:766
		{
			yyVAL.ObjIdComponents = yyDollar[1].DefinedValue
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:770
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
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:791
		{
			yyVAL.Type = RestrictedStringType{LexType: BMPString}
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:792
		{
			yyVAL.Type = RestrictedStringType{LexType: GeneralString}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:793
		{
			yyVAL.Type = RestrictedStringType{LexType: GraphicString}
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:794
		{
			yyVAL.Type = RestrictedStringType{LexType: IA5String}
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:795
		{
			yyVAL.Type = RestrictedStringType{LexType: ISO646String}
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:796
		{
			yyVAL.Type = RestrictedStringType{LexType: NumericString}
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:797
		{
			yyVAL.Type = RestrictedStringType{LexType: PrintableString}
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:798
		{
			yyVAL.Type = RestrictedStringType{LexType: TeletexString}
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:799
		{
			yyVAL.Type = RestrictedStringType{LexType: T61String}
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:800
		{
			yyVAL.Type = RestrictedStringType{LexType: UniversalString}
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:801
		{
			yyVAL.Type = RestrictedStringType{LexType: UTF8String}
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:802
		{
			yyVAL.Type = RestrictedStringType{LexType: VideotexString}
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:803
		{
			yyVAL.Type = RestrictedStringType{LexType: VisibleString}
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:808
		{
			yyVAL.Type = CharacterStringType{}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:813
		{
			yyVAL.Type = TypeReference("GeneralizedTime")
		}
	case 191:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:818
		{
			yyVAL.Type = ConstraintedType{yyDollar[1].Type, yyDollar[2].Constraint}
		}
	case 193:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:826
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].Type}, yyDollar[2].Constraint}
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:827
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].Type}, SingleElementConstraint(yyDollar[2].Elements)}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:830
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].NamedType}, yyDollar[2].Constraint}
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:831
		{
			yyVAL.Type = ConstraintedType{SequenceOfType{yyDollar[4].NamedType}, SingleElementConstraint(yyDollar[2].Elements)}
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line asn1.y:836
		{
			yyVAL.Constraint = Constraint{ConstraintSpec: yyDollar[2].ConstraintSpec}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:839
		{
			yyVAL.ConstraintSpec = yyDollar[1].SubtypeConstraint
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:849
		{
			yyVAL.SubtypeConstraint = yyDollar[1].SubtypeConstraint
		}
	case 202:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line asn1.y:850
		{
			yyVAL.SubtypeConstraint = append(yyDollar[1].SubtypeConstraint, yyDollar[5].ElementSetSpec)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:853
		{
			yyVAL.SubtypeConstraint = SubtypeConstraint{yyDollar[1].ElementSetSpec}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:859
		{
			yyVAL.ElementSetSpec = yyDollar[1].Unions
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:860
		{
			yyVAL.ElementSetSpec = yyDollar[2].Exclusions
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:863
		{
			yyVAL.Unions = Unions{yyDollar[1].Intersections}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:864
		{
			yyVAL.Unions = append(yyDollar[1].Unions, yyDollar[3].Intersections)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:870
		{
			yyVAL.Intersections = Intersections{yyDollar[1].IntersectionElements}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:871
		{
			yyVAL.Intersections = append(yyDollar[1].Intersections, yyDollar[3].IntersectionElements)
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:877
		{
			yyVAL.IntersectionElements = IntersectionElements{Elements: yyDollar[1].Elements}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:878
		{
			yyVAL.IntersectionElements = IntersectionElements{Elements: yyDollar[1].Elements, Exclusions: yyDollar[2].Exclusions}
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:884
		{
			yyVAL.Exclusions = Exclusions{yyDollar[2].Elements}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:893
		{
			yyVAL.Elements = yyDollar[1].Elements
		}
	case 222:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:895
		{
			yyVAL.Elements = yyDollar[2].ElementSetSpec
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:910
		{
			yyVAL.Elements = SingleValue{yyDollar[1].Value}
		}
	case 228:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line asn1.y:915
		{
			yyVAL.Elements = ValueRange{yyDollar[1].RangeEndpoint, yyDollar[3].RangeEndpoint}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:918
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value}
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:919
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value, IsOpen: true}
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:922
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[1].Value}
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:923
		{
			yyVAL.RangeEndpoint = RangeEndpoint{Value: yyDollar[2].Value, IsOpen: true}
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:927
		{
			yyVAL.Value = nil
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:931
		{
			yyVAL.Value = nil
		}
	case 237:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line asn1.y:936
		{
			yyVAL.Elements = SizeConstraint{yyDollar[2].Constraint}
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line asn1.y:941
		{
			yyVAL.Elements = TypeConstraint{yyDollar[1].Type}
		}
	}
	goto yystack /* stack new state and value */
}
