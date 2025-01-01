package domain

type SortType string

const (
	SortByRelevance     SortType = "relevance"
	SortByRelevanceDesc SortType = "-relevance"
	SortByName          SortType = "name"
	SortByNameDesc      SortType = "-name"
	SortByPrice         SortType = "price"
	SortByPriceDesc     SortType = "-price"
)

func GetSortTypes() []SortType {
	return []SortType{SortByRelevance, SortByRelevanceDesc, SortByName, SortByNameDesc, SortByPrice, SortByPriceDesc}
}
