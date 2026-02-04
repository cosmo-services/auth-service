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

// @BasePath /api/v2/auth
func main() {
	cmd.StartApp()
}
