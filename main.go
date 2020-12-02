package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/gopush/pkg/utils/logger"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main()  {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		logger.Errorf("Fail to read file: %v", err)
		os.Exit(1)
	}

	app := fiber.New()

	port := cfg.Section("general").Key("port").String()
	if port == "" {
		port = "9000"
	}

	addr := flag.String("addr", port, "http service address")

	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
