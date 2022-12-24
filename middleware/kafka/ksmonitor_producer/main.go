package main

import (
	"encoding/json"
	"fmt"
	"github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer/md_kafka"
	"github.com/labstack/gommon/log"
	"math/rand"
	"strconv"
	"time"
)

type Event struct {
	ItemApp    string `json:"itemapp"`
	ItemGroup  string `json:"itemgroup"`
	ItemType   string `json:"itemtype"`
	ItemId     string `json:"itemid"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	StatusFrom string `json:"statusfrom"`
	StatusTo   string `json:"statusto"`
	StatusFn   string `json:"statusfn"`
	StatusTn   string `json:"statustn"`
}

func main() {
	clientKfaka := md_kafka.KafkaEngine{}
	err := clientKfaka.GetEngine(3, md_kafka.KafakaEngineConfig{Topic: "Topic_Event", BootstrapServers: "192.168.102.57:29092"})
	if err != nil {
		log.Fatal(err)
	}

	dataarray := make([]Event, 0)

	//接入层
	GroupNames := []string{
		"JQ_JZJY_Switch_G1",
		"WGQ_JZJY_Switch_G1",
		"BJ_JZJY_Switch_G1",
		"XN_JZJY_Switch_G1",
		"BD_JZJY_Switch_G1",
	}
	for i := 0; i < 5; i++ {

		intR := rand.Intn(200)
		for j := 0; j < 3; j++ {
			alert := "正常"
			switch j {
			case 0:
				alert = "正常"
			case 1:
				alert = "报警"
			case 2:
				alert = "报警"

			}
			in := Event{
				ItemApp:   "J",
				ItemGroup: GroupNames[i],
				ItemType:  "switch",
				ItemId:    "10.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}
			dataarray = append(dataarray, in)
		}

	}

	for {

		b, _ := json.Marshal(dataarray[rand.Intn(15)])

		intR := rand.Intn(5000)
		fmt.Println(intR)
		fmt.Println(string(b))
		time.Sleep(time.Duration(intR) * time.Millisecond)
		clientKfaka.Produce("", b)
	}

}
