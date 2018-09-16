package tax

var (
	getTaxesQuery             = `SELECT * FROM shopee_tax`
	getComponentsQuery        = `SELECT * FROM shopee_tax_component`
	insertTaxableProductQuery = `
		INSERT INTO shopee_taxable_product (
			tax_category_id,
			product_name,   
			price,
			tax_price, 
			total_price) VALUES (?, ?, ?, ?, ?)`
)
