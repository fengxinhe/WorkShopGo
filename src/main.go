package main

import (
    //"fmt"
    "io/ioutil"
    "net/http"
    "html/template"
    "io"
    "time"

    "./controller"
    //"regexp"
    //"errors"
)
const STATIC_URL string = "/home/firebug/goweb/static/"
const STATIC_ROOT string = "/home/firebug/goweb/static/"

type Page struct {
    Title string
    Body  []byte
    Static string
}

var templates *template.Template

func init(){
    templates = template.Must(template.ParseGlob("../templates/*")) //Template chching
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
    static_file := req.URL.Path[len(STATIC_URL):]
    if len(static_file) != 0 {
        f, err := http.Dir(STATIC_ROOT).Open(static_file)
        if err == nil {
            content := io.ReadSeeker(f)
            http.ServeContent(w, req, static_file, time.Now(), content)
            return
        }
    }
    http.NotFound(w, req)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
    p.Static = STATIC_URL
    err := templates.ExecuteTemplate(w, tmpl+".html",p)
    if err != nil{
        http.Error(w,err.Error(), http.StatusInternalServerError)
    }
}

// func indexHandler(w http.ResponseWriter, r *http.Request){
//     //content := &Page{Title: "main", Body: []byte("kkk")}
//     content := model.GetIndexContent()
//     //renderTemplate(w, "index",content)
//     content.Static = STATIC_URL
//     templates.ExecuteTemplate(w, "index"+".html",content)
// }


func main() {

    //myMux := http.NewServeMux()
    http.HandleFunc("/", controller.IndexGet)
    http.HandleFunc(STATIC_URL, StaticHandler)

    http.ListenAndServe(":8000", nil)
}
