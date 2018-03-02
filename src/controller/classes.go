package controller

import (
    //"fmt"
    //"io/ioutil"
    "net/http"
    "../view"
    "../model"
    "log"
    //"regexp"
    //"errors"
    "gopkg.in/mgo.v2/bson"

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
    view.Repopulate([]string{"class_title", "class_summary", "class_content","first_tag","secondtag"}, r.Form, v.Data)
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
