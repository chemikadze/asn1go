package utils

import "testing"

func TestParseWiresharkHex(t *testing.T) {
	repr := `
0000   6a 81 ad 30 81 aa a1 03 02 01 05 a2 03 02 01 0a
0010   a3 0e 30 0c 30 0a a1 04 02 02 00 95 a2 02 04 00
0020   a4 81 8d 30 81 8a a0 07 03 05 00 00 00 00 10 a1
0030   17 30 15 a0 03 02 01 01 a1 0e 30 0c 1b 0a 63 68
0040   65 6d 69 6b 61 64 7a 65 a2 10 1b 0e 41 54 48 45
0050   4e 41 2e 4d 49 54 2e 45 44 55 a3 23 30 21 a0 03
0060   02 01 02 a1 1a 30 18 1b 06 6b 72 62 74 67 74 1b
0070   0e 41 54 48 45 4e 41 2e 4d 49 54 2e 45 44 55 a5
0080   11 18 0f 32 30 31 38 30 31 30 33 30 36 30 34 30
0090   37 5a a7 06 02 04 64 21 bb 89 a8 14 30 12 02 01
00a0   12 02 01 11 02 01 10 02 01 17 02 01 19 02 01 1a
`
	parsed := ParseWiresharkHex(repr)
	expectedBeginning := []byte{0x6a, 0x81, 0xad, 0x30}
	for i, exp := range expectedBeginning {
		was := parsed[i]
		if was != exp {
			t.Errorf("Not matched at %v: %v != %v", i, was, exp)
		}
	}
}
