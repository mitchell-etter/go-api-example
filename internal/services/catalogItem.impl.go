package services

func NewCatalogItemService() CatalogItemService {
	return &catalogItemServiceImpl{}
}

type catalogItemServiceImpl struct {
}

func (svc *catalogItemServiceImpl) GetAll() []CatalogListItem {
	return []CatalogListItem{}
}

func (svc *catalogItemServiceImpl) GetBySku(sku string) *CatalogItem {
	return &CatalogItem{}
}
