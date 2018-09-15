package usecaseres

import (
	"github.com/syariatifaris/shopeetax/app/db/tax"
	"github.com/syariatifaris/shopeetax/app/infra/config"
)

//UseCaseResource holders
type UseCaseResource struct {
	TaxRepo tax.Repo
	Config  *config.ConfigurationData
}
