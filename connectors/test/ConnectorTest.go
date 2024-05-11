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
func (c ConfigTest) ServerSideRpcHandler(m map[string]any) any {
	return "1111"
}

// AttributeUpdatesHandler 更新共享属性
func (c ConfigTest) AttributeUpdatesHandler(m map[string]any) {

}
