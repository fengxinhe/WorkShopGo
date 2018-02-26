package main

import (
    "fmt"
    "os"
    "encoding/json"
    "./controller/server"
    "./controller/static"
    "./view"
    "./controller/route"
    "./public/database"
    "./public/jsonconfig"
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
	Database  database.Info   `json:"Database"`
//	Email     email.SMTPInfo  `json:"Email"`
//	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
//	Session   session.Session `json:"Session"`
//	Template  view.Template
	View      view.View       `json:"View"`
    Static    static.StaticInfo    `json:"Static"`
}

var config = &configuration{}

func main() {
    fmt.Println("start")

    jsonconfig.LoadConfig("config" + string(os.PathSeparator)+"config.json", config)
    view.Configure(config.View)
    static.Configure(config.Static)
    database.Connect(config.Database)

	//view.LoadTemplates(config.Template.Root, config.Template.Children)

	// Start the listener
    server.Run(route.LoadHTTP(), route.LoadHTTPS(), config.Server)

    //http.ListenAndServe(":8000", nil)
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
