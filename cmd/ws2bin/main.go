package main

import (
	"fmt"
	"github.com/chemikadze/asn1go/internal/utils"
	"io/ioutil"
	"os"
)

func main() {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read: %v", err.Error())
		os.Exit(1)
	}
	outputBytes := utils.ParseWiresharkHex(string(inputBytes))
	os.Stdout.Write(outputBytes)
	os.Stdout.Sync()
	os.Stdout.Close()
}
