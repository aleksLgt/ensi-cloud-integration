package app

type (
	Options struct {
		Addr, EnsiCloudPrivateToken, EnsiCloudPublicToken, EnsiCloudAddr string
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
			ensiCloudAddr:         opts.EnsiCloudAddr,
			ensiCloudPrivateToken: opts.EnsiCloudPrivateToken,
			ensiCloudPublicToken:  opts.EnsiCloudPublicToken,
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
