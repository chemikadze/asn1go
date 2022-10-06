# asn1go

## Rationale

Currently existing Go libraries for asn1 support are either reflection-based (crypto/asn1) or 
very low-level (golang.org/x/crypto/cryptobyte). Idea is to provide Protobuf-like experience for 
working with ASN1 in Golang.

Originally started to back [gorberos](https://github.com/chemikadze/gorberos) kerberos wannabe-library.

## Architecture

1) Custom Lexer consumes from bufio.Reader and called by Parser
2) Parser is built using [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)
 based on BNF provided in [X.680](https://www.itu.int/ITU-T/studygroups/com17/languages/X.680-0207.pdf) standard. 
 As the result, Parser produces ASN1 module AST.
3) AST is used by Code Generator to produce declarations, serialization, and deserialization code.

## Roadmap

1) Lexer
 - [x] identifiers
 - [x] numbers 
 - [x] keywords
 - [x] symbols
 - [ ] strings, bit strings, hex strings
 - [ ] XML
2) Parser
 - [x] module definition BNF
 - [x] parse Kerberos (rfc4120)
 - [x] yield AST from parser
 - [x] parse SNMPv1 (rfc1157, rfc1155)
 - [ ] parse LDAP (rfc4511) 
 - [ ] SNMPv2 (rfc3411–3418)
3) Code Generator
 - [x] declaration generator
 - [x] crypto/asn1 compatible generation mode
 - [x] verify serialization on Kerberos
 - [ ] DER serialization generator
 - [ ] DER deserialization generator
4) Missing ASN features
 - [ ] SET
 - [ ] ANY (1988?)
 - [ ] WITH COMPONENTS
 - [ ] _Add more as found_

## Adding features

Real-world ASN.1 descriptions from RFC documents are used to gauge completeness of the implementation.
Standard itself is pretty huge, so features are added as needed.

Typically, this requires:

1) Modifying asn1.y to uncomment unsupported branches of syntax notation, and add missing declarations. 
   Refer to goyacc documentation for .y syntax notation.
2) Extending ast.go with necessary fields and types.
3) Modifying codegen.go to produce corresponding Go declarations. Note that not all ASN declarations can be mapped to 
   crypto/asn1. 

For example, refer to [CHOICE implementation](https://github.com/chemikadze/asn1go/commit/884e30ce6a93c4e9df7ad7711889651fbcda01ce).