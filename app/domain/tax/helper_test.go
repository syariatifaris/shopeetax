package tax

import (
	"reflect"
	"testing"

	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
)

func mockTaxes() map[int64]*taxmodel.Tax {
	return map[int64]*taxmodel.Tax{
		1: &taxmodel.Tax{
			Code:       "Food",
			Expression: "0.1 * ${price}",
			MinPrice:   -1,
		},
		2: &taxmodel.Tax{
			Code:       "Tobacco",
			Expression: "10 + (0.02 * ${price})",
			MinPrice:   -1,
		},
		3: &taxmodel.Tax{
			Code:       "Movie",
			Expression: "0.01 * (${price} - 100)",
			MinPrice:   99,
		},
	}
}

func TestCalculateTax(t *testing.T) {
	type args struct {
		input *TaxableProductInput
		taxes map[int64]*taxmodel.Tax
	}
	tests := []struct {
		name    string
		args    args
		want    *taxmodel.TaxableProduct
		wantErr bool
	}{
		{
			name: "possitive test case tobacco",
			args: args{
				input: &TaxableProductInput{
					InitialPrice:  1000,
					ProductName:   "Dunhill",
					TaxCategoryID: 2,
				},
				taxes: mockTaxes(),
			},
			want: &taxmodel.TaxableProduct{
				Price:         1000,
				ProductName:   "Dunhill",
				TaxCategoryID: 2,
				TaxPrice:      30,
				TotalPrice:    1030,
			},
			wantErr: false,
		},
		{
			name: "possitive test case food",
			args: args{
				input: &TaxableProductInput{
					InitialPrice:  1000,
					ProductName:   "Big Mac",
					TaxCategoryID: 1,
				},
				taxes: mockTaxes(),
			},
			want: &taxmodel.TaxableProduct{
				Price:         1000,
				ProductName:   "Big Mac",
				TaxCategoryID: 1,
				TaxPrice:      100,
				TotalPrice:    1100,
			},
			wantErr: false,
		},
		{
			name: "possitive test case entertainment taxable",
			args: args{
				input: &TaxableProductInput{
					InitialPrice:  150,
					ProductName:   "Movie",
					TaxCategoryID: 3,
				},
				taxes: mockTaxes(),
			},
			want: &taxmodel.TaxableProduct{
				Price:         150,
				ProductName:   "Movie",
				TaxCategoryID: 3,
				TaxPrice:      0.5,
				TotalPrice:    150.5,
			},
			wantErr: false,
		},
		{
			name: "possitive test case entertainment not taxable",
			args: args{
				input: &TaxableProductInput{
					InitialPrice:  80,
					ProductName:   "Movie",
					TaxCategoryID: 3,
				},
				taxes: mockTaxes(),
			},
			want: &taxmodel.TaxableProduct{
				Price:         80,
				ProductName:   "Movie",
				TaxCategoryID: 3,
				TaxPrice:      0,
				TotalPrice:    80,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateTax(tt.args.input, tt.args.taxes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateTax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateTax() = %v, want %v", got, tt.want)
			}
		})
	}
}
