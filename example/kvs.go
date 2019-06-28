package main

import (
	"github.com/sevennt/wzap"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file")
	}

	logs := viper.GetStringMap("app.logger")
	for l := range logs {
		outputs := viper.GetStringMap("app.logger." + l)
		outputs["name"] = l
		wzap.Register(l, wzap.New(wzap.WithOutputKVs([]interface{}{outputs})))
	}

	wzap.WInfo("mobile", "123")
	wzap.WInfo("service", "1234")
}
