package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
)

func TestUseCaseData_Cast(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		target interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.Cast(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.Cast() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_castHTTPRequest(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		target interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.castHTTPRequest(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.castHTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_castExecRequest(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		target interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.castExecRequest(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.castExecRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_castNSQBody(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		target interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.castNSQBody(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.castNSQBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_castSubscribeData(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		target interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.castSubscribeData(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.castSubscribeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_NotifySubscribers(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		ctx     context.Context
		timeout time.Duration
		res     *usecaseres.UseCaseResource
		data    interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.NotifySubscribers(tt.args.ctx, tt.args.timeout, tt.args.res, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.NotifySubscribers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCaseData_GetNotifyResultByName(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if gotResult := u.GetNotifyResultByName(tt.args.ctx, tt.args.name); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UseCaseData.GetNotifyResultByName() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestUseCaseData_GetSubscriberByName(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   UseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if got := u.GetSubscriberByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCaseData.GetSubscriberByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCaseData_setSubscribeData(t *testing.T) {
	type fields struct {
		ServiceType   *ServiceType
		HTTPData      interface{}
		NSQBody       []byte
		Subscribers   []UseCase
		SubscribeData *SubscribeData
		Language      string
		Arguments     []string
	}
	type args struct {
		data interface{}
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
			u := &UseCaseData{
				ServiceType:   tt.fields.ServiceType,
				HTTPData:      tt.fields.HTTPData,
				NSQBody:       tt.fields.NSQBody,
				Subscribers:   tt.fields.Subscribers,
				SubscribeData: tt.fields.SubscribeData,
				Language:      tt.fields.Language,
				Arguments:     tt.fields.Arguments,
			}
			if err := u.setSubscribeData(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseData.setSubscribeData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
