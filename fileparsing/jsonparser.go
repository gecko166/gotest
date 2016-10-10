package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Definition der einzelnen Profile
// -> also ein Hash, Key ist ein String, der Inhalt
// des Hashes ist ein Struct mit der Struktur von Db
type Profile map[string]Db

// Dieses Struct beschreibt die Struktur eines DB-Profils
type Db struct {
	Host, Name, User, Pwd string
}

// Methode zum auslesen eines Hashes mit dem Key 'name'
// Rueckgabe eines Structs Db
func (p Profile) getProfile(name string) Db {
	return p[name]
}

func main() {
	// Configfile einlesen -> Rawbytes
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// eine Variable vom Type config definieren
	var config Profile

	// File parsen und in config schreiben
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	// Da die Variable config vom Type Profile ist, und
	// der Type Profile eine Methode getProfile hat,
	// kann ein bestimmter Key angewendet werden
	fmt.Println("CE:", config.getProfile("ce"))
}
