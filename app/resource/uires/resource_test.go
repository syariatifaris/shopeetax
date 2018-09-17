package uires

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

func TestUIResource_RenderJSON(t *testing.T) {
	type fields struct {
		UseCaseResource *usecaseres.UseCaseResource
		Router          *httprouter.Router
		ServiceType     *usecase.ServiceType
	}
	type args struct {
		w          http.ResponseWriter
		data       interface{}
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIResource{
				UseCaseResource: tt.fields.UseCaseResource,
				Router:          tt.fields.Router,
				ServiceType:     tt.fields.ServiceType,
			}
			u.RenderJSON(tt.args.w, tt.args.data, tt.args.statusCode)
		})
	}
}

func TestUIResource_GetPostData(t *testing.T) {
	type fields struct {
		UseCaseResource *usecaseres.UseCaseResource
		Router          *httprouter.Router
		ServiceType     *usecase.ServiceType
	}
	type args struct {
		r *http.Request
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIResource{
				UseCaseResource: tt.fields.UseCaseResource,
				Router:          tt.fields.Router,
				ServiceType:     tt.fields.ServiceType,
			}
			if err := u.GetPostData(tt.args.r, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("UIResource.GetPostData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUIResource_CreateUseCaseData(t *testing.T) {
	type fields struct {
		UseCaseResource *usecaseres.UseCaseResource
		Router          *httprouter.Router
		ServiceType     *usecase.ServiceType
	}
	tests := []struct {
		name   string
		fields fields
		want   *usecase.UseCaseData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIResource{
				UseCaseResource: tt.fields.UseCaseResource,
				Router:          tt.fields.Router,
				ServiceType:     tt.fields.ServiceType,
			}
			if got := u.CreateUseCaseData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UIResource.CreateUseCaseData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUIResource_RenderView(t *testing.T) {
	type fields struct {
		UseCaseResource *usecaseres.UseCaseResource
		Router          *httprouter.Router
		ServiceType     *usecase.ServiceType
	}
	type args struct {
		opt    *ViewOption
		writer http.ResponseWriter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIResource{
				UseCaseResource: tt.fields.UseCaseResource,
				Router:          tt.fields.Router,
				ServiceType:     tt.fields.ServiceType,
			}
			u.RenderView(tt.args.opt, tt.args.writer)
		})
	}
}

func TestUIResource_SetupView(t *testing.T) {
	type fields struct {
		UseCaseResource *usecaseres.UseCaseResource
		Router          *httprouter.Router
		ServiceType     *usecase.ServiceType
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UIResource{
				UseCaseResource: tt.fields.UseCaseResource,
				Router:          tt.fields.Router,
				ServiceType:     tt.fields.ServiceType,
			}
			u.SetupView()
		})
	}
}
