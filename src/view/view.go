package view

import (
    "net/http"
    "html/template"
)

const STATIC_URL string = "/home/firebug/goweb/static/"

var templates *template.Template
var viewInfo View

func init(){
    templates = template.Must(template.ParseGlob("../templates/*")) //Template caching
}

type View struct {
    BaseURI string
    StaticURL string
    Name string
    Data map[string]interface{}
    request *http.Request
}

func Configure(vi View) {
    viewInfo = vi
}

func New(req *http.Request) *View {
    v := &View{}
    v.Data = make(map[string]interface{})

    v.request = req
    v.Data["Static"] = viewInfo.StaticURL
    return v
}


func (v *View)RenderTemplate(w http.ResponseWriter){

    err := templates.ExecuteTemplate(w, v.Name+".html", v.Data)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
}
}
