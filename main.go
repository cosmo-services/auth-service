package main

import (
	"main/cmd"
	_ "main/docs"
)

// @title Cosmo Auth API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security BearerAuth

// @BasePath /api/v1/
func main() {
	cmd.StartApp()
}
