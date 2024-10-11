package main

import (
	"log"

	"github.com/nidoqueen1/article-api/app"
	"github.com/spf13/viper"
)

func main() {
	server := app.New()
	if err := server.Run(viper.GetString("server.addr")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
