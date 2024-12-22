package app

import (
	"context"
	"ensi-cloud-integration/internal/app/http/indexes"
	indexProducts "ensi-cloud-integration/internal/service/ensiCloud/indexes/products"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"ensi-cloud-integration/internal/app/closer"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	mux interface {
		Handle(pattern string, handler http.Handler)
	}

	server interface {
		ListenAndServe() error
		Close() error
		Shutdown(ctx context.Context) error
	}

	ensiCloudClient interface {
		IndexProducts(ctx context.Context) error
	}

	App struct {
		ctx             context.Context
		config          *Config
		mux             mux
		server          server
		ensiCloudClient ensiCloudClient
		closer          *closer.Closer
	}
)

func NewApp(ctx context.Context, config *Config) (*App, error) {
	mux := http.NewServeMux()

	newEnsiCloudClient, err := ensiCloud.New(
		config.ensiCloudAddr,
		config.ensiCloudPrivateToken,
		config.ensiCloudPublicToken,
	)

	if err != nil {
		return nil, fmt.Errorf("the creation of a new ensi cloud client failed: %w", err)
	}

	return &App{
		ctx:    ctx,
		config: config,
		mux:    mux,
		server: &http.Server{
			Addr:              config.addr,
			ReadHeaderTimeout: 3 * time.Second, // TODO
		},
		ensiCloudClient: newEnsiCloudClient,
		closer:          &closer.Closer{},
	}, nil
}

func (a *App) ListenAndServe() error {
	a.mux.Handle(
		a.config.path.indexProducts,
		indexes.NewIndexProductsHandler(
			indexProducts.New(a.ensiCloudClient),
			a.config.path.indexProducts,
		),
	)

	a.closer.Add(a.server.Shutdown)

	// When calling a.closer.Close after the server is stopped, we wait a few seconds for the completion of
	// third-party processes
	a.closer.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)

		return nil
	})

	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("listen and serve failed %s", err)
		}
	}()

	log.Printf("listening on %s", a.config.addr)
	<-a.ctx.Done()

	log.Print("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.closer.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer failed: %v", err)
	}

	return nil
}
