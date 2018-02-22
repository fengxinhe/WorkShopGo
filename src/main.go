package main

import (
    //"fmt"
    "./controller/server"
    "./view"
    "./controller/route"
    //"net/http"
    //"regexp"
    //"errors"
)


// func loadPage(title string) (*Page, error) {
//     filename := title + ".txt"
//     body, err := ioutil.ReadFile(filename)
//     if err != nil {
//         return nil, err
//     }
//     return &Page{Title: title, Body: body}, nil
// }

type configuration struct {
//	Database  database.Info   `json:"Database"`
//	Email     email.SMTPInfo  `json:"Email"`
//	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server
//	Session   session.Session `json:"Session"`
//	Template  view.Template
	View      view.View
}

var config = &configuration{}
const STATIC_URL string = "/home/firebug/goweb/static/"

func main() {


    //view.Configure(config.View)
	//view.LoadTemplates(config.Template.Root, config.Template.Children)

    config.Server= server.Server{
        HostName: "localhost",
        UseHttp: true,
        UseHttps: false,
        HttpPort: "8000",
        HttpsPort: "8080",
    }
	// Start the listener
    server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)

    //http.ListenAndServe(":8000", nil)
}
