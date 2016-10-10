package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Hauptstruct, wird vom Package gefordert
// in DB werden die einzelnen Profile reingeschrieben
type tomlConfig struct {
	DB map[string]Database
}

// Profil eines DB-Eintrags
type Database struct {
	Host string
	Name string
	User string
	Pwd  string
}

// Methode zum auslesen eines Profiles nach Name
func (t tomlConfig) getProfile(name string) Database {
	return t.DB[name]
}

func main() {

	// Definition des Config-Object
	var config tomlConfig

	// TOML-File einlesen, in Variable config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	// Nur ein Profil abfragen
	ce := config.getProfile("ce")
	fmt.Printf("Host: %s, Name: %s, User: %s, Pwd: %s\n", ce.Host, ce.Name, ce.User, ce.Pwd)
}
