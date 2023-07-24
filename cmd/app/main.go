package main

import (
	"LavaPanel/internal/model"
	"LavaPanel/internal/service"
)

func main() {
	model.InitDB()
	service.InitCmdInfo()
	service.InitRouter()
}
