package database

import(
    "time"
    "log"
    "gopkg.in/mgo.v2"
    "fmt"
)
type Type string
const (

	// TypeMongoDB is MongoDB
	TypeMongoDB Type = "MongoDB"
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
)

type MySQLInfo struct {
    Username string
    Password string
    Name string
    Hostname string
    Port int
    Parameter string
}

type Info struct {
    Type Type
    MongoDB MongoInfo
    MySQL MySQLInfo
}
var Mongo *mgo.Session
var database Info


type MongoInfo struct {
    URL string
    Database string
}

func Connect(d Info) {
    var err error
    database = d
    if Mongo, err = mgo.DialWithTimeout(d.MongoDB.URL, 5*time.Second); err!=nil{
        log.Println("MongoDB error", err)
        return
    }
    Mongo.SetSocketTimeout(1 * time.Second)

    if err = Mongo.Ping(); err != nil {
        log.Println("Database erro", err)
    }
    CheckConnection()
}

func CheckConnection() bool{
    if Mongo == nil{
        Connect(database)
    }
    if Mongo != nil {
        fmt.Println("db ok")
        return true
    }
    return false
}

func ReadConfig() Info {
	return database
}
