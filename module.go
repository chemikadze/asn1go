package asn1go

func init() {
	yyErrorVerbose = true
	RESERVED_WORDS = makeReservedWords()
	USEFUL_TYPES = makeBuiltinTypes()
}
