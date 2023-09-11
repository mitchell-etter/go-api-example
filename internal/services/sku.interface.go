package services

type SkuService interface {
	GetAll() []SkuItem
}

type SkuItem struct {
	Sku  string `json:"sku"`
	Name string `json:"name"`
}
