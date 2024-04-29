package modbus

import (
	"go-gateway/connectors/base"
)

type ConfigModbus struct {
	base.ConfigBase
	UnitId           int          `json:"unitId"`
	Port             string       `json:"port"`
	Baudrate         int          `json:"baudrate"`
	Databits         int          `json:"databits"`
	Stopbits         int          `json:"stopbits"`
	Parity           string       `json:"parity"`
	Timeseries       []InfoModbus `json:"timeseries"`
	Rpc              []InfoModbus `json:"rpc"`
	AttributeUpdates []InfoModbus `json:"attributeUpdates"`
}

type InfoModbus struct {
	Tag string `json:"tag"`
	//Type         reflect.Type
	FunctionCode int     `json:"functionCode"`
	ObjectsCount uint16  `json:"objectsCount"`
	Address      uint16  `json:"address"`
	Take         float32 `json:"take"`
}
