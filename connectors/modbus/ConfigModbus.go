package modbus

import (
	"go-gateway/connectors/base"
	"reflect"
)

type ConfigModbus struct {
	base.ConfigBase
	UnitId           int32
	Port             string
	Baudrate         int64
	Databits         int32
	Stopbits         int32
	Parity           string
	Timeseries       []InfoModbus
	Rpc              []InfoModbus
	AttributeUpdates []InfoModbus
}

type InfoModbus struct {
	Tag          string
	Type         reflect.Type
	FunctionCode int32
	ObjectsCount int32
	Address      int32
}
