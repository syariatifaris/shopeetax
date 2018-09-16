package uires

import (
	"fmt"
	"runtime"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/syariatifaris/shopeetax/app/infra/format"
	"github.com/syariatifaris/shopeetax/app/resource/usecaseres"
	"github.com/syariatifaris/shopeetax/app/usecase"
)

const (
	relPath         = "./"
	assetsRelPath   = "./assets"
	siteDir         = "/site/"
	staticSiteDir   = "/assets/site/"
	staticViewDir   = "assets/site/view/"
	staticMasterDir = "master/"
)

//UIResource structure
type UIResource struct {
	UseCaseResource *usecaseres.UseCaseResource
	Router          *httprouter.Router
	ServiceType     *usecase.ServiceType
}

//Template holds customizeable view template
type Template struct {
	ViewHTMLFile string
	ViewJSFile   string
	ViewCSSFile  string
}

//MasterPage master file structure
//props:
//	JSPageFile: complete relative path for js file
//	CSSPageFile: complete relative path for css file
//	JSObjects: parsed js object (as json object)
type MasterPage struct {
	JSPageFile  string
	CSSPageFile string
	JSObjects   string
}

//ViewOption consists view components
//props:
//	Module: specific render module
//	ViewModel: passing arguments to view
//	JSMode: passing arguments to js
type ViewOption struct {
	Module         string
	SiteTitle      string
	PageTitle      string
	ViewModel      interface{}
	JSModel        interface{}
	CustomTemplate *Template
	MasterPage
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

//RenderView renders the template
//args:
//	opt: View options
//	writer: http response writer
func (u *UIResource) RenderView(opt *ViewOption, writer http.ResponseWriter) {
	var fileName string
	if opt.CustomTemplate != nil && opt.CustomTemplate.ViewHTMLFile != "" {
		fileName = opt.CustomTemplate.ViewHTMLFile
	} else {
		fpcs := make([]uintptr, 1)
		n := runtime.Callers(2, fpcs)
		if n == 0 {
			writer.Write([]byte(fmt.Sprint("unable to get runtime caller")))
			return
		}
		fun := runtime.FuncForPC(fpcs[0] - 1)
		if fun == nil {
			writer.Write([]byte(fmt.Sprint("unable to get handler function")))
			return
		}
		funcs := strings.Split(fun.Name(), ".")
		if len(funcs) == 0 {
			writer.Write([]byte(fmt.Sprint("invalid handler function runtime call")))
			return
		}
		token := funcs[len(funcs)-1]
		fileName = strings.ToLower(format.ToUnderScore(token))
	}
	fileNameExt := fmt.Sprint(fileName, ".html")
	var moduleDir string
	if opt.Module != "" {
		moduleDir = fmt.Sprint(staticViewDir, opt.Module, "/")
	} else {
		moduleDir = staticViewDir
	}
	masterFiles, err := ioutil.ReadDir(fmt.Sprint(relPath, staticViewDir, staticMasterDir))
	if err != nil {
		writer.Write([]byte(fmt.Sprint("unable to read master template dir, err: ", err.Error())))
		return
	}
	var mFiles []string

	mFiles = append(mFiles, fmt.Sprint(relPath, moduleDir, fileNameExt))
	for _, f := range masterFiles {
		mFiles = append(mFiles, fmt.Sprint(relPath, staticViewDir, staticMasterDir, f.Name()))
	}

	t := template.New("result template")
	t, err = template.ParseFiles(mFiles...)
	if err != nil {
		writer.Write([]byte(fmt.Sprint("unable to parse template files, err: ", err.Error())))
		return
	}

	if opt != nil {
		if opt.SiteTitle == "" {
			opt.SiteTitle = opt.PageTitle
		}
		cssFile := fileName
		jsFile := fileName
		if cust := opt.CustomTemplate; cust != nil {
			if cust.ViewJSFile != "" {
				jsFile = cust.ViewJSFile
			}
			if cust.ViewCSSFile != "" {
				cssFile = cust.ViewCSSFile
			}
		}
		if opt.Module != "" {
			jsFile = fmt.Sprint(opt.Module, "/", jsFile)
			cssFile = fmt.Sprint(opt.Module, "/", cssFile)
		}
		opt.MasterPage = MasterPage{
			JSPageFile:  fmt.Sprint(siteDir, "js/", jsFile, ".js"),
			CSSPageFile: fmt.Sprint(siteDir, "css/", cssFile, ".css"),
		}
		js, _ := json.Marshal(opt.JSModel)
		opt.MasterPage.JSObjects = string(js)
	}

	err = t.Execute(writer, *opt)
	if err != nil {
		writer.Write([]byte(fmt.Sprint("unable to render template, err:", err.Error())))
		return
	}
}

//SetupView setups view components: public path dirs
func (u *UIResource) SetupView() {
	u.Router.ServeFiles(fmt.Sprint(siteDir, "*filepath"), http.Dir(fmt.Sprint(assetsRelPath, siteDir)))
}
