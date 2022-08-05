package main

import (
	"fmt"
	"time"
)

func main() {
	voiceAlertStep, _ := time.ParseDuration("-5m")

	fmt.Println(voiceAlertStep.Minutes())
	fmt.Println(voiceAlertStep.Seconds())

	t := time.Now().Add(voiceAlertStep)
	fmt.Println(t.Format("2006-01-02 15:04:05"))

}
