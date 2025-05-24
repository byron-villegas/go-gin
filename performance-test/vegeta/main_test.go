package performance

import (
	"os"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

var apiHost string = "http://localhost:8080/api"

func init() {
	if os.Getenv("API_HOST") != "" {
		apiHost = os.Getenv("API_HOST")
	}
}

func TestProductEndpointPerformance(t *testing.T) {
	rate := vegeta.Rate{Freq: 10, Per: time.Second} // 10 requests por segundo
	duration := 5 * time.Second

	endpoint := "products"

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    apiHost + "/products",
	})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "GET /products") {
		metrics.Add(res)
	}
	metrics.Close()

	reportsDir := "reports"

	if err := os.MkdirAll(reportsDir, os.ModePerm); err != nil {
		t.Fatalf("failed to create reports directory: %v", err)
	}

	reportFile := "vegeta-" + endpoint + "-report"

	// Reporte en formato texto
	textFile, err := os.Create(reportsDir + "/" + reportFile + ".txt")
	if err != nil {
		t.Fatalf("failed to create text report file: %v", err)
	}
	defer textFile.Close()

	textReporter := vegeta.NewTextReporter(&metrics)
	if err := textReporter.Report(textFile); err != nil {
		t.Errorf("failed to write vegeta text report: %v", err)
	}

	// Reporte en formato JSON
	jsonFile, err := os.Create(reportsDir + "/" + reportFile + ".json")
	if err != nil {
		t.Fatalf("failed to create JSON report file: %v", err)
	}
	defer jsonFile.Close()

	jsonReporter := vegeta.NewJSONReporter(&metrics)
	if err := jsonReporter.Report(jsonFile); err != nil {
		t.Errorf("failed to write vegeta JSON report: %v", err)
	}

	if metrics.Success < 0.95 {
		t.Errorf("success rate too low: %.2f%%", metrics.Success*100)
	}
	t.Logf("99th percentile latency: %s", metrics.Latencies.P99)
}
