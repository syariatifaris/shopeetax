package tax

import (
	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
)

//InsertTaxRequest request structure
type InsertTaxRequest struct {
	TaxProduct *TaxableProductInput `json:"tax_product"`
}

//InsertTaxResponse response structure
type InsertTaxResponse struct {
	TaxableProduct *taxmodel.TaxableProduct `json:"taxable_product"`
}

//GetTaxableProductsResponse response structure
type GetTaxableProductsResponse struct {
	TaxableProducts []*taxmodel.TaxableProduct `json:"taxable_products"`
	Summary         *TaxableProductSummary     `json:"summary"`
}

//TaxableProductInput structure
type TaxableProductInput struct {
	ProductName   string  `json:"product_name"`
	TaxCategoryID int64   `json:"tax_category_id"`
	InitialPrice  float64 `json:"initial_price"`
}

//TaxableProductSummary structure
type TaxableProductSummary struct {
	TotalAmount    float64 `json:"total_amount"`
	TotalTaxAmount float64 `json:"total_tax_amount"`
	GrandTotal     float64 `json:"grand_total"`
}
