package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jason-shen/gopush/config"
	"github.com/jason-shen/gopush/ent"
	"github.com/jason-shen/gopush/ent/migrate"
	"github.com/jason-shen/gopush/handlers"
	"github.com/jason-shen/gopush/middlewares"
	"github.com/jason-shen/gopush/pkg/utils/logger"
	"github.com/jason-shen/gopush/routes"
	_ "github.com/lib/pq"
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

	conf := config.New()

	client, err := ent.Open("postgres", fmt.Sprintf( "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Name, conf.Database.Password))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := fiber.New()

	middlewares.SetupMiddleware(app)

	handler := handlers.NewHandlers(client, conf)

	routes.SetupApiV1(app, handler)

	port := cfg.Section("general").Key("port").String()
	if port == "" {
		port = "9000"
	}

	addr := flag.String("addr", port, "http service address")

	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))
}
