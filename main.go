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

func main() {
	// TODO: path from parameter
	scenarios, err := readYml("./data.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	resultCh := make(chan *ScenarioResult)
	ctx := context.Background()

	for _, s := range scenarios {
		execute(ctx, &s, resultCh)
	}

	// TODO: resume result
	for g := range resultCh {
		fmt.Println(g)
	}
}

func readYml(path string) ([]Scenario, error) {
	yfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := make(map[string]Scenario)

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		return nil, err
	}

	ss := make([]Scenario, len(data))
	for k, v := range data {
		v.ID = k
		ss = append(ss, v)
	}

	return ss, nil
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

		resp, err := req.doHTTP(context.Background(), &(RequesterReq{
			Method:  s.Method,
			URL:     s.URL,
			Body:    s.Body,
			Headers: s.Headers,
		}))

		if err != nil {
			sc.Err = err
			return
		}
		sc.Status = string(resp.StatusCode)

	}(resultCh, sce)

}
