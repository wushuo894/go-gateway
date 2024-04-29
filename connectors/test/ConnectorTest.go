package test

// Run 运行
func (c ConfigTest) Run() {
	//for {
	//	println(c.A)
	//	c.Telemetry()
	//	time.Sleep(1 * time.Second)
	//}
}

// ServerSideRpcHandler TB > gateway 数据
func (ConfigTest) ServerSideRpcHandler(m map[string]any) any {
	return "1111"
}
