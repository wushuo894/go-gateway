package util

import (
	"fmt"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"go-gateway/connectors/base"
	"go-gateway/connectors/modbus"
	"go-gateway/connectors/test"
	"go-gateway/entity"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
)

type Gateway struct {
}

var Config = &entity.GatewayConfig{
	ThingsBoard: entity.ThingsBoardConfig{},
	Connectors:  []base.ConfigBase{},
}

func Load() {
	m := map[string]reflect.Type{
		"MODBUS": reflect.TypeOf(modbus.ConfigModbus{}),
		"TEST":   reflect.TypeOf(test.ConfigTest{}),
	}
	println(m)
	config, _ := os.Open("config")
	dir, _ := config.ReadDir(-1)

	var tbGatewayConfigFile = ""

	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		exts := []string{".json", ".json5"}
		ext := filepath.Ext(file.Name())
		if !slices.Contains(exts, ext) {
			continue
		}
		if !strings.HasPrefix(file.Name(), "tb_gateway") {
			continue
		}
		tbGatewayConfigFile = file.Name()
	}

	if len(tbGatewayConfigFile) < 1 {
		println("找不到配置文件 tb_gateway.json5")
		os.Exit(1)
		return
	}

	file, err := os.Open("config/" + tbGatewayConfigFile)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}

	byteValue, err := ioutil.ReadAll(file)
	err1 := json5.Unmarshal(byteValue, &Config)
	if err1 != nil {
		println(err1.Error())
		os.Exit(1)
		return
	}
	connectors := Config.Connectors
	for _, configBase := range connectors {
		c := reflect.New(m[configBase.DeviceType])
		connectorBase := c.Elem().Interface().(base.ConnectorBase)
		println(m[configBase.DeviceType].Name())
		configBase.Connector = &connectorBase
		go connectorBase.Run()
	}
}
