package modbus

import (
	"encoding/binary"
	"github.com/grid-x/modbus"
	"log"
	"strconv"
	"sync"
	"time"
)

func (c ConfigModbus) Func(client modbus.Client, info InfoModbus, value any) (results []byte, err error) {
	address := info.Address
	count := info.ObjectsCount

	m := map[int]func() (results []byte, err error){
		modbus.FuncCodeReadDiscreteInputs: func() (results []byte, err error) {
			return client.ReadDiscreteInputs(address, count)
		},
		modbus.FuncCodeReadCoils: func() (results []byte, err error) {
			return client.ReadCoils(address, count)
		},
		modbus.FuncCodeWriteSingleCoil: func() (results []byte, err error) {
			return client.WriteSingleCoil(address, value.(uint16))
		},
		modbus.FuncCodeWriteMultipleCoils: func() (results []byte, err error) {
			return client.WriteMultipleCoils(address, count, value.([]byte))
		},
		modbus.FuncCodeReadInputRegisters: func() (results []byte, err error) {
			return client.ReadInputRegisters(address, count)
		},
		modbus.FuncCodeReadHoldingRegisters: func() (results []byte, err error) {
			return client.ReadHoldingRegisters(address, count)
		},
		modbus.FuncCodeWriteSingleRegister: func() (results []byte, err error) {
			return client.WriteSingleRegister(address, value.(uint16))
		},
		modbus.FuncCodeWriteMultipleRegisters: func() (results []byte, err error) {
			return client.WriteMultipleRegisters(address, count, value.([]byte))
		},
		modbus.FuncCodeReadFIFOQueue: func() (results []byte, err error) {
			return client.ReadFIFOQueue(address)
		},
	}
	return m[info.FunctionCode]()
}

var (
	HandleMap = map[string]modbus.Client{}
	LockerMap = map[string]*sync.Mutex{}
	Locker    = &sync.Mutex{}
)

func (c ConfigModbus) item(info InfoModbus, value any) (any, error) {
	Locker.Lock()
	client := HandleMap[c.Port+","+strconv.Itoa(c.UnitId)]
	locker := LockerMap[c.Port]

	if locker == nil {
		locker = &sync.Mutex{}
		LockerMap[c.Port] = locker
	}

	if client == nil {
		handler := modbus.NewRTUClientHandler(c.Port)
		handler.BaudRate = c.Baudrate
		handler.DataBits = c.Databits
		handler.Parity = c.Parity
		handler.StopBits = c.Stopbits
		handler.SlaveID = byte(c.UnitId)
		handler.Timeout = 3 * time.Second
		err := handler.Connect()
		if err != nil {
			Locker.Unlock()
			log.Fatalln("Connect Error:", err)
			return nil, err
		}
		log.Println("Connect ", c.Port)
		client = modbus.NewClient(handler)
		HandleMap[c.Port+","+strconv.Itoa(c.UnitId)] = client
	}
	Locker.Unlock()

	locker.Lock()
	results, err := c.Func(client, info, value)
	locker.Unlock()
	if err != nil {
		return nil, err
	}
	v := binary.BigEndian.Uint16(results)

	var ok = float32(v)

	take := info.Take
	if take == 0 {
		take = 1
	}

	resultStr := strconv.FormatFloat(float64(ok*take), 'f', 2, 64)

	return resultStr, nil
}
