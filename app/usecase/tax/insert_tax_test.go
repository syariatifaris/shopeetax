package tax

import (
	"context"
	"reflect"
	"testing"

	taxdomain "github.com/syariatifaris/shopeetax/app/domain/tax"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

func TestGetHandleInsertTax(t *testing.T) {
	tests := []struct {
		name string
		want usecase.HandleUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHandleInsertTax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHandleInsertTax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertTaxUseCase_Name(t *testing.T) {
	tests := []struct {
		name string
		c    *insertTaxUseCase
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &insertTaxUseCase{}
			if got := c.Name(); got != tt.want {
				t.Errorf("insertTaxUseCase.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertTaxUseCase_NotifyResult(t *testing.T) {
	tests := []struct {
		name string
		c    *insertTaxUseCase
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &insertTaxUseCase{}
			if got := c.NotifyResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertTaxUseCase.NotifyResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertTaxUseCase_HandleUseCase(t *testing.T) {
	type args struct {
		ctx  context.Context
		res  *usecaseres.UseCaseResource
		data *usecase.UseCaseData
	}
	tests := []struct {
		name    string
		c       *insertTaxUseCase
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &insertTaxUseCase{}
			got, err := c.HandleUseCase(tt.args.ctx, tt.args.res, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertTaxUseCase.HandleUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertTaxUseCase.HandleUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertTaxUseCase_Notify(t *testing.T) {
	type args struct {
		ctx  context.Context
		res  *usecaseres.UseCaseResource
		data *usecase.UseCaseData
	}
	tests := []struct {
		name    string
		c       *insertTaxUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &insertTaxUseCase{}
			if err := c.Notify(tt.args.ctx, tt.args.res, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("insertTaxUseCase.Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_insertTaxUseCase_validate(t *testing.T) {
	type args struct {
		request *taxdomain.InsertTaxRequest
	}
	tests := []struct {
		name    string
		c       *insertTaxUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &insertTaxUseCase{}
			if err := c.validate(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("insertTaxUseCase.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
