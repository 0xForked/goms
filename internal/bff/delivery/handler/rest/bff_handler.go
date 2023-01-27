package rest

import (
	"github.com/aasumitro/goms/internal/bff/domain/contract"
	"github.com/gin-gonic/gin"
)

type bffRESTHandler struct {
	service contract.IBFFService
}

func NewBFFRegistrarHandler(service contract.IBFFService, router *gin.RouterGroup) {
	handler := bffRESTHandler{service: service}
	// STORES ROUTER LIST
	router.GET("/stores", handler.FetchStore)
	router.GET("/stores/:id", handler.ShowStore)
	router.POST("/stores", handler.CreateStore)
	router.PATCH("/stores/:id", handler.EditStore)
	router.DELETE("/stores/:id", handler.DeleteStore)
	// BOOKS ROUTER LIST
	router.GET("/books", handler.FetchBook)
	router.GET("/books/:id", handler.ShowBook)
	router.POST("/books", handler.CreateBook)
	router.PATCH("/books/:id", handler.EditBook)
	router.DELETE("/books/:id", handler.DeleteBook)
}
