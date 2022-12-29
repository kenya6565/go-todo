package config

import (
	"crypto/sha1"
	"fmt"
	"go-todo/model"
	"go-todo/utils"
	"log"

	"github.com/google/uuid"

	"gopkg.in/go-ini/ini.v1"
)

// global variable for the use of other files
var Config model.ConfigList

// execute before main()
func init() {
	// read config.ini
	LoadConfig()

	// output a logfile
	utils.LoggingSettings(Config.LogFile)

}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = model.ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8090"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}

// generate uuid used when creating users table in sqlite3
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// generate password used when creating users table in sqlite3
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
