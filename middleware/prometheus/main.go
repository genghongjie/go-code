package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func main() {
	ExampleAPI_Query()
	//ExampleAPI_QuerSnapshot()
	fmt.Println(fmt.Sprintf("%v", time.Now()))
}

func ExampleAPI_Query() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	api := v1.NewAPI(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := api.Query(ctx, "up", time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	j, _ := json.Marshal(result)
	fmt.Println(string(j))
	//fmt.Printf("Result:\n%v\n", result)
}

func ExampleAPI_QueryRange() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := api.QueryRange(ctx, "rate(prometheus_tsdb_head_samples_appended_total[5m])", r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func ExampleAPI_Series() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lbls, warnings, err := api.Series(ctx, []string{
		"{__name__=~\"scrape_.+\",job=\"node\"}",
		"{__name__=~\"scrape_.+\",job=\"prometheus\"}",
	}, time.Now().Add(-time.Hour), time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Println("Result:")
	for _, lbl := range lbls {
		fmt.Println(lbl)
	}
}
