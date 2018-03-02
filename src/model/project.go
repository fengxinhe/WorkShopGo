package model

import(
    	"gopkg.in/mgo.v2/bson"
    //    "../public/database"
    //    "log"
    //    "fmt"
)

type Project struct{
    ProjectID       bson.ObjectId       `bson:"project_id"`
    ProjectTitle    string              `bson:"project_title"`
    ProjectSummary  string              `bson:"project_summary"`
    ProjectContent  string              `bson:"project_content"`
    FirstTag        string              `bson:"project_tag1"`
    SecondTag       string              `bson:"project_tag2"`
    ProjectHeat     int                 `bson:"project_heat"`
}


func GetProjectByTag(tag1 string, tag2 string) *[]Class{

    classes := &[]Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
                Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},}
    return classes
}

func GetProjects() *[]Contest{
    contests := &[]Contest{Contest{ContestTitle: "t1", ContestSummary: "ccccc"},
                Contest{ContestTitle: "t2", ContestSummary: "dddddd"},}
    return contests
}

// func GetClassByHeat() *[]Class{
//     //class := from db
//
// }
