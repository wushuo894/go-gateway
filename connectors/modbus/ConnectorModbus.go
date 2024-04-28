package modbus

// Run 运行
func (c ConfigModbus) Run() {
	c.Telemetry()
}

// ServerSideRpcHandler TB > gateway 数据
func (ConfigModbus) ServerSideRpcHandler() {

}
