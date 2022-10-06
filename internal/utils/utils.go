package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

var runCommandError = `failed to run program %q: %v
output:

%v
`

func RunCommandForResult(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(runCommandError, append([]string{command}, args...), err.Error(), string(out))
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
