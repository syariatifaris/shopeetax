package server

import (
	"fmt"

	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/service/api"
	"github.com/syariatifaris/shopeetax/app/usecase"
	taxusecase "github.com/syariatifaris/shopeetax/app/usecase/tax"
)

//InitShopeeHTTPServer creates the shopee tax framework UI resouece
//args:
//	ui: UI resouece
//	port: http api port
//returns:
//	Server: server implementation
func InitShopeeHTTPServer(ui *uires.UIResource, port int) Server {
	return &httpServer{
		ui:     ui,
		port:   port,
		router: ui.Router,
	}
}

//Run runs the server
func (h *httpServer) Run() error {
	h.registerRouters()
	//run server
	return gracehttp.Serve(&http.Server{
		Addr:    fmt.Sprint("0.0.0.0:", h.port),
		Handler: h.router,
	})
}

func (h *httpServer) registerRouters() {
	h.router.HandlerFunc(http.MethodPost, "/shopee/tax/insert", h.handle(h.ui, api.InsertTaxHandler, taxusecase.GetHandleInsertTax()))
}

//handle handles middleware to run handler-usecase coupled
//args:
//	uiRes: ui resouece
//	httpHandleFunc: anon function for http handler
//	useCase: use case anon function
//	subscribers event subscribed usecase
//returns:
//	http.HandlerFunc: anon function for http handler
func (h *httpServer) handle(uiRes *uires.UIResource, httpHandlerFunc httpHandlerFunc, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		response, err := httpHandlerFunc(uiRes, request, useCase, subscribers...)
		if err != nil {
			h.ui.RenderJSON(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		h.ui.RenderJSON(writer, response, http.StatusOK)
	}
}
