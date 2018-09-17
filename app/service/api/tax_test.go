package api

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

func TestInsertTaxHandler(t *testing.T) {
	type args struct {
		ui          *uires.UIResource
		request     *http.Request
		useCase     usecase.HandleUseCase
		subscribers []usecase.UseCase
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InsertTaxHandler(tt.args.ui, tt.args.request, tt.args.useCase, tt.args.subscribers...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertTaxHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertTaxHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTaxableProductsHandler(t *testing.T) {
	type args struct {
		ui          *uires.UIResource
		request     *http.Request
		useCase     usecase.HandleUseCase
		subscribers []usecase.UseCase
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTaxableProductsHandler(tt.args.ui, tt.args.request, tt.args.useCase, tt.args.subscribers...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaxableProductsHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaxableProductsHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexPage(t *testing.T) {
	type args struct {
		ui          *uires.UIResource
		request     *http.Request
		writer      http.ResponseWriter
		useCase     usecase.HandleUseCase
		subscribers []usecase.UseCase
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndexPage(tt.args.ui, tt.args.request, tt.args.writer, tt.args.useCase, tt.args.subscribers...)
		})
	}
}
