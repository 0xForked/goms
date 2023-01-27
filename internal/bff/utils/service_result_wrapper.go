package utils

import (
	"database/sql"
	"net/http"
)

type ServiceErrorData struct {
	Code    int
	Message any
}

func WrapDataRows[T any](data []*T, err error) (valueData []*T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	return data, errData
}

func WrapDataRow[T any](data *T, err error) (valueData *T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	return data, errData
}

func WrapPrimitiveValue[T any](data T, err error) (valueData T, errData *ServiceErrorData) {
	errData = validateErrorMessage(err)

	return data, errData
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
