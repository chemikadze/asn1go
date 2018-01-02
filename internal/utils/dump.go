package utils

import (
	"bytes"
	"strings"
	"strconv"
	"time"
)

// Mon Jan 2 15:04:05 -0700 MST 2006
const WS_TIME_FORMAT = "2006-01-02 15:04:05"

func ParseWiresharkHex(str string) []byte {
	buf := bytes.NewBuffer([]byte{})
	for _, line := range strings.Split(str, "\n") {
		elems := strings.Split(line, " ")
		for _, byteRepr := range elems[1:] {
			if len(byteRepr) == 0 {
				continue
			}
			x, err := strconv.ParseInt(byteRepr, 16, 9)
			if err != nil {
				panic(err)
			}
			buf.WriteByte(byte(x))
		}
	}
	return buf.Bytes()
}

func ParseWiresharkTime(repr string) time.Time {
	t, err := time.Parse(WS_TIME_FORMAT, repr)
	if err != nil {
		panic(err)
	}
	return t
}