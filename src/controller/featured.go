package controller

import (
    "fmt"
    //"os"
    //"io"
    "io/ioutil"
    "net/http"
    "../view"
    "../model"
    "log"
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

func ShowProjectGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "project"
    var project = model.GetProject()
    var content = ""
    for i:=0; i<project.ProjectStepNum;i++{
        title := project.ProjectStepTitle[i]
        path:=project.ProjectSteps[title]
        content += readfile(path)
    }

    v.Data["Content"]=content
    v.RenderTemplate(w)
    return
}

func readfile(path string) string{
    data,_ := ioutil.ReadFile(path)
    return string(data[:])
}

func CreateProjectGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "new_project"
    //view.Repopulate([]string{"class_surface_img","class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
    v.RenderTemplate(w)
}

// type Project struct{
//     ProjectID       bson.ObjectId       `bson:"project_id,omitempty"`
//     ProjectTitle    string              `bson:"project_title"`
//     ProjectSteps    map[string]string   `bson:"project_steps"`
//     ProjectCategory string              `bson:"project_tag1"`
//     ProjectChannel  string              `bson:"project_tag2"`
//     ProjectHeat     int                 `bson:"project_heat"`
//      ProjectStepNum int
//      ProjectStepTitle []string
// }
func CreateProjectPost(w http.ResponseWriter, r *http.Request) {

    //var class model.Project
    r.ParseMultipartForm(32 << 20)
   //file, handler, err := r.FormFile("editdata")
   var project model.Project
   project.ProjectID = bson.NewObjectId()
   project.ProjectTitle = r.FormValue("project_title")
   stepcount := r.FormValue("stepcount")
   project.ProjectHeat = 10
   project.ProjectSteps=make(map[string]string)

   cnt, _ := strconv.Atoi(stepcount)
   project.ProjectStepNum = cnt

   for i := 1; i <= cnt; i++ {
       var temp=r.FormValue("summernotecode"+strconv.Itoa(i))
       var steptitle = r.FormValue("step_title"+strconv.Itoa(i))
       project.ProjectStepTitle=append(project.ProjectStepTitle, steptitle)
       var fn=project.ProjectID.String()+"_step_"+strconv.Itoa(i)
       project.ProjectSteps[steptitle]=saveFile(fn, temp)
   }

   project.ProjectCategory = "test"
   project.ProjectChannel = "tttest"
   if err := model.CreateProject(project); err != nil {
       log.Println(err)
       //respondWithError(w, http.StatusInternalServerError, err.Error())
       return
   } else {
       http.Redirect(w, r, "/", http.StatusFound)
       return
   }

   //dd :=r.FormValue("summernotecode1")
   // if err != nil {
   //     fmt.Println(err)
   //     return
   // }
   // //defer file.Close()

   //fmt.Println(handler.Filename)

    //var c=r.FormValue("content")
    fmt.Println("ok")

    //CreateProjectPost(w, r)
}

func saveFile(fn string, content string) string {
    data := []byte(content)
    path := "/home/firebug/goweb/static/projects/files/"+fn
    ioutil.WriteFile(path, data,0600)
    return path
}
