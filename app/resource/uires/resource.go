package uires

import (
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

//UIResource structure
type UIResource struct {
	UseCaseResource *usecaseres.UseCaseResource
	Router          *httprouter.Router
	ServiceType     *usecase.ServiceType
}

//RenderJSON renders the json to web response
//args:
//	w: http response writer
//	data: passed data to be written
//	statusCode: http status code
func (u *UIResource) RenderJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	byteData, _ := json.Marshal(data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(byteData)
}

//GetPostData gets the post data from http request
//args:
//	r: http request
//	v: target interface where the data post will be passed
//returns:
//	error: error state
func (u *UIResource) GetPostData(r *http.Request, v interface{}) error {
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("method %s not allowed %d", r.Method, http.StatusMethodNotAllowed)
}

//CreateUseCaseData creates usec case data and add the service type definition
func (u *UIResource) CreateUseCaseData() *usecase.UseCaseData {
	return &usecase.UseCaseData{
		ServiceType: u.ServiceType,
	}
}
