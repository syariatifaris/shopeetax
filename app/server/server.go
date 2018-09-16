package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

//httpHandlerFunc abstraction of http handler
type httpHandlerFunc func(ui *uires.UIResource, request *http.Request, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase) (interface{}, error)

//httpHandlerViewFunc abstraction of http handler with view
type httpHandlerViewFunc func(ui *uires.UIResource, request *http.Request, writer http.ResponseWriter, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase)

//Server contract
type Server interface {
	Run() error
}

//httpServer structure as Server implementation for http api
type httpServer struct {
	ui     *uires.UIResource
	router *httprouter.Router
	port   int
}
