// Binary asn1go generates go code from ASN.1 definitions.
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
If output is omitted, writes go file to stdout. 
If input is omitted as well, reads from stdin.

Input should be an ASN.1 module file.
Output is a go file.
`

type flagsType struct {
	inputName      string
	outputName     string
	packageName    string
	defaultIntRepr string
}

func failWithError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func parseFlags(args []string) (res flagsType) {
	cmd := flag.NewFlagSet(args[0], flag.ExitOnError)
	cmd.StringVar(&res.packageName, "package", "", "package name for generated code")
	cmd.StringVar(&res.defaultIntRepr, "default-integer-repr", "int64", "Type for numeric types (int64, big.Int)")
	cmd.Parse(args[1:])
	if cmd.NArg() > 0 {
		res.inputName = cmd.Arg(0)
	}
	if cmd.NArg() == 2 {
		res.outputName = cmd.Arg(1)
	}
	if cmd.NArg() > 2 {
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
		if err := output.Truncate(0); err != nil {
			failWithError("Failed to truncate file: %v", err.Error())
		}
	}
	return input, output
}

func main() {
	flags := parseFlags(os.Args)
	input, output := openChannels(flags.inputName, flags.outputName)

	module, err := asn1go.ParseStream(input)
	if err != nil {
		failWithError(err.Error())
	}
	params := asn1go.GenParams{
		Package:     flags.packageName,
		IntegerRepr: asn1go.IntegerRepr(flags.defaultIntRepr),
	}
	gen := asn1go.NewCodeGenerator(params)
	err = gen.Generate(*module, output)
	if err != nil {
		failWithError(err.Error())
	}

	output.Close()
	input.Close()
}
