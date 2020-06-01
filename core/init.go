package core

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/spf13/viper"
	"log"
	"tawny/tawny"
)

var SubscriptionStore = map[string]map[string]*chan (tawny.Message){}
var PresenceStore = map[string]map[string]*chan (bool){}

var Db *badger.DB

func Init() {
	var err error
	Db, err = badger.Open(badger.DefaultOptions(viper.GetString(BADGER_PATH)))
	if err != nil {
		log.Fatal(err)
	}
}
