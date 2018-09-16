package api

import (
	"github.com/syariatifaris/shopeetax/app/resource/uires"
	"github.com/syariatifaris/shopeetax/app/usecase"
	"net/http"
)

//InsertTaxHandler inserts the new tax
//args:
//	ui: ui resouece
//	request: http request object
//	use case: use case implmentation business for
//	subscribers usecase which subscribe the event
//returns:
//	interface: return response
//	error: state operation error
func InsertTaxHandler(ui *uires.UIResource, request *http.Request, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase) (interface{}, error) {
	data := ui.CreateUseCaseData()
	err := ui.GetPostData(request, &data.HTTPData)
	if err != nil {
		return nil, err
	}
	data.Subscribers = subscribers
	return useCase(request.Context(), ui.UseCaseResource, data)
}

//GetTaxableProductsHandler gets taxable product data
//args:
//	ui: ui resouece
//	request: http request object
//	use case: use case implmentation business for
//	subscribers usecase which subscribe the event
//returns:
//	interface: return response
//	error: state operation error
func GetTaxableProductsHandler(ui *uires.UIResource, request *http.Request, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase) (interface{}, error) {
	return useCase(request.Context(), ui.UseCaseResource, ui.CreateUseCaseData())
}

//IndexPage show test page UI
func IndexPage(ui *uires.UIResource, request *http.Request, writer http.ResponseWriter, useCase usecase.HandleUseCase, subscribers ...usecase.UseCase) {
	ui.RenderView(&uires.ViewOption{
		Module:    "tax",
		PageTitle: "Tax Calculator",
		ViewModel: struct {
			VisitorName string `json:"visitor_name"`
			VisitorAge  int    `json:"visitor_age"`
		}{
			VisitorName: "Faris M Syariati",
			VisitorAge:  27,
		},
	}, writer)
}
