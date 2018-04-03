package figya

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Playbook struct {
	Vars      map[string]interface{} `json:"vars"`
	Tasks     []*Task                `json:"tasks"`
	succeeded bool
}

func NewPlaybook(path string) (*Playbook, error) {
	jsn, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var playbook Playbook
	err = json.Unmarshal(jsn, &playbook)
	if err != nil {
		return nil, err
	}
	for _, task := range playbook.Tasks {
		path := fmt.Sprintf("modules/%s/%s.so", task.ModuleName, task.ModuleName)
		module, err := LoadModule(path, task.Params)
		if err != nil {
			return nil, err
		}
		task.Module = module
	}
	return &playbook, nil
}

func (p *Playbook) print(result *Result) error {
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b))
	return nil
}

func (p *Playbook) Run() ([]*Result, error) {
	var results []*Result
	for _, task := range p.Tasks {
		result := task.Module.Run()
		err := p.print(result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
		if result.Error != nil {
			p.succeeded = false
			return results, fmt.Errorf(*result.Error)
		}
	}
	p.succeeded = true
	return results, nil
}
