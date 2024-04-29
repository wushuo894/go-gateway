package modbus

import (
	"encoding/binary"
	"github.com/goburrow/modbus"
	"log"
	"strconv"
)

func (c ConfigModbus) Func(client modbus.Client, info InfoModbus, value any) (results []byte, err error) {
	address := info.Address
	count := info.ObjectsCount

	if err != nil {
		return nil, err
	}

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

func (c ConfigModbus) item(info InfoModbus, value any) (any, error) {
	handler := modbus.NewRTUClientHandler(c.Port)
	handler.BaudRate = c.Baudrate
	handler.DataBits = c.Databits
	handler.Parity = c.Parity
	handler.StopBits = c.Stopbits
	handler.SlaveId = byte(c.UnitId)

	err := handler.Connect()
	if err != nil {
		return nil, err
	}
	defer func(handler *modbus.RTUClientHandler) {
		err := handler.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(handler)
	client := modbus.NewClient(handler)
	results, err := c.Func(client, info, value)
	if err != nil {
		return nil, err
	}

	var ok = float32(binary.BigEndian.Uint16(results))

	resultStr := strconv.FormatFloat(float64(ok*info.Take), 'f', 2, 64)
	return resultStr, nil
}
