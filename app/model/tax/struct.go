package tax

//Tax table database structure
type Tax struct {
	ID          int64   `json:"id" db:"id"`
	Code        string  `json:"code" db:"code"`
	Description string  `json:"desc" db:"description"`
	Expression  string  `json:"expression" db:"expression"`
	MinPrice    float64 `json:"min_price" db:"min_price"`
}

//Component table database structure
type Component struct {
	ID          int64  `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Description string `json:"desc" db:"description"`
	VarType     string `json:"var_type" db:"var_type"`
}

//TaxableProduct taxable product table database structure
type TaxableProduct struct {
	ID            int64   `json:"id" db:"id"`
	ProductName   string  `json:"product_name" db:"product_name"`
	TaxCategoryID int64   `json:"tax_category_id" db:"tax_category_id"`
	Price         float64 `json:"price" db:"price"`
	TaxPrice      float64 `json:"tax_price" db:"tax_price"`
	TotalPrice    float64 `json:"total_price" db:"total_price"`
	TaxCode       string  `json:"tax_code,omitempty" db:"tax_code"`
}
