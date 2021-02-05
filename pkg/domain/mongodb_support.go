package domain

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//GetMongoDBDatabase provide a MongoDB session
func GetMongoDBDatabase(db string) (*mgo.Database, error) {
	mongoDBDialInfo := getConnectionInfo()
	log.Println("Mongo dial info: ", mongoDBDialInfo)
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Printf("Error getting Mongo session: %e\n", err)
		return nil, err
	}
	return session.DB(db), nil
}

func getConnectionInfo() *mgo.DialInfo {
	host := os.Getenv("MONGO_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	mongoURI := "mongodb://" + host + ":27017"
	mongoDB := "test"
	dialInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {
		return nil
	}

	dialInfo.Database = mongoDB
	dialInfo.Timeout = time.Second * 15

	return dialInfo
}
