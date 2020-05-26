package core

import (
	"github.com/spf13/viper"
)

const (
	BADGER_PATH  = "BADGER_PATH"
	GATEWAY_PORT = "GATEWAY_PORT"
)

func InitConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tawny")
	viper.SetDefault(BADGER_PATH, "db-temp/badger")
	viper.SetDefault(GATEWAY_PORT, "8900")
}
