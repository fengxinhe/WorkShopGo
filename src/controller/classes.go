package controller

import (
    "fmt"
    "os"
    "io"
    //"io/ioutil"
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
)

func ClassGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "class"
    v.Data["Title"] = "Xinhe Feng"
    v.Data["Classes"] = model.GetClasses()
    //v.Data["Contests"] = model.GetContest()
    v.RenderTemplate(w)
    return
}

func CreateClassGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "create_class"
    view.Repopulate([]string{"class_surface_img","class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
    v.RenderTemplate(w)
}
// type Class struct {
//     ClassID         bson.ObjectID       `bson:"class_id"`
//     ClassTitle      string              `bson:"class_title"`
//     ClassSummary    string
//     FirstTag        string
//     SecondTag       string
//     // Content []byte
//     // ImgUrl string
//     // VideoUrl string
// }
func CreateClassPost(w http.ResponseWriter, r *http.Request) {

    var class model.Class
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
    //imaging.Save(resizedimg, filepath)
    class.ClassSurfaceImg = filepath
    class.ClassID = bson.NewObjectId()
    class.ClassTitle = r.FormValue("class_title")
    class.ClassSummary = r.FormValue("class_summary")
    class.ClassContent = r.FormValue("class_content")
    class.FirstTag = r.FormValue("first_tag")
    class.SecondTag = r.FormValue("second_tag")
    class.ClassHeat = 10

    if err := model.CreateClass(class); err != nil {
        log.Println(err)
        //respondWithError(w, http.StatusInternalServerError, err.Error())
		return
    } else {
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    CreateClassPost(w, r)
}


func WriteNewPost(w http.ResponseWriter, r *http.Request) {
    
}
