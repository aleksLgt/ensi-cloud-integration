package app

type (
	EnsiCloudConfig struct {
		Addr, PrivateToken, PublicToken string
	}

	Options struct {
		Addr      string
		EnsiCloud EnsiCloudConfig
	}

	configEnsiCloudService struct {
		ensiCloudAddr, ensiCloudPrivateToken, ensiCloudPublicToken string
	}

	path struct {
		indexProducts,
		indexCategories,
		searchCatalog,
		searchRecommendedQueryProducts,
		searchRecommendedProducts,
		searchCrossSellProducts string
	}

	Config struct {
		addr string
		configEnsiCloudService
		path path
	}
)

func NewConfig(opts *Options) *Config {
	return &Config{
		addr: opts.Addr,
		configEnsiCloudService: configEnsiCloudService{
			ensiCloudAddr:         opts.EnsiCloud.Addr,
			ensiCloudPrivateToken: opts.EnsiCloud.PrivateToken,
			ensiCloudPublicToken:  opts.EnsiCloud.PublicToken,
		},
		path: path{
			indexProducts:                  "POST /indexes/products",
			indexCategories:                "POST /indexes/categories",
			searchCatalog:                  "POST /catalog/search",
			searchRecommendedQueryProducts: "POST /adviser/recommendation-query-products:search",
			searchRecommendedProducts:      "POST /adviser/recommendation-products:search",
			searchCrossSellProducts:        "POST /adviser/cross-sell-products:search",
		},
	}
}
