package entity

import "go-gateway/connectors/base"

type GatewayConfig struct {
	Connectors  []base.ConfigBase `json:"connectors"`
	ThingsBoard ThingsBoardConfig `json:"things_board"`
	LogLevel    string            `json:"log_level"`
}
