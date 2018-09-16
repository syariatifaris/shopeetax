package tax

var (
	getTaxesQuery             = `SELECT * FROM shopee_tax`
	getComponentsQuery        = `SELECT * FROM shopee_tax_component`
	getTaxableProductsQuery        = `SELECT a.*, b.code as "tax_code" FROM shopee_taxable_product as a INNER JOIN shopee_tax as b ON a.tax_category_id = b.id`
	insertTaxableProductQuery = `
		INSERT INTO shopee_taxable_product (
			tax_category_id,
			product_name,   
			price,
			tax_price, 
			total_price) VALUES (?, ?, ?, ?, ?)`
)
