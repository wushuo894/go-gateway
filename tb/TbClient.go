package tb

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"log"
	"time"
)

type QueueType struct {
	Ts     int64                  `json:"ts"`
	Values map[string]interface{} `json:"values"`
}

var Queue = map[string][]QueueType{}

func Connect() mqtt.Client {
	//println(util.Config)
	//thingsBoard := util.Config.ThingsBoard
	opts := mqtt.NewClientOptions()
	//opts.AddBroker(thingsBoard.Host)
	//opts.SetUsername(thingsBoard.UserName)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	client.Connect()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			jsonData, _ := json5.Marshal(Queue)
			println(string(jsonData))
			client.Publish("v1/gateway/telemetry", 1, false, string(jsonData))
		}
	}()

	return client
}
