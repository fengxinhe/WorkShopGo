package controller

import (
    "fmt"
    "os"
    "io"
    //"io/ioutil"
    "net/http"
    "../view"
//    "../model"
//    "log"
    //"regexp"
    //"errors"
    ///"gopkg.in/mgo.v2/bson"

    //"image"
    //"image/jpeg"
    //"github.com/disintegration/imaging"
)


func CreateProjectGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "new_project"
    view.Repopulate([]string{"class_surface_img","class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
    v.RenderTemplate(w)
}



func CreateProjectPost(w http.ResponseWriter, r *http.Request) {

    //var class model.Project
    r.ParseMultipartForm(32 << 20)
   file, handler, err := r.FormFile("class_surface_img")
   if err != nil {
       fmt.Println(err)
       return
   }
   //imgSize, err := strconv.Atoi(r.FormValue("imgsize"))
   defer file.Close()

   filepath := "/home/firebug/goweb/static/images/class/"+handler.Filename
   fmt.Println("file"+filepath)
   f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
   if err != nil {
       fmt.Println(err)
       return
   }
   defer f.Close()
    io.Copy(f, file)



    CreateProjectPost(w, r)
}
