package figya

type Runner func() *Result

type Predicate func() (bool, error)

type Action func() *Result

type Result struct {
	Succeeded    bool        `json:"succeeded"`
	Changed      bool        `json:"changed"`
	Error        *string     `json:"error,omitempty"`
	Module       string      `json:"module"`
	ModuleOutput interface{} `json:"module_output,omitempty"`
}

func DoIf(module string, condition Predicate, do Action) *Result {
	done, err := condition()
	if err != nil {
		errStr := err.Error()
		return &Result{
			Succeeded: false,
			Changed:   false,
			Error:     &errStr,
			Module:    module,
		}
	}
	if !done {
		return do()
	}
	return &Result{
		Succeeded: true,
		Changed:   false,
		Module:    module,
	}
}
