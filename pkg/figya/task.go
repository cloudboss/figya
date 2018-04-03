package figya

import (
	"strings"
)

type Whener func() (bool, error)

type Task struct {
	Name       string `json:"name"`
	ModuleName string `json:"module"`
	Params     map[string]interface{}
	Module     Module
	WhenField  string `json:"when"`
}

func (t *Task) When() Whener {
	when := t.WhenField
	if when == "" {
		return func() (bool, error) {
			return true, nil
		}
	}
	parts := strings.Fields(t.WhenField)
	command := parts[0]
	args := []string{}
	if len(parts) > 1 {
		args = parts[1:]
	}
	return func() (bool, error) {
		commandOutput, err := RunCommand(command, args...)
		if err != nil {
			return false, err
		}
		return commandOutput.ExitStatus == 0, nil
	}
}
