package rest

import (
	contract "github.com/aasumitro/goms/internal/bff/domain/contract"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/internal/bff/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FetchStore
// Stores godoc
// @Schemes
// @Summary 	 Store List
// @Description  Get store list and show as JSON data.
// @Tags 		 Stores
// @Accept       json
// @Produce      json
// @Success 200 {object} utils.SuccessRespond{data=[]entity.Store} "OK RESPOND"
// @Failure 500 {object} utils.ErrorRespond "INTERNAL SERVER ERROR RESPOND"
// @Router /api/v1/stores [GET]
func (handler *bffRESTHandler) FetchStore(ctx *gin.Context) {
	data, err := handler.service.AllStore(ctx, nil, nil)
	if err != nil {
		utils.WrapHTTPMessage(ctx, err.Code, err.Message)
		return
	}
	utils.WrapHTTPMessage(ctx, http.StatusOK, data)
}

func (handler *bffRESTHandler) ShowStore(ctx *gin.Context) {
	var (
		id   int
		err  error
		args *contract.WithParam
	)

	if id, err = strconv.Atoi(ctx.Param("id")); err != nil {
		utils.WrapHTTPMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if param := ctx.Request.URL.Query().Get("with"); param == "books" {
		relation := contract.WithRelationID
		args = &relation
	}

	data, errWrapper := handler.service.FirstStore(ctx, args, &entity.Store{ID: uint32(id)})
	if errWrapper != nil {
		utils.WrapHTTPMessage(ctx, errWrapper.Code, errWrapper.Message)
		return
	}
	utils.WrapHTTPMessage(ctx, http.StatusOK, data)
}

func (handler *bffRESTHandler) CreateStore(ctx *gin.Context) {

}

func (handler *bffRESTHandler) EditStore(ctx *gin.Context) {

}

func (handler *bffRESTHandler) DeleteStore(ctx *gin.Context) {

}
