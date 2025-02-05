package wbson

import "go.mongodb.org/mongo-driver/v2/bson"

func NewObjectIdHex(hex string) bson.ObjectID {
	b, _ := bson.ObjectIDFromHex(hex)
	return b
}
