package util

import (
	"encoding/json"
	"github.com/titanous/json5"
	"go-gateway/connectors/base"
	"go-gateway/connectors/modbus"
	"go-gateway/connectors/test"
	"go-gateway/entity"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Gateway struct {
}

var Config = entity.GatewayConfig{
	ThingsBoard: entity.ThingsBoardConfig{},
	Connectors:  []*base.ConfigBase{},
}

var ConfigFuncMap = map[string]func(bt []byte, cb base.ConfigBase) base.ConnectorBase{
	"MODBUS": func(bt []byte, cb base.ConfigBase) base.ConnectorBase {
		config := modbus.ConfigModbus{}
		bytes, err := json.Marshal(cb)
		if err != nil {
			log.Fatalln(err)
		}
		err = json5.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatalln(err)
		}
		err = json5.Unmarshal(bt, &config)
		if err != nil {
			log.Fatalln(err)
		}
		return config
	},
	"TEST": func(bt []byte, cb base.ConfigBase) base.ConnectorBase {
		config := test.ConfigTest{}
		bytes, err := json.Marshal(cb)
		if err != nil {
			log.Fatalln(err)
		}
		err = json5.Unmarshal(bytes, &config)
		if err != nil {
			log.Fatalln(err)
		}
		err = json5.Unmarshal(bt, &config)
		if err != nil {
			log.Fatalln(err)
		}
		return config
	},
}

func Load() {
	config, err := os.Open("config")
	if err != nil {
		log.Fatalln(err)
	}
	dir, err := config.ReadDir(-1)
	if err != nil {
		log.Fatalln(err)
	}

	var tbGatewayConfigFile = ""

	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		ext := filepath.Ext(file.Name())
		if !slices.Contains([]string{".json", ".json5"}, ext) {
			continue
		}
		if !strings.HasPrefix(file.Name(), "tb_gateway") {
			continue
		}
		tbGatewayConfigFile = file.Name()
	}

	if len(tbGatewayConfigFile) < 1 {
		log.Fatalln("找不到配置文件 tb_gateway.json5")
	}

	file, err := os.Open("config/" + tbGatewayConfigFile)
	if err != nil {
		log.Fatalln("无法打开文件:", err)
		return
	}

	byteValue, err := io.ReadAll(file)
	err = json5.Unmarshal(byteValue, &Config)
	if err != nil {
		log.Fatalln(err)
	}
}
