package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v2"
)

func readYml(path string) (map[string]Scenario, error) {
	scenarios := make(map[string]Scenario, 0)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return scenarios, err
	}

	if err := yaml.Unmarshal(file, scenarios); err != nil {
		return scenarios, err
	}

	return scenarios, nil
}

func main() {
	ctx := context.Background()

	scenarios, err := readYml("./data.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	resultCh := make(chan *ScenarioResult, len(scenarios))
	for _, s := range scenarios {
		execute(ctx, &s, resultCh)
	}

	// TODO: resume result
	for g := range resultCh {
		fmt.Println(g)
	}
}

func execute(ctx context.Context, s *Scenario, resultCh chan<- *ScenarioResult) {
	// TODO: processes and iterations

	c := &http.Client{}
	if s.Timeout != 0 {
		c.Timeout = s.Timeout * time.Second
	}
	req := NewRequester(c)

	sce := &ScenarioResult{
		Scenario: s,
	}

	go func(rCh chan<- *ScenarioResult, sc *ScenarioResult) {
		start := time.Now()

		defer func(ch chan<- *ScenarioResult, s *ScenarioResult) {
			s.RespTime = time.Since(start)
			ch <- s
			close(ch)
		}(rCh, sc)

		resp, err := req.doHTTP(context.Background(), &(Request{
			Method:  s.Method,
			URL:     s.URL,
			Body:    s.Body,
			Headers: s.Headers,
		}))

		if err != nil {
			sc.Err = err
			return
		}
		sc.Status = fmt.Sprintf("%d", resp.StatusCode)

	}(resultCh, sce)
}
