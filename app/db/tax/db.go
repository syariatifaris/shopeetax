package tax

import (
	"context"
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
func GetTaxRepo(res *repores.RepoResource) Repo {
	repoOnce.Do(func() {
		taxRepo := &taxRepo{
			RepoResource: res,
		}
		repo = taxRepo
	})
	return repo
}

//Repo as tax db repo
type Repo interface {
	GetTaxes(context.Context) ([]*taxmodel.Tax, error)
	db.RepoDB
}

type taxRepo struct {
	*repores.RepoResource
}

//GetTaxes get shopee taxes
//args:
//	ctx: context from above
//returns:
//	list of tax
//	error operation
func (t *taxRepo) GetTaxes(ctx context.Context) ([]*taxmodel.Tax, error) {
	var taxes []*taxmodel.Tax
	err := t.ShopeeDB.SelectContext(ctx, &taxes, getTaxQuery)
	if err != nil {
		return nil, err
	}
	return taxes, nil
}

//BeginTransaction get database transaction
//returns:
//	sqlx's transaction
func (t *taxRepo) BeginTransaction() (*sqlx.Tx, error) {
	return t.ShopeeDB.Beginx()
}
