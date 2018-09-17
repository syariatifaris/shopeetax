package tax

import (
	"context"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
	"github.com/syariatifaris/shopeetax/app/resource/repores"
)

func TestGetTaxRepo(t *testing.T) {
	type args struct {
		ctx context.Context
		res *repores.RepoResource
	}
	tests := []struct {
		name string
		args args
		want Repo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTaxRepo(tt.args.ctx, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTaxRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_GetTaxes(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*taxmodel.Tax
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			got, err := t.GetTaxes(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.GetTaxes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRepo.GetTaxes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_GetComponents(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*taxmodel.Component
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			got, err := t.GetComponents(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.GetComponents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRepo.GetComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_GetTaxableProducts(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*taxmodel.TaxableProduct
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			got, err := t.GetTaxableProducts(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.GetTaxableProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRepo.GetTaxableProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_InsertTaxableProduct(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	type args struct {
		ctx context.Context
		tx  *sqlx.Tx
		tp  *taxmodel.TaxableProduct
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
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			if err := t.InsertTaxableProduct(tt.args.ctx, tt.args.tx, tt.args.tp); (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.InsertTaxableProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taxRepo_GetHolders(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	tests := []struct {
		name   string
		fields fields
		want   *Holder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			if got := t.GetHolders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRepo.GetHolders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_BeginTransaction(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	tests := []struct {
		name    string
		fields  fields
		want    *sqlx.Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			got, err := t.BeginTransaction()
			if (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.BeginTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taxRepo.BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taxRepo_holder(t *testing.T) {
	type fields struct {
		RepoResource *repores.RepoResource
		Holder       *Holder
	}
	type args struct {
		ctx context.Context
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
			t := &taxRepo{
				RepoResource: tt.fields.RepoResource,
				Holder:       tt.fields.Holder,
			}
			if err := t.holder(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("taxRepo.holder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
