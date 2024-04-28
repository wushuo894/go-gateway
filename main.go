package main

import (
	"go-gateway/tb"
	"go-gateway/util"
	"time"
)

func main() {
	util.Load()

	tb.Connect()
	//go func() {
	//	configTest := test.ConfigTest{
	//		base.ConfigBase{
	//			DeviceName: "test",
	//			DeviceId:   "",
	//			FileName:   "",
	//			Connector:  nil,
	//		},
	//	}
	//
	//}()

	for {
		time.Sleep(10 * time.Second)
	}

}
