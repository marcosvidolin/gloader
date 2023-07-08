package main

import (
	"time"

	"gopkg.in/yaml.v2"
)

type Scenario struct {
	ID         string            `yaml:"id"`
	Desc       string            `yaml:"description"`
	URL        string            `yaml:"url"`
	Method     string            `yaml:"method"`
	Headers    map[string]string `yaml:"headers"`
	Body       string            `yaml:"body,omitempty"`
	Iterations int               `yaml:"iterations"`
	Processes  int               `yaml:"processes"`
	Timeout    time.Duration     `yaml:"timeout"`
}

type ScenarioResult struct {
	Scenario *Scenario
	Status   string
	Err      error
	RespTime time.Duration
}

type scenarioReader struct {
	file []byte
}

func NewScenarioReader(file []byte) scenarioReader {
	return scenarioReader{
		file: file,
	}
}

func (s scenarioReader) Read() ([]*Scenario, error) {
	data := make(map[string]Scenario)
	err := yaml.Unmarshal(s.file, &data)
	if err != nil {
		return nil, err
	}

	ss := make([]*Scenario, 0, len(data))
	for k, v := range data {
		v.ID = k
		ss = append(ss, &v)
	}

	return ss, nil
}
