package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-api-example/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListCatalogItems_NoneFound(t *testing.T) {
	controller, _ := newTestController()
	context, recorder := newTestHttpContext()

	controller.ListCatalogItems(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []services.CatalogListItem
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err, "Error occurred deserializing response")
	assert.Empty(t, response, "Expecting empty results")
}

func TestListCatalogItems_OneFound(t *testing.T) {
	controller, service := newTestController()
	service.catalogListItems = append(service.catalogListItems, services.CatalogListItem{})
	context, recorder := newTestHttpContext()

	controller.ListCatalogItems(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response []services.CatalogListItem
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err, "Error occurred deserializing response")
	assert.Len(t, response, 1)
}

func TestGetCatalogItem_NotFound(t *testing.T) {
	controller, _ := newTestController()
	context, recorder := newTestHttpContext()
	context.Params = append(context.Params, gin.Param{Key: "id", Value: "anything"})

	controller.GetCatalogItem(context)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func TestGetCatalogItem_Found(t *testing.T) {
	controller, service := newTestController()
	service.catalogItem = &services.CatalogItem{}
	context, recorder := newTestHttpContext()
	context.Params = append(context.Params, gin.Param{Key: "id", Value: "anything"})

	controller.GetCatalogItem(context)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var response services.CatalogListItem
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err, "Error occurred deserializing response")
}

func newTestController() (*CatalogItemsController, *mockCatalogItemService) {
	mockService := &mockCatalogItemService{}
	var service services.CatalogItemService = mockService
	controller := NewCatalogItemsController(&service)
	return controller, mockService
}

func newTestHttpContext() (*gin.Context, *httptest.ResponseRecorder) {
	responseRecorder := httptest.NewRecorder()
	httpContext, _ := gin.CreateTestContext(responseRecorder)
	return httpContext, responseRecorder
}

type mockCatalogItemService struct {
	catalogListItems []services.CatalogListItem
	catalogItem      *services.CatalogItem
}

func (service *mockCatalogItemService) GetAll() []services.CatalogListItem {
	return service.catalogListItems
}

func (service *mockCatalogItemService) GetBySku(sku string) *services.CatalogItem {
	return service.catalogItem
}
