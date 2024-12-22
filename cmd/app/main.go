package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ensi-cloud-integration/internal/app"
)

const (
	defaultAddr           = ":8082"
	ensiCloudAddr         = "https://cloud-api-master-dev-cloud.ensi.tech"
	ensiCloudPrivateToken = "someToken"
	ensiCloudPublicToken  = "someToken"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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

func initOpts() app.Options {
	options := app.Options{}

	// TODO тут разобраться с этой конструкцией
	flag.StringVar(
		&options.Addr,
		"addr",
		defaultAddr,
		fmt.Sprintf("server address, default: %q", defaultAddr),
	)

	flag.StringVar(
		&options.EnsiCloudAddr,
		"ensi_cloud_addr",
		ensiCloudAddr,
		fmt.Sprintf("ensi cloud address, default: %q", ensiCloudAddr),
	)

	flag.StringVar(
		&options.EnsiCloudPrivateToken,
		"ensi_cloud_private_token",
		ensiCloudPrivateToken,
		fmt.Sprintf("ensi cloud private token, default: %q", ensiCloudPrivateToken),
	)

	flag.StringVar(
		&options.EnsiCloudPublicToken,
		"ensi_cloud_public_token",
		ensiCloudPublicToken,
		fmt.Sprintf("ensi cloud public token, default: %q", ensiCloudPublicToken),
	)

	flag.Parse()

	return options
}
