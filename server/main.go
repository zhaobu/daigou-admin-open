package main

import (
	"daigou-admin/core"
	"daigou-admin/initialize"
)

//"runtime"

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	initialize.Init()
	core.Run()
}
