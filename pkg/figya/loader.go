package figya

import (
	"fmt"
	"os"
	"plugin"
	"strings"
)

func FindModule(searchPath, module string) (string, error) {
	paths := strings.FieldsFunc(searchPath, func(c rune) bool {
		return c == ':'
	})
	for _, path := range paths {
		fullPath := fmt.Sprintf("%s/%s.so", path, module)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}
	return "", fmt.Errorf("module %s not found", module)
}

func LoadModule(path string, params map[string]interface{}) (Module, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	sym, err := plug.Lookup("New")
	if err != nil {
		return nil, err
	}
	new, ok := sym.(func(map[string]interface{}) (Module, error))
	if !ok {
		return nil, fmt.Errorf("could not load module %v", path)
	}
	return new(params)
}
