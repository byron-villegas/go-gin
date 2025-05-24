package main

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	var outputFile *os.File
	var err error
	outputFile, err = os.Create("godog-report.json")

	if err != nil {
		t.Fatalf("failed to create report file: %v", err)
	}

	opts := godog.Options{
		Format:   "cucumber",
		Paths:    []string{"."},
		Output:   outputFile,
		TestingT: t,
		Strict:   true,
	}
	godog.TestSuite{
		Name:                "acceptance",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()
}
