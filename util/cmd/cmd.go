package cmd

import "os/exec"

// execute is the exec.Command for production
var execute = exec.Command

// ExecutionCMD is a common function for executing command in the linux
func ExecutionCMD(cmd string) (string, error) {
	execCmd := execute("bash", "-c", cmd)
	rawResult, err := execCmd.Output()
	return string(rawResult), err
}

// SetupTest replaces the executing command to the fake function implemented by golang
func SetupTestCmd(fakeExecute func(command string, args ...string) *exec.Cmd) {
	execute = fakeExecute
}

// TearDownTestCmd recovers the execute command function
func TearDownTestCmd() {
	execute = exec.Command
}
