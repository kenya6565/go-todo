package model

type Company struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}
