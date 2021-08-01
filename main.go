package main

import (
	_ "net/http"
	"submit/logger"
)

func main() {

	logger.Info("Submit v%s initializing...", "1.22.2")
	server := NewServer()

	server.Run(":3001")

}
