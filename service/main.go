package main

import (
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tpisntgod/Agenda/service/server"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	ser := server.NewServer()
	ser.Run(":" + port)
}
