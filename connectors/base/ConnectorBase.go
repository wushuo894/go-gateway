package base

import (
	"sync"
	"time"
)

type QueueType struct {
	Ts     int64                  `json:"ts"`
	Values map[string]interface{} `json:"values"`
}

var (
	Queue       = &map[string][]QueueType{}
	QueueLocker = &sync.Mutex{}
)

type ConnectorBase interface {
	// Run 运行
	Run()

	// ServerSideRpcHandler TB > gateway 数据
	ServerSideRpcHandler(m map[string]any) any

	// AttributeUpdatesHandler 更新共享属性
	AttributeUpdatesHandler(m map[string]any)
}

// Run 运行
func (cb ConfigBase) Run() {
}

// Telemetry 上送设备数据
func (cb ConfigBase) Telemetry(values *map[string]any) {
	QueueLocker.Lock()
	defer QueueLocker.Unlock()
	ts := time.Now().UnixMilli()

	if len(*values) < 1 {
		delete(*Queue, cb.DeviceName)
		return
	}

	queueType := QueueType{
		Ts:     ts,
		Values: *values,
	}
	(*Queue)[cb.DeviceName] = []QueueType{queueType}

}
