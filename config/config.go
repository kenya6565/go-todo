package config

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

// global variable for the use of other files
var Config ConfigList