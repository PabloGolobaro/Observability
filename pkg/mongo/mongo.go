package mongo

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const database = "valutes"

var instance *DB

//InitMongo init Data Base
func InitMongo(url string) {
	db, err := NewConnection(url)
	if err != nil {
		log.Panicf("DB init false with err: %s", err.Error())
		panic(err)
	}
	instance = db
}

//GetDB return copy mgo instance
func GetDB() *DB {
	return instance
}

//DB is connection struct to mongoDB
//use mgo
type DB struct {
	client *mongo.Client
}

//NewConnection return new DB connection
//connection with url
//in url must be pass and user if needed
func NewConnection(url string) (*DB, error) {
	var db = DB{}
	var err error
	cwt, _ := context.WithTimeout(context.Background(), time.Second*10)
	clientOptions := options.Client().ApplyURI(url)
	db.client, err = mongo.Connect(cwt, clientOptions)
	if err != nil {
		return nil, err
	}
	if err := db.client.Ping(cwt, readpref.Primary()); err != nil {
		return nil, err
	}
	return &db, nil
}

//IsConnected check connection to mongo db Server
func (db *DB) IsConnected() bool {
	return db.client != nil
}

//Insert insert document to collection
//if collection is not created this function create collection
func (db *DB) Insert(coll string, v ...interface{}) error {
	_, err := db.client.Database(database).Collection(coll).InsertOne(context.Background(), v)
	return err
}

//Find find document in collection
func (db *DB) Find(coll string, query map[string]interface{}, v interface{}) error {
	bsonQuery := bson.M{}
	for k, qv := range query {
		bsonQuery[k] = qv
	}
	cursor, err := db.client.Database(database).Collection(coll).Find(context.Background(), bsonQuery)
	if err != nil {
		return err
	}
	err = cursor.All(context.Background(), v)
	return err
}

//FindByID find document by ID
func (db *DB) FindByID(coll string, id string, v interface{}) error {
	err := db.client.Database(database).Collection(coll).FindOne(context.Background(), bson.M{"id": id}).Decode(v)
	return err
}

//FindWithQuery you can call this function with query
//you can must use mgo.bson format
func (db *DB) FindWithQuery(coll string, query interface{}, v interface{}) error {

	err := db.client.Database(database).Collection(coll).FindOne(context.Background(), query).Decode(v)
	return err
}

//FindWithQueryAll you can find all document in collection with this function
//you can call this function with mgo.bson query
func (db *DB) FindWithQueryAll(coll string, query interface{}, v interface{}) error {

	cursor, err := db.client.Database(database).Collection(coll).Find(context.Background(), query)
	if err != nil {
		return err
	}
	err = cursor.All(context.Background(), v)
	return err
}

/*
//RemoveWithIDs delete all document in collection by ids
func (db *DB) RemoveWithIDs(coll string, ids interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var sess = db.client.Copy()
	defer sess.Close()

	_, err := sess.DB("").C(coll).RemoveAll(bson.M{"_id": bson.M{"$in": ids}})

	return err
}

//Update document by query
//warning you can update all document with this query
func (db *DB) Update(coll string, query interface{}, set interface{}) error {
	if !db.IsConnected() {
		return &IsNotConnected
	}
	var err error
	var sess = db.client.Copy()
	defer sess.Close()

	_, err = sess.DB("").C(coll).UpdateAll(query, set)

	return err
}*/

//FindAll find all document in collection
func (db *DB) FindAll(coll string, v interface{}) error {
	cursor, err := db.client.Database(database).Collection(coll).Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	err = cursor.All(context.Background(), v)
	return err
}
