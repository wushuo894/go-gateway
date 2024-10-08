package tb

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go-gateway/connectors/base"
	"go-gateway/util"
	"io"
	"log"
	"os"
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
	opts.SetAutoReconnect(true)
	opts.SetUsername(userName)
	opts.ConnectTimeout = 10 * time.Second
	if len(password) > 0 {
		opts.SetPassword(password)
	}

	if len(clientId) > 0 {
		opts.SetClientID(clientId)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("Error connecting to mqtt server", token.Error())
	}

	connectors := &util.Config.Connectors

	// RPC
	client.Subscribe("v1/gateway/rpc", 1, func(client mqtt.Client, msg mqtt.Message) {
		m := map[string]any{}
		err := json.Unmarshal(msg.Payload(), &m)
		if err != nil {
			log.Println(err)
		}
		device := m["device"].(string)

		for _, configBase := range *connectors {
			if configBase.DeviceName != device {
				continue
			}
			if &configBase.Connector == nil {
				log.Fatalln("Connector Not Found")
			}
			s, _ := json.Marshal(m)
			log.Println(string(s))
			data := m["data"].(map[string]any)
			ret := (*configBase.Connector).ServerSideRpcHandler(data)
			m["id"] = data["id"]
			m["data"] = ret
			s, _ = json.Marshal(m)
			token := client.Publish(msg.Topic(), byte(1), false, s)
			token.Wait()
			log.Println(string(s))
			return
		}
	})

	// 共享属性
	client.Subscribe("v1/gateway/attributes", 1, func(client mqtt.Client, msg mqtt.Message) {
		// {"data":{"test":"123"},"device":"无线空调控制器"}
		log.Println("v1/gateway/attributes")
		log.Println(string(msg.Payload()))

		m := map[string]any{}
		err := json.Unmarshal(msg.Payload(), &m)
		if err != nil {
			log.Println(err)
		}
		device := m["device"].(string)
		for _, configBase := range *connectors {
			if configBase.DeviceName != device {
				continue
			}
			data := m["data"].(map[string]any)
			(*configBase.Connector).AttributeUpdatesHandler(data)
			return
		}
	})

	go func() {
		for {
			time.Sleep(3 * time.Second)
			base.QueueLocker.Lock()
			if len(*base.Queue) < 1 {
				base.QueueLocker.Unlock()
				continue
			}
			jsonData, err := json.Marshal(base.Queue)
			if err != nil {
				log.Println(err)
				base.QueueLocker.Unlock()
				continue
			}
			base.Queue = &map[string][]base.QueueType{}
			base.QueueLocker.Unlock()
			log.Println(string(jsonData))
			client.Publish("v1/gateway/telemetry", 1, false, string(jsonData))
		}
	}()

	go func() {
		for _, configBase := range *connectors {
			fileName := configBase.FileName
			file, err := os.Open("config/" + fileName)
			if err != nil {
				fmt.Println("无法打开文件:", err)
				continue
			}
			byteValue, err := io.ReadAll(file)
			if err != nil {
				log.Fatalln(err)
			}
			connectorBase := util.ConfigFuncMap[configBase.DeviceType](byteValue, *configBase)
			configBase.Connector = &connectorBase
			go connectorBase.Run()
		}
	}()

	return client
}
