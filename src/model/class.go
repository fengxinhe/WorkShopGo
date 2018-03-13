package model

import(
    	"gopkg.in/mgo.v2/bson"
        "../public/database"
        "log"
        "fmt"
)
type Class struct {
    ClassID         bson.ObjectId       `bson:"class_id,omitempty"`
    ClassTitle      string              `bson:"class_title"`
    ClassSummary    string              `bson:"class_summary"`
    FirstTag        string              `bson:"class_tag1"`
    SecondTag       string              `bson:"class_tag2"`
    ClassContent    string              `bson:"class_content"`
    ClassHeat       int                 `bson:"class_heat"`
    ClassSurfaceImg string              `bson:"class_imgs"`
    
    // Content []byte
    // ImgUrl string
    // VideoUrl string
}

type Contest struct {
    ContestTitle string
    ContestSummary string
    //ImgUrl string
}

func GetClasses() (*[]Class){
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

func GetClassesByHeat() (*[]Class) {
    var classes []Class
    session := database.Mongo.Copy()
    defer session.Close()
    c := session.DB(database.ReadConfig().MongoDB.Database).C("class")
    err := c.Find(bson.M{}).Sort("-class_heat").Limit(3).All(&classes)
    if err != nil {
        log.Println("get class error",err)
    }
    //fmt.Println("getclass")
    return &classes
}
