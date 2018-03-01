package model

import(
    	"gopkg.in/mgo.v2/bson"
        "../public/database"
        "log"
        "fmt"
)
type Class struct {
    ClassID         bson.ObjectId       `bson:"class_id"`
    ClassTitle      string              `bson:"class_title"`
    ClassSummary    string
    FirstTag        string
    SecondTag       string
    // Content []byte
    // ImgUrl string
    // VideoUrl string
}

type Contest struct {
    ContestTitle string
    ContestSummary string
    //ImgUrl string
}

type IndexContent struct {
    Title string
    Classes []Class
    Contests []Contest
    Static string
}

func GetClass() (*[]Class){
//    classes := &[]Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
//                Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},}

    var classes []Class
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("class")
    err := c.Find(bson.M{}).All(&classes)
    if err != nil {
        log.Println("get class error",err)
    }
    fmt.Println("getclass")
    return &classes
}

func CreateClass(class Class) error{
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("class")
    err := c.Insert(&class)
    fmt.Println("create_class")
    return err
}

func GetContest() *[]Contest{
    contests := &[]Contest{Contest{ContestTitle: "t1", ContestSummary: "ccccc"},
                Contest{ContestTitle: "t2", ContestSummary: "dddddd"},}
    return contests
}

func GetIndexContent() *IndexContent {
    content := &IndexContent{
        Title: "DUDU",
        Classes: []Class{Class{ClassTitle: "cl1", ClassSummary: "aaaaa"},
                    Class{ClassTitle: "cl2", ClassSummary: "bbbbb"},},
        Contests: []Contest{Contest{ContestTitle: "t1", ContestSummary: "ccccc"},
                    Contest{ContestTitle: "t2", ContestSummary: "dddddd"},},
    }
    return content
}
