package tax

import (
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/shopeetax/app/db"
	taxmodel "github.com/syariatifaris/shopeetax/app/model/tax"
	"github.com/syariatifaris/shopeetax/app/resource/repores"
)

var repo Repo
var repoOnce sync.Once

//GetTaxRepo gets tax repo once
//args:
//	res: repo resources
//returns:
//	TaxRepository
func GetTaxRepo(ctx context.Context, res *repores.RepoResource) Repo {
	repoOnce.Do(func() {
		t := &taxRepo{
			RepoResource: res,
		}
		t.Holder = new(Holder)
		err := t.holder(ctx)
		if err != nil {
			panic(fmt.Sprint("unable to initialize holder data ", err.Error()))
		}
		repo = t
	})
	return repo
}

//Repo as tax db repo
type Repo interface {
	GetTaxes(context.Context) ([]*taxmodel.Tax, error)
	GetComponents(context.Context) ([]*taxmodel.Component, error)
	InsertTaxableProduct(context.Context, *sqlx.Tx, *taxmodel.TaxableProduct) error
	GetHolders() *Holder
	db.RepoDB
}

type taxRepo struct {
	*repores.RepoResource
	*Holder
}

//GetTaxes get shopee taxes
//args:
//	ctx: context from above
//returns:
//	list of tax
//	error operation
func (t *taxRepo) GetTaxes(ctx context.Context) ([]*taxmodel.Tax, error) {
	var taxes []*taxmodel.Tax
	err := t.ShopeeDB.SelectContext(ctx, &taxes, getTaxesQuery)
	if err != nil {
		return nil, err
	}
	return taxes, nil
}

func (t *taxRepo) GetComponents(ctx context.Context) ([]*taxmodel.Component, error) {
	var components []*taxmodel.Component
	err := t.ShopeeDB.SelectContext(ctx, &components, getComponentsQuery)
	if err != nil {
		return nil, err
	}
	return components, nil
}

func (t *taxRepo) InsertTaxableProduct(ctx context.Context, tx *sqlx.Tx, tp *taxmodel.TaxableProduct) error {
	insertQuery := tx.Rebind(insertTaxableProductQuery)
	res, err := tx.Exec(insertQuery,
		tp.TaxCategoryID,
		tp.ProductName,
		tp.Price,
		tp.TaxPrice,
		tp.TotalPrice)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	tp.ID = lastID
	return nil
}

func (t *taxRepo) GetHolders() *Holder {
	return t.Holder
}

//BeginTransaction get database transaction
//returns:
//	sqlx's transaction
func (t *taxRepo) BeginTransaction() (*sqlx.Tx, error) {
	return t.ShopeeDB.Beginx()
}

func (t *taxRepo) holder(ctx context.Context) error {
	components, err := t.GetComponents(ctx)
	if err != nil {
		return err
	}
	t.Components = []*taxmodel.Component{}
	t.Components = components
	taxes, err := t.GetTaxes(ctx)
	if err != nil {
		return err
	}
	t.Taxes = []*taxmodel.Tax{}
	t.Taxes = taxes
	t.mapRules()
	return nil
}
