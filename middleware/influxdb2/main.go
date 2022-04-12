package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	// Create a client
	// You can generate an API Token from the "API Tokens Tab" in the UI
	client := influxdb2.NewClient("http://192.168.33.253:8086", "L_71iNg8DyZL7djxaQXm8b1stSPIV3Mah6SHB6y5Ihr6Gj-W3mMMJTAUdLFnVltqqs6DwKz2PQe3LUvYxwA9Rw==")
	// always close client at the end
	defer client.Close()

	//// get non-blocking write client
	//writeAPI := client.WriteAPI("jitstack", "monitor")
	//
	//// write line protocol
	//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
	//// Flush writes
	//writeAPI.Flush()
	//

	// get non-blocking write client
	writeAPI := client.WriteAPI("jitstack", "monitor")

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").AddTag("owner", "hank").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	// write point asynchronously
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()

	// Get query client
	queryAPI := client.QueryAPI("jitstack")

	query := `from(bucket:"monitor")|> range(start: -12h) |> filter(fn: (r) => r["_measurement"] == "stat")`

	//`from(bucket: "monitor")|> range(start: -5m, stop: "now")|> filter(fn: (r) => r["_measurement"] == "stat")|> aggregateWindow(every: "Property", fn: mean, createEmpty: false)|> yield(name: "mean")`
	// get QueryTableResult
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}

	tempStrut := make([]map[string]interface{}, 0)

	//if result.TableChanged() {
	//	//fmt.Printf("table: %s\n", result.TableMetadata().Columns())
	//
	//	for _, v := range result.TableMetadata().Columns() {
	//		fmt.Printf("table: %v\n", v)
	//
	//	}
	//}
	// Iterate over query response
	for result.Next() {

		tempStrut = append(tempStrut, result.Record().Values())

	}
	fmt.Printf("%v \n", tempStrut)

	// check for an error
	if result.Err() != nil {
		fmt.Printf("query parsing error: %\n", result.Err().Error())
	}

}
