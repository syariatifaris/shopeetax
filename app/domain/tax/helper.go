package tax

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"strings"

	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
)

//CalculateTax calculates tax based on product input and its rule
func CalculateTax(input *TaxableProductInput, taxes map[int64]*taxmodel.Tax) (*taxmodel.TaxableProduct, error) {
	rule, ok := taxes[input.TaxCategoryID]
	if !ok {
		return nil, fmt.Errorf("unable to find tax with category id %d", input.TaxCategoryID)
	}
	taxProduct := &taxmodel.TaxableProduct{
		Price:         input.InitialPrice,
		TaxCategoryID: input.TaxCategoryID,
		ProductName:   input.ProductName,
	}
	if input.InitialPrice <= rule.MinPrice {
		taxProduct.TaxPrice = 0
	} else {
		replacer := strings.NewReplacer(
			"${", "",
			"}", "",
		)
		rule.Expression = replacer.Replace(rule.Expression)
		expression, err := govaluate.NewEvaluableExpression(rule.Expression)
		if err != nil {
			return nil, err
		}
		parameters := make(map[string]interface{}, 1)
		parameters["price"] = input.InitialPrice
		result, err := expression.Evaluate(parameters)
		if err != nil {
			return nil, err
		}
		taxProduct.TaxPrice = result.(float64)
	}
	taxProduct.TotalPrice = taxProduct.Price + taxProduct.TaxPrice
	return taxProduct, nil
}
