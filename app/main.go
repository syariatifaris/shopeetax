package main

import (
	"fmt"
	"os"

	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	taxdb "github.com/syariatifaris/shopeetax/app/db/tax"
	"github.com/syariatifaris/shopeetax/app/infra/config"
	"github.com/syariatifaris/shopeetax/app/infra/db"
	"github.com/syariatifaris/shopeetax/app/resource/repores"
	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/server"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

//main runs application
func main() {
	httpServer := server.InitShopeeHTTPServer(initUIResouece(), 9090)
	errChan := make(chan error)
	go func() {
		fmt.Println("server started on: 0.0.0.0:9090")
		errChan <- httpServer.Run()
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		fmt.Println("server terminated due error", err.Error())
	case <-term:
		fmt.Println("signal terminate detected, bye!")
	}
}

//initUIResouece initialize ui resorce
//returns:
//	UI Resouece
func initUIResouece() *uires.UIResource {
	cfg := config.NewConfiguration()
	repoResource := initRepoResource(cfg)
	return &uires.UIResource{
		UseCaseResource: initUseCaseResource(repoResource, cfg),
		Router:          httprouter.New(),
		ServiceType:     &usecase.HTTPServiceType,
	}
}

//initRepoResource initialize ui resorce
//returns:
//	Repo Resource
func initRepoResource(cfg *config.ConfigurationData) *repores.RepoResource {
	shopeeDB := db.NewSqlxConnection(cfg.Database)
	return &repores.RepoResource{
		ShopeeDB: shopeeDB,
	}
}

//initUseCaseResource initiialize use case resource
//args:
//	repoResource: repo resource injection
//returns:
//	UsecaseResource: usecase reosource object
func initUseCaseResource(res *repores.RepoResource, cfg *config.ConfigurationData) *usecaseres.UseCaseResource {
	return &usecaseres.UseCaseResource{
		Config:  cfg,
		TaxRepo: taxdb.GetTaxRepo(res),
	}
}
