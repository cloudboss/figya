package figya

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
)

type CommandOutput struct {
	ExitStatus  int      `json:"exit_status"`
	Stdout      string   `json:"stdout"`
	Stderr      string   `json:"stderr"`
	StdoutLines []string `json:"stdout_lines"`
	StderrLines []string `json:"stderr_lines"`
}

func RunCommand(command string, args ...string) (*CommandOutput, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(command, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	waitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus)
	stdoutStr := strings.TrimSpace(stdout.String())
	stderrStr := strings.TrimSpace(stderr.String())
	commandOutput := &CommandOutput{
		ExitStatus:  waitStatus.ExitStatus(),
		Stdout:      stdoutStr,
		Stderr:      stderrStr,
		StdoutLines: strings.Split(stdoutStr, "\n"),
		StderrLines: strings.Split(stderrStr, "\n"),
	}
	return commandOutput, nil
}
