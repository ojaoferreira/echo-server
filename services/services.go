package services

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ojaoferreira/echo-server/models"
)

func DataService(r *http.Request) models.Data {
	data := models.Data{}

	data.Datetime = time.Now().Format("2006-01-02 15:04:05.000")

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	data.Hostname = hostname

	data.Request.Header = http.Header{}
	data.Request.Header.Add("Path", r.RequestURI)
	data.Request.Header.Add("Method", r.Method)
	data.Request.Header.Add("Procotol", r.Proto)
	for key, value := range r.Header {
		data.Request.Header[key] = value
	}

	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		arrEnv := strings.Split(env, "=")
		envMap[arrEnv[0]] = arrEnv[1]
	}
	data.EnvironmentVariables = envMap

	return data
}
