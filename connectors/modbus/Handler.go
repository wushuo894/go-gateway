package modbus

import (
	"github.com/simonvetter/modbus"
	"log"
	"sync"
	"time"
)

func (c ConfigModbus) Func(client *modbus.ModbusClient, info InfoModbus, value any) (results any, err error) {
	address := info.Address
	m := map[uint8]func() (results any, err error){
		0x03: func() (results any, err error) {
			return client.ReadRegister(address, modbus.INPUT_REGISTER)
		},
		0x06: func() (results any, err error) {
			return nil, client.WriteRegister(address, value.(uint16))
		},
	}
	return m[info.FunctionCode]()
}

var (
	HandleMap = map[string]*modbus.ModbusClient{}
	Locker    = &sync.Mutex{}
)

func (c ConfigModbus) item(info InfoModbus, value any) (any, error) {
	Locker.Lock()
	defer Locker.Unlock()
	client := HandleMap[c.Port]

	if client == nil {
		newClient, err := modbus.NewClient(&modbus.ClientConfiguration{
			URL:      "rtu://" + c.Port,
			Speed:    c.Baudrate,        // default
			DataBits: c.Databits,        // default, optional
			Parity:   modbus.PARITY_ODD, // default, optional
			StopBits: c.Stopbits,        // default if no parity, optional
			Timeout:  3 * time.Second,
		})
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		client = newClient
		HandleMap[c.Port] = client

		err = client.Open()
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	//defer func(client *modbus.ModbusClient) {
	//	err := client.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}(client)

	err := client.SetUnitId(c.UnitId)
	if err != nil {
		return nil, err
	}
	results, err := c.Func(client, info, value)
	if err != nil {
		return results, err
	}
	return results, nil
}
