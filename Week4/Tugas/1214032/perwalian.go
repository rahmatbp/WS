package perwalian

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"

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

func InsertWaktu(jam string, hari string, tanggal string) (InsertedID interface{}) {
	var waktu Waktu
	waktu.Jam = jam
	waktu.Hari = hari
	waktu.Tanggal = tanggal
	return InsertOneDoc("tugas_db", "perwalian", waktu)
}

func GetMahasiswaFromJam(jam string) (datang Waktu) {
	mahasiswa := MongoConnect("tugas_db").Collection("perwalian")
	filter := bson.M{"jam": jam}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&datang)
	if err != nil {
		fmt.Printf("getMahasiswaFromJam: %v\n", err)
	}
	return datang
}
