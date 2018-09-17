package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

func TestInitShopeeHTTPServer(t *testing.T) {
	type args struct {
		ui   *uires.UIResource
		port int
	}
	tests := []struct {
		name string
		args args
		want Server
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitShopeeHTTPServer(tt.args.ui, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitShopeeHTTPServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_Run(t *testing.T) {
	type fields struct {
		ui     *uires.UIResource
		router *httprouter.Router
		port   int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				ui:     tt.fields.ui,
				router: tt.fields.router,
				port:   tt.fields.port,
			}
			if err := h.Run(); (err != nil) != tt.wantErr {
				t.Errorf("httpServer.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_httpServer_registerRouters(t *testing.T) {
	type fields struct {
		ui     *uires.UIResource
		router *httprouter.Router
		port   int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				ui:     tt.fields.ui,
				router: tt.fields.router,
				port:   tt.fields.port,
			}
			h.registerRouters()
		})
	}
}

func Test_httpServer_handle(t *testing.T) {
	type fields struct {
		ui     *uires.UIResource
		router *httprouter.Router
		port   int
	}
	type args struct {
		uiRes           *uires.UIResource
		httpHandlerFunc httpHandlerFunc
		useCase         usecase.HandleUseCase
		subscribers     []usecase.UseCase
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				ui:     tt.fields.ui,
				router: tt.fields.router,
				port:   tt.fields.port,
			}
			if got := h.handle(tt.args.uiRes, tt.args.httpHandlerFunc, tt.args.useCase, tt.args.subscribers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServer.handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_httpServer_handleView(t *testing.T) {
	type fields struct {
		ui     *uires.UIResource
		router *httprouter.Router
		port   int
	}
	type args struct {
		uiRes               *uires.UIResource
		httpHandlerViewFunc httpHandlerViewFunc
		useCase             usecase.HandleUseCase
		subscribers         []usecase.UseCase
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpServer{
				ui:     tt.fields.ui,
				router: tt.fields.router,
				port:   tt.fields.port,
			}
			if got := h.handleView(tt.args.uiRes, tt.args.httpHandlerViewFunc, tt.args.useCase, tt.args.subscribers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpServer.handleView() = %v, want %v", got, tt.want)
			}
		})
	}
}
