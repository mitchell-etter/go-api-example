package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCatalogItems_NoneFound(t *testing.T) {
	svc := &catalogItemServiceImpl{}

	items := svc.GetAll()

	assert.Empty(t, items)
}

func TestGetAllCatalogItems_OneFound(t *testing.T) {
	svc := &catalogItemServiceImpl{
		allItems: []CatalogItem{
			{
				Sku:         "test",
				Name:        "test item",
				Description: "Item for testing",
			},
		},
	}

	items := svc.GetAll()

	assert.Len(t, items, 1)
}

func TestGetBySku_EmptySku(t *testing.T) {
	svc := &catalogItemServiceImpl{}

	item := svc.GetBySku("")

	assert.Nil(t, item)
}

func TestGetBySku_NotFound(t *testing.T) {
	svc := &catalogItemServiceImpl{
		allItems: []CatalogItem{
			{
				Sku:         "test",
				Name:        "test item",
				Description: "Item for testing",
			},
		},
	}

	item := svc.GetBySku("nonexistent")

	assert.Nil(t, item)
}

func TestGetBySku_Found(t *testing.T) {
	svc := &catalogItemServiceImpl{
		allItems: []CatalogItem{
			{
				Sku:         "test",
				Name:        "test item",
				Description: "Item for testing",
			},
		},
	}

	item := svc.GetBySku("test")

	assert.NotNil(t, item)
}
