package controller

import (
    //"fmt"
    //"io/ioutil"
    "net/http"
    "../view"
    "../model"
    //"regexp"
    //"errors"
)

func IndexGet(w http.ResponseWriter, r *http.Request) {
    v := view.New(r)
    v.Name = "index"
    //v.Data["Title"] = "Xinhe Feng"
    v.Data["Classes"] = model.GetClassesByHeat()
    //v.Data["Projects"] = model.GetProjects()
    v.RenderTemplate(w)
    return
}
