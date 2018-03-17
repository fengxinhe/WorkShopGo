package model

import(
    	"gopkg.in/mgo.v2/bson"
        "../public/database"
        "log"
       "fmt"
)

type Project struct{
    ProjectID           bson.ObjectId       `bson:"project_id"`
    ProjectTitle        string              `bson:"project_title"`
    ProjectSteps        map[string]string   `bson:"project_steps"`
    ProjectCategory     string              `bson:"project_category"`
    ProjectChannel      string              `bson:"project_channel"`
    ProjectHeat         int                 `bson:"project_heat"`
    ProjectStepNum      int                 `bson:"project_stepnum"`
    ProjectStepTitle    []string            `bson:"project_step_title"`
}

type Step struct {
    ProjectID       bson.ObjectId
    StepTitle       string
    Imgs            []string
    Videos          []string
    Text            string

}

func CreateProject(project Project) error{
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("project")
    err := c.Insert(&project)
    fmt.Println("create_projrct")
    return err
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

func GetProject() *Project{
    var project Project
    //var projects []Project
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("project")
    err := c.Find(bson.M{"ProjectStepNum": 3}).One(&project)
    if err != nil {
        log.Println("get project error",err)
    }
    //project=projects[0]
    //fmt.Println(project)
    return &project
}
