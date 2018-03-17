package controller

import (
    "fmt"
    //"os"
    //s"io"
    //"io/ioutil"
    "net/http"
    "../view"
    "../model"
//    "log"
    //"regexp"
    //"errors"
    "gopkg.in/mgo.v2/bson"

    //"image"
    //"image/jpeg"
    //"github.com/disintegration/imaging"
    //"github.com/gorilla/schema"
    //"strings"

     "strconv"
)


func CreateProjectGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "new_project"
    view.Repopulate([]string{"class_surface_img","class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
    v.RenderTemplate(w)
}

// type Project struct{
//     ProjectID       bson.ObjectId       `bson:"project_id,omitempty"`
//     ProjectTitle    string              `bson:"project_title"`
//     ProjectSteps    map[string]string   `bson:"project_steps"`
//     ProjectCategory string              `bson:"project_tag1"`
//     ProjectChannel  string              `bson:"project_tag2"`
//     ProjectHeat     int                 `bson:"project_heat"`
// }
func CreateProjectPost(w http.ResponseWriter, r *http.Request) {

    //var class model.Project
    r.ParseMultipartForm(32 << 20)
   //file, handler, err := r.FormFile("editdata")
   var project model.Project
   project.ProjectSteps=make(map[string]string)

   stepcount := r.FormValue("stepcount")
   cnt, _ := strconv.Atoi(stepcount)
   fmt.Println(r.FormValue("summernotecode2"))

   for i := 1; i <= cnt; i++ {
       var temp=r.FormValue("summernotecode"+strconv.Itoa(i))
       //fmt.Println(temp)
       project.ProjectSteps["step"+strconv.Itoa(i)]=temp
   }

   project.ProjectID = bson.NewObjectId()
   //dd :=r.FormValue("summernotecode1")
   // if err != nil {
   //     fmt.Println(err)
   //     return
   // }
   // //defer file.Close()

   //fmt.Println(handler.Filename)

    //var c=r.FormValue("content")
    fmt.Println(project.ProjectSteps)



    //CreateProjectPost(w, r)
}
