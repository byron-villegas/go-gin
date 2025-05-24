package main

import (
	"fmt"
	"io"
	"net/http"

	"os"

	"github.com/cucumber/godog"
)

var apiHost string = "http://localhost:8080/api"

func init() {
	if os.Getenv("API_HOST") != "" {
		apiHost = os.Getenv("API_HOST")
	}
}

var endpoint string
var response *http.Response
var body []byte

func anEndpoint(ep string) error {
	endpoint = ep
	return nil
}

func iSendAGETRequestTo() error {
	var err error
	response, err = http.Get(apiHost + endpoint)
	if err != nil {
		return err
	}
	body, _ = io.ReadAll(response.Body)
	return nil
}

func theResponseCodeShouldBe(code int) error {
	if response.StatusCode != code {
		return fmt.Errorf("expected %d, got %d", code, response.StatusCode)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^An endpoint "(.+)"$`, anEndpoint)
	ctx.Step(`^I send a GET request$`, iSendAGETRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
}
