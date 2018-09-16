package tax

import (
	"context"
	"sync"

	taxdomain "github.com/syariatifaris/shopeetax/app/domain/tax"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

var getTaxableUC usecase.UseCase
var getTaxableOnce sync.Once

type getTaxableUseCase struct{}

//GetHandleGetTaxableProduct gets handle create tax main logic
func GetHandleGetTaxableProduct() usecase.HandleUseCase {
	getTaxableOnce.Do(func() {
		getTaxableUC = &getTaxableUseCase{}
	})
	return getTaxableUC.HandleUseCase
}

func (g *getTaxableUseCase) Name() string {
	return "GetTaxableUseCase"
}

func (g *getTaxableUseCase) HandleUseCase(ctx context.Context, res *usecaseres.UseCaseResource, data *usecase.UseCaseData) (interface{}, error) {
	products, err := res.TaxRepo.GetTaxableProducts(ctx)
	return &taxdomain.GetTaxableProductsResponse{
		TaxableProducts: products,
	}, err
}

func (g *getTaxableUseCase) Notify(ctx context.Context, res *usecaseres.UseCaseResource, data *usecase.UseCaseData) error {
	panic("not implemented")
}

func (g *getTaxableUseCase) NotifyResult() interface{} {
	panic("not implemented")
}
