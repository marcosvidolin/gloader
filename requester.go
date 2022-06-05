package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

type Requester struct {
	client *http.Client
}

func NewRequester(c *http.Client) *Requester {
	return &Requester{
		client: c,
	}
}

type RequesterReq struct {
	Name    string
	URL     string
	Method  string
	Headers map[string]string
	Body    string
}

type RequesterResp struct {
	Body          string
	Status        string
	StatusCode    int
	ContentLength int64
}

func (r *Requester) doHTTP(ctx context.Context, rr *RequesterReq) (*RequesterResp, error) {
	req, err := http.NewRequest(rr.Method, rr.URL, strings.NewReader(rr.Body))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	for k, v := range rr.Headers {
		req.Header.Add(k, v)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &RequesterResp{
		Body:          string(b),
		Status:        resp.Status,
		StatusCode:    resp.StatusCode,
		ContentLength: resp.ContentLength,
	}, nil
}
