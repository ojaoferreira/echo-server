package models

import "net/http"

type Data struct {
	Hostname             string            `json:"hostname"`
	Datetime             string            `json:"datetime"`
	Request              Request           `json:"request"`
	EnvironmentVariables map[string]string `json:"envs"`
}

type Request struct {
	Header http.Header `json:"headers"`
}
