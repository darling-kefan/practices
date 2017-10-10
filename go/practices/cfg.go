package main

import (
	"github.com/darling-kefan/toolkit/cfg"
	"log"
	"strings"
)

func main() {
	cfg.Init(cfg.FileProvider{
		Filename: ".env",
	})
	log.Println(cfg.Log())
	log.Println(cfg.MustString("DB_DATABASE"))

	var config = make(map[string]string)
	line := `AUTH="shouqiang"`
	index := strings.Index(line, "=")

	key, val := line[:index], line[index+1:]

	// add the item to the config
	if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
		config[key] = strings.Trim(val, "\"")
	} else if strings.HasPrefix(val, "'") && strings.HasSuffix(val, "'") {
		log.Println("'''''''''''''''''")
		config[key] = strings.Trim(val, "'")
	} else {
		config[key] = val
	}
	log.Println(config)
}
