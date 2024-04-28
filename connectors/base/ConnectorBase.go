package base

import (
	"time"
)

type QueueType struct {
	Ts     int64                   `json:"ts"`
	Values *map[string]interface{} `json:"values"`
}

var Queue = &map[string][]QueueType{}

type ConnectorBase interface {
	// Run 运行
	Run()

	// ServerSideRpcHandler TB > gateway 数据
	ServerSideRpcHandler()
}

// Run 运行
func (cb ConfigBase) Run() {
}

// Telemetry 上送设备数据
func (cb ConfigBase) Telemetry() {
	ts := time.Now().UnixMilli()

	queueType := QueueType{
		Ts: ts,
		Values: &map[string]interface{}{
			"temperature": 42,
			"humidity":    80,
		},
	}

	(*Queue)[cb.DeviceName] = []QueueType{queueType}
}
