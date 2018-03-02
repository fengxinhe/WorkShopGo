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
    v.Data["Classes"] = model.GetClass()
    v.Data["Contests"] = model.GetContest()
    v.RenderTemplate(w)
    return
}
