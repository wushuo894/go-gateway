package util

import (
	"fmt"
	"github.com/titanous/json5"
	"go-gateway/connectors/base"
	"go-gateway/connectors/modbus"
	"go-gateway/connectors/test"
	"go-gateway/entity"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
)

type Gateway struct {
}

var Config = entity.GatewayConfig{
	ThingsBoard: entity.ThingsBoardConfig{},
	Connectors:  []base.ConfigBase{},
}

var ConfigMap = map[string]reflect.Type{
	"MODBUS": reflect.TypeOf(modbus.ConfigModbus{}),
	"TEST":   reflect.TypeOf(test.ConfigTest{}),
}

func Load() {
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

	byteValue, err := io.ReadAll(file)
	err1 := json5.Unmarshal(byteValue, &Config)
	if err1 != nil {
		println(err1.Error())
		os.Exit(1)
		return
	}
}
