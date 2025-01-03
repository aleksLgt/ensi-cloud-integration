package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"ensi-cloud-integration/internal/app/closer"
	"ensi-cloud-integration/internal/app/http/adviser/crossSellProducts"
	"ensi-cloud-integration/internal/app/http/adviser/recommendationProducts"
	"ensi-cloud-integration/internal/app/http/adviser/recommendationQueryProducts"
	"ensi-cloud-integration/internal/app/http/catalog"
	"ensi-cloud-integration/internal/app/http/indexes/categories"
	"ensi-cloud-integration/internal/app/http/indexes/products"
	"ensi-cloud-integration/internal/clients/ensiCloud"
	"ensi-cloud-integration/internal/domain"
	searchCrossSellProducts "ensi-cloud-integration/internal/service/ensiCloud/adviser/crossSellProducts"
	searchRecommendationProducts "ensi-cloud-integration/internal/service/ensiCloud/adviser/recommendationProducts"
	searchRecommendationQueryProducts "ensi-cloud-integration/internal/service/ensiCloud/adviser/recommendationQueryProducts"
	searchCatalog "ensi-cloud-integration/internal/service/ensiCloud/catalog"
	indexCategories "ensi-cloud-integration/internal/service/ensiCloud/indexes/categories"
	indexProducts "ensi-cloud-integration/internal/service/ensiCloud/indexes/products"
	"ensi-cloud-integration/pkg/logger"
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
		IndexProducts(ctx context.Context, request *domain.IndexProductsRequest) (*domain.IndexProductsResponse, error)
		IndexCategories(ctx context.Context, request *domain.IndexCategoriesRequest) (*domain.IndexCategoriesResponse, error)
		SearchCatalog(ctx context.Context, request *domain.SearchCatalogRequest) (*domain.SearchCatalogResponse, error)

		SearchCrossSellProducts(
			ctx context.Context,
			request *domain.SearchCrossSellProductsRequest,
		) (*domain.SearchCrossSellProductsResponse, error)

		SearchRecommendationProducts(
			ctx context.Context,
			request *domain.SearchRecommendationProductsRequest,
		) (*domain.SearchRecommendationProductsResponse, error)

		SearchRecommendationQueryProducts(
			ctx context.Context,
			request *domain.SearchRecommendationQueryProductsRequest,
		) (*domain.SearchRecommendationQueryProductsResponse, error)
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

	_, err := logger.New()
	if err != nil {
		panic(err)
	}

	loggerCustom, err := logger.With("service", "ensi-cloud-integration")
	if err != nil {
		// panic, as an error can only occur if the logger is not initialized
		panic(err)
	}

	ctx = logger.ToContext(ctx, loggerCustom)

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
			Handler:           mux,
			ReadHeaderTimeout: 3 * time.Second,
		},
		ensiCloudClient: newEnsiCloudClient,
		closer:          &closer.Closer{},
	}, nil
}

func (a *App) ListenAndServe() error {
	a.registerRoutes()

	a.closer.Add(a.server.Shutdown)

	// When calling a.closer.Close after the server is stopped, we wait a few seconds for the completion of
	// third-party processes
	a.closer.Add(func(ctx context.Context) error {
		time.Sleep(3 * time.Second)

		return nil
	})

	go func() {
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Panicw(a.ctx, "listen and serve failed", "error", err)
		}
	}()

	logger.Infow(a.ctx, fmt.Sprintf("listening on %s", a.config.addr))
	<-a.ctx.Done()

	logger.Infow(a.ctx, "shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.closer.Close(shutdownCtx); err != nil {
		return fmt.Errorf("closer failed: %v", err)
	}

	return nil
}

func (a *App) registerRoutes() {
	a.mux.Handle(
		a.config.path.indexProducts,
		products.NewIndexProductsHandler(
			indexProducts.New(a.ensiCloudClient),
			a.config.path.indexProducts,
		),
	)

	a.mux.Handle(
		a.config.path.indexCategories,
		categories.NewIndexCategoriesHandler(
			indexCategories.New(a.ensiCloudClient),
			a.config.path.indexCategories,
		),
	)

	a.mux.Handle(
		a.config.path.searchCatalog,
		catalog.NewSearchCatalogHandler(
			searchCatalog.New(a.ensiCloudClient),
			a.config.path.searchCatalog,
		),
	)

	a.mux.Handle(
		a.config.path.searchCrossSellProducts,
		crossSellProducts.NewSearchCrossSellProductsHandler(
			searchCrossSellProducts.New(a.ensiCloudClient),
			a.config.path.searchCrossSellProducts,
		),
	)

	a.mux.Handle(
		a.config.path.searchRecommendedProducts,
		recommendationProducts.NewSearchRecommendationProductsHandler(
			searchRecommendationProducts.New(a.ensiCloudClient),
			a.config.path.searchRecommendedProducts,
		),
	)

	a.mux.Handle(
		a.config.path.searchRecommendedQueryProducts,
		recommendationQueryProducts.NewSearchRecommendationQueryProductsHandler(
			searchRecommendationQueryProducts.New(a.ensiCloudClient),
			a.config.path.searchRecommendedQueryProducts,
		),
	)
}
