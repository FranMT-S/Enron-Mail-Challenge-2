package main

import (
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/config"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/server"
)

func main() {
	config.LoadConfig()
	server := server.NewServer(config.CFG.ServerPort)
	server.Start()
}
