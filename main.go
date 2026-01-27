package main

import (
	"log"
	"main/cmd"
	_ "main/docs"
)

// @title Cosmo Auth API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security BearerAuth
func main() {
	if err := cmd.StartApp(); err != nil {
		log.Fatal(err)
	}
}
