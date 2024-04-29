package test

import "go-gateway/connectors/base"

type ConfigTest struct {
	base.ConfigBase
	A int `json:"a"`
}
