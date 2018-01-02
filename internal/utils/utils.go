package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
)

var runCommandError = `Failed to run program: %v
stdout:

%v

stderr:

%v
`

func RunCommandForResult(command string, args ...string) error {
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

func RunCommand(command string, args ...string) (stdout, stderr string, err error) {
	cmd := exec.Command(command, args...)
	bufStdout := bytes.NewBufferString("")
	bufStderr := bytes.NewBufferString("")
	cmd.Stdout, cmd.Stderr = bufStdout, bufStderr
	err = cmd.Run()
	stdout = bufStdout.String()
	stderr = bufStderr.String()
	return
}

func CreateTestTemp() (string, error) {
	return ioutil.TempDir("", "asn1go_test")
}
