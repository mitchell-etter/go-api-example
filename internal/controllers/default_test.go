package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRoot(t *testing.T) {
	controller := NewDefaultController()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	controller.GetRoot(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
}
