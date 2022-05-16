package main

import (
	"fmt"
	"time"

	"k8s.io/kubernetes/staging/src/k8s.io/apimachinery/pkg/util/json"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:    "http://192.168.33.252:8086",
		Timeout: 1 * time.Second,
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()
	fmt.Println("2")

	q := client.NewQuery("SELECT * FROM demo", "test_database", "")
	response, err := c.Query(q)
	if err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
	a, _ := json.Marshal(response.Results)
	fmt.Println(string(a))

}
