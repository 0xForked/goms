package utils_test

import (
	"database/sql"
	"errors"
	"github.com/aasumitro/goms/internal/bff/domain/entity"
	"github.com/aasumitro/goms/internal/bff/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateDataRow(t *testing.T) {
	type args[T any] struct {
		data *T
		err  error
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantValueData *T
		wantErrData   *utils.ServiceErrorData
	}
	tests := []testCase[entity.Store]{
		{
			name: "Validate Row Should Success",
			args: args[entity.Store]{
				data: &entity.Store{ID: 1, Name: "ipsum"},
				err:  nil,
			},
			wantValueData: &entity.Store{ID: 1, Name: "ipsum"},
			wantErrData:   nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[entity.Store]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceErrorData{Code: 500, Message: "LOREM"},
		},
		{
			name: "Validate Row Should Not Found",
			args: args[entity.Store]{
				data: nil,
				err:  sql.ErrNoRows,
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceErrorData{Code: 404, Message: sql.ErrNoRows.Error()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.WrapDataRow(tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "ValidateDataRow(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "ValidateDataRow(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}

func TestValidateDataRows(t *testing.T) {
	type args[T any] struct {
		data []*T
		err  error
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantValueData []*T
		wantErrData   *utils.ServiceErrorData
	}
	tests := []testCase[entity.Store]{
		{
			name: "Validate Row Should Success",
			args: args[entity.Store]{
				data: []*entity.Store{
					{ID: 1, Name: "ipsum"},
					{ID: 2, Name: "lorem"},
				},
				err: nil,
			},
			wantValueData: []*entity.Store{
				{ID: 1, Name: "ipsum"},
				{ID: 2, Name: "lorem"},
			},
			wantErrData: nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[entity.Store]{
				data: nil,
				err:  errors.New("LOREM"),
			},
			wantValueData: nil,
			wantErrData:   &utils.ServiceErrorData{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.WrapDataRows(tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}

func TestValidatePrimitiveValue(t *testing.T) {
	type args[T any] struct {
		data T
		err  error
	}
	type testCase[T any] struct {
		name          string
		args          args[T]
		wantValueData T
		wantErrData   *utils.ServiceErrorData
	}
	tests := []testCase[int]{
		{
			name: "Validate int Should Success",
			args: args[int]{
				data: 1,
				err:  nil,
			},
			wantValueData: 1,
			wantErrData:   nil,
		},
		{
			name: "Validate Row Should Error",
			args: args[int]{
				data: 0,
				err:  errors.New("LOREM"),
			},
			wantValueData: 0,
			wantErrData:   &utils.ServiceErrorData{Code: 500, Message: "LOREM"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValueData, gotErrData := utils.WrapPrimitiveValue(tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantValueData, gotValueData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
			assert.Equalf(t, tt.wantErrData, gotErrData, "ValidateDataRows(%v, %v)", tt.args.data, tt.args.err)
		})
	}
}
