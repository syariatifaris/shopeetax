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
