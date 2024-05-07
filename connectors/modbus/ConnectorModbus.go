package modbus

import (
	"log"
	"time"
)

// Run 运行
func (c ConfigModbus) Run() {
	for {
		m := map[string]any{}
		for _, infoModbus := range c.Timeseries {
			ret, err := c.item(infoModbus, nil)
			if err != nil {
				log.Println("Run Error in item:", c.DeviceName, infoModbus, err)
			}
			if ret != nil {
				m[infoModbus.Tag] = ret
			}

			time.Sleep(1 * time.Second)
		}
		if len(m) > 0 {
			c.Telemetry(&m)
		}
		time.Sleep(1 * time.Second)
	}
}

// ServerSideRpcHandler TB > gateway 数据
func (c ConfigModbus) ServerSideRpcHandler(m map[string]any) any {
	method := m["method"].(string)
	params := m["params"].(float64)
	for _, modbus := range c.Rpc {
		tag := modbus.Tag
		if tag != method {
			continue
		}

		ret, err := c.item(modbus, uint16(params))
		if err != nil {
			log.Println("Error in item:", modbus, err)
		}
		return ret
	}
	return nil
}
