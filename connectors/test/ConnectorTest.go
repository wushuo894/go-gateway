package test

import (
	"time"
)

// Run 运行
func (c ConfigTest) Run() {
	for {
		println(c.DeviceName)
		c.Telemetry()
		time.Sleep(1 * time.Second)
	}
}

// ServerSideRpcHandler TB > gateway 数据
func (ConfigTest) ServerSideRpcHandler() {

}
