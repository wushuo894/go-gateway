package tb

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/titanous/json5"
	"go-gateway/connectors/base"
	"go-gateway/connectors/test"
	"go-gateway/util"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)

func Connect() mqtt.Client {
	thingsBoard := util.Config.ThingsBoard
	host := thingsBoard.Host
	port := thingsBoard.Port
	userName := thingsBoard.UserName
	password := thingsBoard.Password
	clientId := thingsBoard.ClientId

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://" + host + ":" + strconv.Itoa(port))
	opts.SetUsername(userName)
	if len(password) > 0 {
		opts.SetPassword(password)
	}

	if len(clientId) > 0 {
		opts.SetClientID(clientId)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}

	client.Connect()

	client.Subscribe("v1/gateway/rpc", 0, func(client mqtt.Client, msg mqtt.Message) {

	})

	go func() {
		for {
			time.Sleep(3 * time.Second)
			jsonData, _ := json.Marshal(base.Queue)
			println(string(jsonData))
			client.Publish("v1/gateway/telemetry", 1, false, string(jsonData))
		}
	}()

	go func() {
		for _, configBase := range util.Config.Connectors {
			fileName := configBase.FileName
			file, err := os.Open("config/" + fileName)
			if err != nil {
				fmt.Println("无法打开文件:", err)
				return
			}
			byteValue, err := io.ReadAll(file)
			if err != nil {
				println(err)
			}
			c := reflect.New(util.ConfigMap[configBase.DeviceType])
			connectorBase := c.Elem().Interface().(test.ConfigTest)
			json5.Unmarshal(byteValue, &connectorBase)
			println(&connectorBase)

			//go connectorBase.Run()
		}
	}()

	return client
}
