package base

import (
	"go-gateway/tb"
	"time"
)

type ConnectorBase interface {
	// Run 运行
	Run()

	// ServerSideRpcHandler TB > gateway 数据
	ServerSideRpcHandler()
}

// Telemetry 上送设备数据
func (cb ConfigBase) Telemetry() {
	ts := time.Now().UnixMilli()

	queueType := tb.QueueType{
		Ts: ts,
		Values: map[string]interface{}{
			"temperature": 42,
			"humidity":    80,
		},
	}

	tb.Queue[cb.DeviceName] = []tb.QueueType{queueType}
}
