package main

import "time"

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

type Sumary struct {
	Name        string
	AvgRespTime time.Duration
	RespErros   map[string]string
}
