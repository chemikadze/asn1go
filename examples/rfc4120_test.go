package examples

import (
	"encoding/asn1"
	"fmt"
	"github.com/chemikadze/asn1go/internal/utils"
	"testing"
)

//go:generate go run ../cmd/asn1go/main.go -package examples rfc4120.asn1 rfc4120_generated.go

func TestMessagesDeclared(t *testing.T) {
	var (
		_ KDC_REQ
		_ KDC_REP
		_ AS_REQ
		_ AS_REP
		_ AP_REQ
		_ AP_REP
	)
}

type testCase struct {
	bytes    []byte
	value    interface{} // should be pointer
	expected interface{} // should be value
}

func messageTest(t *testing.T, item testCase) {
	// verify it can be parsed
	parsed := item.value
	rest, err := asn1.Unmarshal(item.bytes, parsed)
	if err != nil {
		t.Errorf("Failed to parse: %v", err.Error())
	}
	if len(rest) != 0 {
		t.Errorf("Expected no trailing data, got %v bytes", len(rest))
	}
	if es, ps := fmt.Sprintf("&%+v", item.expected), fmt.Sprintf("%+v", parsed); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}

	// verify that it can be generated and serialization is reversible
	generatedBytes, err := asn1.Marshal(item.expected)
	if err != nil {
		t.Fatalf("Failed to marshall message: %v", err.Error())
	}
	_, err = asn1.Unmarshal(generatedBytes, parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshall message: %v", err.Error())
	}
	if es, ps := fmt.Sprintf("&%+v", item.expected), fmt.Sprintf("%+v", parsed); es != ps {
		t.Errorf("Repr mismatch:\n exp: %v\n got: %v", es, ps)
	}
}

func TestKdcReq(t *testing.T) {
	msgBytes := utils.ParseWiresharkHex(`
0000   30 81 aa a1 03 02 01 05 a2 03 02 01 0a a3 0e 30
0010   0c 30 0a a1 04 02 02 00 95 a2 02 04 00 a4 81 8d
0020   30 81 8a a0 07 03 05 00 00 00 00 10 a1 17 30 15
0030   a0 03 02 01 01 a1 0e 30 0c 1b 0a 63 68 65 6d 69
0040   6b 61 64 7a 65 a2 10 1b 0e 41 54 48 45 4e 41 2e
0050   4d 49 54 2e 45 44 55 a3 23 30 21 a0 03 02 01 02
0060   a1 1a 30 18 1b 06 6b 72 62 74 67 74 1b 0e 41 54
0070   48 45 4e 41 2e 4d 49 54 2e 45 44 55 a5 11 18 0f
0080   32 30 31 38 30 31 30 33 30 36 30 34 30 37 5a a7
0090   06 02 04 64 21 bb 89 a8 14 30 12 02 01 12 02 01
00a0   11 02 01 10 02 01 17 02 01 19 02 01 1a
`)
	expected := AS_REQ{
		Pvno:     5,
		Msg_type: 10,
		Padata: []PA_DATA{
			{149, []byte{}},
		},
		Req_body: KDC_REQ_BODY{
			Kdc_options: asn1.BitString{[]byte{0x00, 0x00, 0x00, 0x10}, 32},
			Cname:       PrincipalName{1, []KerberosString{"chemikadze"}},
			Realm:       "ATHENA.MIT.EDU",
			Sname:       PrincipalName{2, []KerberosString{"krbtgt", "ATHENA.MIT.EDU"}},
			Till:        utils.ParseWiresharkTime("2018-01-03 06:04:07"),
			Nonce:       1679932297,
			Etype:       []Int32{18, 17, 16, 23, 25, 26},
		},
	}

	messageTest(t, testCase{bytes: msgBytes, value: new(AS_REQ), expected: expected})
}

func TestKrbError(t *testing.T) {
	msgBytes := utils.ParseWiresharkHex(`
0000   30 81 b2 a0 03 02 01 05 a1 03 02 01 1e a2 11 18
0010   0f 32 30 32 33 30 33 32 37 31 35 35 31 33 37 5a
0020   a4 11 18 0f 32 30 31 38 30 31 30 32 30 36 30 34
0030   30 37 5a a5 05 02 03 04 88 a8 a6 03 02 01 06 a7
0040   10 1b 0e 41 54 48 45 4e 41 2e 4d 49 54 2e 45 44
0050   55 a8 17 30 15 a0 03 02 01 01 a1 0e 30 0c 1b 0a
0060   63 68 65 6d 69 6b 61 64 7a 65 a9 10 1b 0e 41 54
0070   48 45 4e 41 2e 4d 49 54 2e 45 44 55 aa 23 30 21
0080   a0 03 02 01 02 a1 1a 30 18 1b 06 6b 72 62 74 67
0090   74 1b 0e 41 54 48 45 4e 41 2e 4d 49 54 2e 45 44
00a0   55 ab 12 1b 10 43 4c 49 45 4e 54 5f 4e 4f 54 5f
00b0   46 4f 55 4e 44
`)
	expected := KRB_ERROR{
		Pvno:       5,
		Msg_type:   30,
		Ctime:      utils.ParseWiresharkTime("2023-03-27 15:51:37"),
		Stime:      utils.ParseWiresharkTime("2018-01-02 06:04:07"),
		Susec:      297128,
		Error_code: 6,
		Crealm:     "ATHENA.MIT.EDU",
		Cname:      PrincipalName{1, []KerberosString{"chemikadze"}},
		Realm:      "ATHENA.MIT.EDU",
		Sname:      PrincipalName{2, []KerberosString{"krbtgt", "ATHENA.MIT.EDU"}},
		E_text:     "CLIENT_NOT_FOUND",
	}

	messageTest(t, testCase{bytes: msgBytes, value: new(AS_REQ), expected: expected})
}
