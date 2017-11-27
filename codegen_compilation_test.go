package asn1go

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"text/template"
)

var runCommandError = `Failed to run program: %v
stdout:

%v

stderr:

%v
`

var driverProgramTemplate = `
package main

import (
	"./{{ .ModuleName }}"
	"fmt"
	"os"
	"encoding/asn1"
)

var _ = os.Exit
var _ = fmt.Println

func main() {
  ok := true
  {{ $ctx := . }}
  {{ range $index, $assignment := .TypeAssignments }}
  	{{ $typeName := printf "%v.%v" $ctx.ModuleName (call $ctx.Goify $assignment.TypeReference.Name) }}
  	{
  		fmt.Println("Testing {{ $typeName }}...")
	  	var x {{ $typeName }}
	  	data, err := asn1.Marshal(x)
	  	if err != nil {
	  		fmt.Println("Marshal error: " + err.Error())
	  		ok = false
	  	}
	  	var y {{ $typeName }}
	  	_, err = asn1.Unmarshal(data, &y)
	  	if err != nil {
	  		fmt.Println("Unmarshal error: " + err.Error())
	  		ok = false
	  	}
  	}
  {{ end }}
  if !ok {
  	fmt.Println("Test failed")
  	os.Exit(1)
  }
}
`

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	stdout := bytes.NewBufferString("")
	stderr := bytes.NewBufferString("")
	cmd.Stdout, cmd.Stderr = stdout, stderr
	err := cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf(runCommandError, err.Error(), stdout, stderr))
	} else {
		return nil
	}
}

func createTestTemp() (string, error) {
	return ioutil.TempDir("", "asn1go_test")
}

func renderModule(baseDir, moduleName, module string) (filePath string, err error) {
	// create module
	modulePath := filepath.Join(baseDir, moduleName)
	err = os.Mkdir(modulePath, 0755)
	if err != nil {
		return "", err
	}
	// render module contents
	filePath = filepath.Join(modulePath, "module.go")
	err = ioutil.WriteFile(filePath, []byte(module), 0644)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

type driverProgramContext struct {
	ModuleName      string
	Goify           func(string) string
	TypeAssignments []TypeAssignment
}

func filterTypeAssignments(assignments AssignmentList) (res []TypeAssignment) {
	for _, assignment := range assignments {
		if typeAssignment, ok := assignment.(TypeAssignment); ok {
			res = append(res, typeAssignment)
		}
	}
	return res
}

func renderDriverProgram(driverPath, moduleName string, module ModuleDefinition) error {
	templ, err := template.New("main.go").Parse(driverProgramTemplate)
	if err != nil {
		return err
	}
	driverFile, err := os.Create(driverPath)
	if err != nil {
		return err
	}
	defer driverFile.Close()
	ctx := driverProgramContext{
		ModuleName:      moduleName,
		Goify:           goifyName,
		TypeAssignments: filterTypeAssignments(module.ModuleBody.AssignmentList),
	}
	err = templ.Execute(driverFile, ctx)
	if err != nil {
		return err
	}
	return nil
}

func tryCompileModule(moduleName, module string) error {
	tempPath, err := createTestTemp()
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempPath)
	// create module
	filePath, err := renderModule(tempPath, moduleName, module)
	if err != nil {
		return err
	}
	// test module compiles
	err = runCommand("go", "build", filePath)
	if err != nil {
		return err
	}
	return nil
}

func dryrunModule(moduleName, module string, moduleAst ModuleDefinition) error {
	tempPath, err := createTestTemp()
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempPath)
	// create module
	_, err = renderModule(tempPath, moduleName, module)
	if err != nil {
		return err
	}
	// create driver program
	driverPath := filepath.Join(tempPath, "main.go")
	err = renderDriverProgram(driverPath, moduleName, moduleAst)
	if err != nil {
		return err
	}
	// test module compiles
	err = runCommand("go", "run", driverPath)
	if err != nil {
		return err
	}
	return nil
}

func TestKerberosCompiles(t *testing.T) {
	ast, err := ParseFile("examples/rfc4120.asn1")
	if err != nil {
		t.Fatal(err.Error())
	}
	module, err := generateDeclarationsString(*ast)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = tryCompileModule(ast.ModuleIdentifier.Reference, module)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestKerberosRuns(t *testing.T) {
	ast, err := ParseFile("examples/rfc4120.asn1")
	if err != nil {
		t.Fatal(err.Error())
	}
	module, err := generateDeclarationsString(*ast)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = dryrunModule(ast.ModuleIdentifier.Reference, module, *ast)
	if err != nil {
		t.Fatal(err.Error())
	}
}
