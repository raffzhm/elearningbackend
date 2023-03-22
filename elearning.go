package elngbackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func RegisterMhs(nm string, em string, nps string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nm
	mahasiswa.Email = em
	mahasiswa.Password = nps

	return InsertOneDoc("dbmhs", "mahasiswa", mahasiswa)
}

func EnrolMatakuliah(mn string, mk string, ml string) (InsertedID interface{}) {
	var matakuliah Matakuliah
	matakuliah.Nama = mn
	matakuliah.Kode = mk
	matakuliah.Lokasi = ml

	return InsertOneDoc("dbmhs", "matakuliah", matakuliah)
}

func GetDataMatakuliahFromKode(kode string, db *mongo.Database, col string) (data Matakuliah) {
	user := db.Collection(col)
	filter := bson.M{"No Code": kode}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getDataCompFromPhoneNumber: %v\n", err)
	}
	return data
}