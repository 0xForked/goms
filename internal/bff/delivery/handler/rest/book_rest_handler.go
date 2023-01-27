package rest

import (
	"github.com/aasumitro/goms/internal/bff/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FetchBook
// Books godoc
// @Schemes
// @Summary 	 Book List
// @Description  Get book list and show as JSON data.
// @Tags 		 Books
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]entity.Book} "OK RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/books [GET]
func (handler *bffRESTHandler) FetchBook(ctx *gin.Context) {
	utils.WrapHTTPMessage(ctx, http.StatusOK, gin.H{"hey": "Hello World from Book Data"})
}

func (handler *bffRESTHandler) ShowBook(ctx *gin.Context) {

}

func (handler *bffRESTHandler) CreateBook(ctx *gin.Context) {

}

func (handler *bffRESTHandler) EditBook(ctx *gin.Context) {

}

func (handler *bffRESTHandler) DeleteBook(ctx *gin.Context) {

}
