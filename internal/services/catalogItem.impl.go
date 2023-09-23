package services

func NewCatalogItemService() CatalogItemService {
	return &catalogItemServiceImpl{
		allItems: []CatalogItem{
			{
				Sku:         "001",
				Name:        "Toaster",
				Description: "It browns bread.",
			},
			{
				Sku:         "002",
				Name:        "Microwave Oven",
				Description: "It heats food unevenly.",
			},
		},
	}
}

type catalogItemServiceImpl struct {
	allItems []CatalogItem
}

func (svc *catalogItemServiceImpl) GetAll() []CatalogListItem {
	listItems := make([]CatalogListItem, 0, len(svc.allItems))

	for _, item := range svc.allItems {
		listItems = append(listItems, CatalogListItem{
			Name: item.Name,
			Sku:  item.Sku,
		})
	}

	return listItems
}

func (svc *catalogItemServiceImpl) GetBySku(sku string) *CatalogItem {
	if len(sku) == 0 {
		return nil
	}

	for _, item := range svc.allItems {
		if item.Sku == sku {
			return &item
		}
	}

	return nil
}
