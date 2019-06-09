package db

import (

	"time"
	"log"
	// "reflect"

	mgo "gopkg.in/mgo.v2"
	
)

type Ab struct{
	A string `json:"a"`
}

func ConnPool () *mgo.Session {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  60 * time.Second,
		Database: "test",
		Username: "",
		Password: "",
	}
	
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	// defer mongoSession.Close()
	return mongoSession
	// c := session.DB(database).C(collection)
	// err := c.Find(query).One(&result)
}

var Pool = ConnPool()

func DBops(db string, collection string) *mgo.Collection{

	var Connection = Pool.DB(db).C(collection)
	// log.Println("type>>>", reflect.TypeOf(Conn))
	//----------------Query Demo----------------------------
	// type x Ab
	// err := Connection.Insert(x{"Inserted"})
	// if err != nil{
	// 	log.Println("error in insertion", err)
	// }else{
	// 	log.Println("Inserted")
	// }
	return Connection
}






