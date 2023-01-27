package utils

import (
	"database/sql"
	"net/http"
)

type ServiceErrorData struct {
	Code    int
	Message any
}

func WrapDataRow[T any](data *T, err error) (valueData *T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	if data == nil {
		errData = whenDataIsNil()
	}

	return data, errData
}

func WrapDataRows[T any](data []*T, err error) (valueData []*T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	if data == nil {
		errData = whenDataIsNil()
	}

	return data, errData
}

func WrapPrimitiveValue[T any](data T, err error) (valueData T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	return data, errData
}

func whenDataIsNil() *ServiceErrorData {
	return &ServiceErrorData{
		Code:    http.StatusNotFound,
		Message: "data you're looking for not found",
	}
}

func validateErrorMessage(err error) *ServiceErrorData {
	var errData *ServiceErrorData

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errData = &ServiceErrorData{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}
		default:
			errData = &ServiceErrorData{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}

	return errData
}
