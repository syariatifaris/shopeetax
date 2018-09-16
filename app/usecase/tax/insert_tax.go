package tax

import (
	"context"
	"errors"
	"fmt"
	"sync"

	taxdomain "github.com/syariatifaris/shopeetax/app/domain/tax"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

var insertTaxUC usecase.UseCase
var insertTaxOnce sync.Once

//GetHandleInsertTax gets handle create tax main logic
func GetHandleInsertTax() usecase.HandleUseCase {
	insertTaxOnce.Do(func() {
		insertTaxUC = &insertTaxUseCase{}
	})
	return insertTaxUC.HandleUseCase
}

//insertTaxUseCase contract
type insertTaxUseCase struct{}

//Name gets the use case name
//returns:
//	name of use case
func (c *insertTaxUseCase) Name() string {
	return "InsertTaxUseCase"
}

//NotifyResult gets the notify result
//returns:
//	gets the notify result
func (c *insertTaxUseCase) NotifyResult() interface{} {
	return nil
}

//HandleUseCase handle use case main function
//args:
//	ctx: context from request
//	res: use case resource
//	data: use case data
func (c *insertTaxUseCase) HandleUseCase(ctx context.Context, res *usecaseres.UseCaseResource, data *usecase.UseCaseData) (interface{}, error) {
	var request *taxdomain.InsertTaxRequest
	err := data.Cast(&request)
	if err != nil {
		return nil, err
	}
	if err = c.validate(request); err != nil {
		return nil, err
	}
	taxable, err := taxdomain.CalculateTax(request.TaxProduct, res.TaxRepo.GetHolders().TaxRules)
	if err != nil {
		return nil, err
	}
	tx, err := res.TaxRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	err = res.ExecuteTransaction(ctx, tx, func() error {
		return res.TaxRepo.InsertTaxableProduct(ctx, tx, taxable)
	})
	if err != nil {
		return nil, err
	}
	code, err := taxdomain.GetTaxCategoryCode(request.TaxProduct.TaxCategoryID, res.TaxRepo.GetHolders().TaxRules)
	if err != nil {
		return nil, err
	}
	taxable.TaxCode = code
	return &taxdomain.InsertTaxResponse{
		TaxableProduct: taxable,
	}, nil
}

//Notify notifies this usecase when something happen from other use case
//args:
//	ctx context
//	res usecase resource
//return error state
func (c *insertTaxUseCase) Notify(ctx context.Context, res *usecaseres.UseCaseResource, data *usecase.UseCaseData) error {
	fmt.Println("notify on insert tax use case")
	return nil
}

func (c *insertTaxUseCase) validate(request *taxdomain.InsertTaxRequest) error {
	var errs []error
	if request == nil {
		return errors.New("unable to get request, request nil")
	}
	if request.TaxProduct.InitialPrice <= 0 {
		errs = append(errs, errors.New("invalid initial price"))
	}
	if len(request.TaxProduct.ProductName) == 0 {
		errs = append(errs, errors.New("invalid product name"))
	}
	if request.TaxProduct.TaxCategoryID <= 0 {
		errs = append(errs, errors.New("invalid tax type"))
	}
	if len(errs) > 0 {
		return fmt.Errorf("%+v", errs)
	}
	return nil
}
