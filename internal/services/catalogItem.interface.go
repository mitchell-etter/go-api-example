package services

type CatalogItemService interface {
	GetAll() []CatalogListItem
	GetBySku(string) *CatalogItem
}

type CatalogListItem struct {
	Sku  string `json:"sku"`
	Name string `json:"name"`
}

type CatalogItem struct {
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
