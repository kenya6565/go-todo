package config

import (
	"go-todo/model"
	"go-todo/utils"
	"log"

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
		Static:    cfg.Section("web").Key("static").String(),
	}
}
