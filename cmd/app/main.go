package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"ensi-cloud-integration/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	initEnv()
	opts := initOpts()

	service, err := app.NewApp(ctx, app.NewConfig(&opts))

	if err != nil {
		log.Panicf("error when creating a new app: %s", err)
	}

	err = service.ListenAndServe()
	if err != nil {
		log.Panicf("error starting server: %s\n", err)
	}
}

func initEnv() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func initOpts() app.Options {
	options := app.Options{
		Addr: getEnv("DEFAULT_ADDR", ":8082"),
		EnsiCloud: app.EnsiCloudConfig{
			Addr:         getEnv("ENSI_CLOUD_ADDR", ""),
			PrivateToken: getEnv("ENSI_CLOUD_PRIVATE_TOKEN", ""),
			PublicToken:  getEnv("ENSI_CLOUD_PUBLIC_TOKEN", ""),
		},
	}

	return options
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
