package main

import (
	"flag"
	"fmt"
	"github.com/chemikadze/asn1go"
	"os"
)

var usage = `
asn1go [[input] output]

Generates go file from input and writes to output.
If output is omitted, uses stdout. If input is omitted,
reads from stdin.
`

type flagsType struct {
	inputName  string
	outputName string
}

func failWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func parseFlags() (res flagsType) {
	flag.Parse()
	if flag.NArg() > 0 {
		res.inputName = flag.Arg(0)
	}
	if flag.NArg() == 2 {
		res.outputName = flag.Arg(1)
	}
	if flag.NArg() > 2 {
		failWithError(usage)
	}
	return res
}

func openChannels(inputName, outputName string) (input, output *os.File) {
	var err error
	input = os.Stdin
	output = os.Stdout
	if len(inputName) != 0 {
		input, err = os.Open(inputName)
		if err != nil {
			failWithError("Can't open %s for reading: %v", inputName, err.Error())
		}
	}
	if len(outputName) != 0 {
		output, err = os.OpenFile(outputName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			failWithError("File %v can not be written: %v", inputName, err.Error())
		}
	}
	return input, output
}

func main() {
	flag.Parse()
	flags := parseFlags()
	input, output := openChannels(flags.inputName, flags.outputName)

	module, err := asn1go.ParseStream(input)
	if err != nil {
		failWithError(err.Error())
	}

	gen := asn1go.NewCodeGenerator(asn1go.GEN_DECLARATIONS)
	err = gen.Generate(*module, output)
	if err != nil {
		failWithError(err.Error())
	}

	output.Close()
	input.Close()
}
