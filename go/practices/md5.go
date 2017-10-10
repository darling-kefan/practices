package main

import (
	"fmt"
	"crypto/md5"
)

func main() {
	username := "admin"
	password := "tvmining@2017"
	authSalt := "{Kj]1O~MRng=0US)"
	passwordMd5 := fmt.Sprintf("%x", md5.Sum([]byte(username+password)))
	ret := fmt.Sprintf("%x", md5.Sum([]byte(passwordMd5+authSalt)))
	fmt.Println(passwordMd5, ret)
}
