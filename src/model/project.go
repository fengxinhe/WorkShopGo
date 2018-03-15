package model

import(
    	"gopkg.in/mgo.v2/bson"
        "../public/database"
        "log"
    //    "fmt"
)

type Project struct{
    ProjectID       bson.ObjectId       `bson:"project_id,omitempty"`
    ProjectTitle    string              `bson:"project_title"`
    ProjectSummary  string              `bson:"project_summary"`
    ProjectContent  string              `bson:"project_content"`
    FirstTag        string              `bson:"project_tag1"`
    SecondTag       string              `bson:"project_tag2"`
    ProjectHeat     int                 `bson:"project_heat"`
}

type Step struct {
    ProjectID       bson.ObjectId
    Imgs            string
    Videos          string
    Text            string

}


func GetProjectByTag(tag1 string, tag2 string) *[]Project{
    var projects []Project

    return &projects
}

func GetProjects() *[]Project{
    var projects []Project
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("project")
    err := c.Find(bson.M{}).All(&projects)
    if err != nil {
        log.Println("get project error",err)
    }
    return &projects
}

// func GetClassByHeat() *[]Class{
//     //class := from db
//
// }
