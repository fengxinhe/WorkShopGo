package main

import (
    //"fmt"
    "io/ioutil"
    "net/http"
    "html/template"
    //"regexp"
    //"errors"
)

type Page struct {
    Title string
    Body  []byte
}

var templates *template.Template

func init(){
    templates = template.Must(template.ParseGlob("../../templates/*")) //Template chching
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
    err := templates.ExecuteTemplate(w, tmpl+".html",p)
    if err != nil{
        http.Error(w,err.Error(), http.StatusInternalServerError)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request){
    content := &Page{Title: "main", Body: []byte("kkk")}
    renderTemplate(w, "index",content)
}


func main() {

    myMux := http.NewServeMux()
    myMux.HandleFunc("/", indexHandler)


    http.ListenAndServe(":8000", myMux)
}
