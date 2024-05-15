package main

import (
	"go-gateway/tb"
	"go-gateway/util"
	"time"
)

func main() {
	//client, err := modbus.NewClient(&modbus.ClientConfiguration{
	//	URL:      "rtu:///dev/tty.usbserial-B00282SI",
	//	Speed:    9600,
	//	DataBits: 8,
	//	Parity:   modbus.PARITY_ODD,
	//	StopBits: 1,
	//	Timeout:  3 * time.Second,
	//})
	//
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//err = client.Open()
	//if err != nil {
	//	return
	//}
	//err = client.SetUnitId(1)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//values, err := client.ReadRegister(0, modbus.INPUT_REGISTER)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//println(values)

	util.Load()

	tb.Connect()

	for {
		time.Sleep(10 * time.Second)
	}
}
