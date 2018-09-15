package repores

import (
	"github.com/jmoiron/sqlx"
)

//RepoResource holds object dependencies for Repository
//i.e: database, redis, apicall, grpc
type RepoResource struct {
	ShopeeDB *sqlx.DB
}
