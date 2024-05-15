package modbus

import (
	"go-gateway/connectors/base"
)

type ConfigModbus struct {
	base.ConfigBase
	UnitId           uint8        `json:"unitId"`
	Port             string       `json:"port"`
	Baudrate         uint         `json:"baudrate"`
	Databits         uint         `json:"databits"`
	Stopbits         uint         `json:"stopbits"`
	Parity           string       `json:"parity"`
	Timeseries       []InfoModbus `json:"timeseries"`
	Rpc              []InfoModbus `json:"rpc"`
	AttributeUpdates []InfoModbus `json:"attributeUpdates"`
}

type InfoModbus struct {
	Tag string `json:"tag"`
	//Type         reflect.Type
	FunctionCode uint8   `json:"functionCode"`
	ObjectsCount uint16  `json:"objectsCount"`
	Address      uint16  `json:"address"`
	Take         float32 `json:"take"`
}
