package model

import "time"

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

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}
