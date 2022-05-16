package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/genghongjie/go-code/middleware/kafka/util"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type student struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}
	configFile := os.Args[1]
	conf := util.ReadConfig(configFile)

	topic := "purchases"
	p, err := kafka.NewProducer(&conf)
	fmt.Printf("1111")
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	//可选的chan Event，可以用来监听发送的结果
	delivery_chan := make(chan kafka.Event, 10000)

	for n := 0; n < 10; n++ {
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, delivery_chan)

		s := student{
			Id:   1,
			Age:  18,
			Name: "Hank",
		}
		bb, _ := json.Marshal(s)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			//Key:            []byte(key),
			Value: bb,
		}, delivery_chan)
	}
	//监听结果
	e := <-delivery_chan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	close(delivery_chan)
	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()
}
