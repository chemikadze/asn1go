# asn1go

## Rationale

Built-in Go libraries for asn1 support are either reflection-based (crypto/asn1) or 
very low-level (golang.org/x/crypto/cryptobyte). Idea is to provide Protobuf-like experience for 
working with ASN1 in Golang.

Note: currently provided code generator implementation creates definitions to be used with crypto/asn1, 
so all its limitations (no real CHOICE support) apply for this project as well.

## Architecture

1) Custom Lexer consumes from bufio.Reader and called by Parser
2) Parser is built using [goyacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)
 based on BNF provided in [X.680](https://www.itu.int/ITU-T/studygroups/com17/languages/X.680-0207.pdf) standard. 
 As the result, Parser produces ASN1 module AST.
3) AST is used by Code Generator to produce declarations, serialization, and deserialization code.

## Supported features

### Feature categories

| Feature           | Parsing     | Codegen       |
|-------------------|-------------|---------------|
| Exports           | Syntax only | No            |
| Imports           | Yes         | No            |
| Type assignments  | Yes         | Yes           |
| Value assignments | Yes         | Partial [^f1] |
| XML               | No          |               |
| Objects           | No          |               |
| Parameterization  | No          |               |

[^f1]: Only literal values are supported, referenced values are not implemented.

### Types

| Type              | Parsing   | Codegen                                |
|-------------------|-----------|----------------------------------------|
| BIT STRING        | Yes       | Yes; named bits not translated         |
| BOOLEAN           | Yes       | Yes                                    |
| CHARACTER STRING  | Yes       | Yes                                    |
| CHOICE            | Yes       | Yes; common demoninator type is used   |
| Embedded PDV      | No        |                                        |
| External          | No        |                                        |
| ENUMERATED        | Yes [^t4] | Yes; alternative values not translated |
| Instance Of       | No        |                                        |
| INTEGER           | Yes       | Yes                                    |
| NULL              | Yes       |                                        |
| Object Class      | No        |                                        |
| Object Identifier | Yes       |                                        | 
| OCTET STRING      | Yes       | Yes                                    |
| REAL              | Yes       | Yes                                    |
| Relative OID      | No        |                                        |
| SEQUENCE          | Yes [^t1] | Yes                                    |
| SEQUENCE OF       | Yes       | Yes                                    |
| SET               | Yes [^t1] | Yes                                    |
| SET OF            | Yes       | Yes                                    |
| ANY               | Yes [^t2] | Yes                                    |
| Tagged types      | Yes       | Yes [^t3]                              |
| Constrained types | Partial   | Partial; generates wrapped type        |

[^t1]: With ASN.1 syntax limitations: two component type lists, exceptions and extension addition groups are not supported, extensions are not exposed in generated Go code.
[^t2]: Not defined in the latest ASN.1 standard.
[^t3]: Used by encoding/asn1 only in SEQUENCE and SET fields. CHOICE with tagged alternatives is represented as RawValue.
[^t4]: With ASN.1 syntax limitations: explicit extensibility and non-literal values are not supported.

### Values

| Value               | Parsing  | Codegen |
|---------------------|----------|---------|
| BOOLEAN             | Yes      | Yes     |
| INTEGER             | Yes      | Yes     |
| OID                 | Yes      | No      |
| Real                | Yes      | Yes     |
| Referenced          | No       |         |
| Object class fields | No       |         |
| BIT STRING          | No       |         |
| Other               | No       |         |

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
 - [x] parse SNMPv1 (rfc1157, rfc1155); no codegen, depends on CHOICE
 - [x] parse LDAP (rfc4511, partially - required minor modifications); no codegen, depends on CHOICE
 - [ ] parse X.509 (rfc 5280) - depends on ANY
 - [ ] SNMPv2 (rfc3411–3418)
3) Code Generator
 - [x] declaration generator
 - [x] crypto/asn1 compatible generation mode
 - [x] verify serialization on Kerberos
 - [ ] DER serialization generator
 - [ ] DER deserialization generator
4) Supported ASN features
 - [x] SET type
 - [x] ANY type (1988 standard) - mapped to interface{}
 - [x] CHOICE type - mapped to interface{}, or asn1.RawValue if selections are tagged
 - [ ] Extensions in SEQUENCE, SET, CHOICE
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