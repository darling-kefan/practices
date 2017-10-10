package main

import (
	"fmt"
)

var createTablesStatements = []string{
	`CREATE DATABASE IF NOT EXISTS maintain DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE maintain;`,
	    `CREATE TABLE IF NOT EXISTS projects (                                                                                                                                 `+"`id`"+` int(11) unsigned NOT NULL AUTO_INCREMENT,                                                                                                               ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci`,
}


func main() {
	fmt.Println(createTablesStatements)
}
