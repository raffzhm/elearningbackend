package rafi1214005

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email		 string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Matakuliah struct {
	Nama       string             `bson:"namamatakuliah,omitempty" json:"namamatakuliah,omitempty"`
	Kode       string             `bson:"kode,omitempty" json:"kode,omitempty"`
	Lokasi     string             `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	Enrol      string             `bson:"enrol,omitempty" json:"enrol,omitempty"`
}

