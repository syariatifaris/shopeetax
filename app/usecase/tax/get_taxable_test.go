package tax

import (
	"context"
	"reflect"
	"testing"

	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

func TestGetHandleGetTaxableProduct(t *testing.T) {
	tests := []struct {
		name string
		want usecase.HandleUseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHandleGetTaxableProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHandleGetTaxableProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTaxableUseCase_Name(t *testing.T) {
	tests := []struct {
		name string
		g    *getTaxableUseCase
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &getTaxableUseCase{}
			if got := g.Name(); got != tt.want {
				t.Errorf("getTaxableUseCase.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTaxableUseCase_HandleUseCase(t *testing.T) {
	type args struct {
		ctx  context.Context
		res  *usecaseres.UseCaseResource
		data *usecase.UseCaseData
	}
	tests := []struct {
		name    string
		g       *getTaxableUseCase
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &getTaxableUseCase{}
			got, err := g.HandleUseCase(tt.args.ctx, tt.args.res, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTaxableUseCase.HandleUseCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTaxableUseCase.HandleUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTaxableUseCase_Notify(t *testing.T) {
	type args struct {
		ctx  context.Context
		res  *usecaseres.UseCaseResource
		data *usecase.UseCaseData
	}
	tests := []struct {
		name    string
		g       *getTaxableUseCase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &getTaxableUseCase{}
			if err := g.Notify(tt.args.ctx, tt.args.res, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("getTaxableUseCase.Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getTaxableUseCase_NotifyResult(t *testing.T) {
	tests := []struct {
		name string
		g    *getTaxableUseCase
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &getTaxableUseCase{}
			if got := g.NotifyResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTaxableUseCase.NotifyResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
