package core

import (
	"github.com/spf13/viper"
)

const (
	ADMIN_KEY    = "ADMIN_KEY"
	BADGER_PATH  = "BADGER_PATH"
	HTTPS_ENABLE = "HTTPS_ENABLE"
	HTTPS_DOMAIN = "HTTPS_DOMAIN"
	HTTPS_EMAIL  = "HTTPS_EMAIL"
	HTTPS_PORT   = "HTTPS_PORT"
	HTTP_PORT    = "HTTP_PORT"
	GRPC_PORT    = "GRPC_PORT"
)

func InitConfig() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tawny")
	viper.SetDefault(BADGER_PATH, "db-temp/badger")
	viper.SetDefault(HTTPS_ENABLE, false)
	viper.SetDefault(GRPC_PORT, "4000")
	if viper.GetBool(HTTPS_ENABLE) {
		viper.SetDefault(HTTPS_PORT, "443")
		viper.SetDefault(HTTP_PORT, "80")
	} else {
		viper.SetDefault(HTTP_PORT, "8900")
	}
	key := viper.GetString(ADMIN_KEY)
	if key == "" {
		panic("A TAWNY_ADMIN_KEY need to be define")
	}
}
