package wbson

import (
	"dario.cat/mergo"
	"github.com/WingGao/webutils/werror"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/v2/bson"
	"reflect"
)

func NewObjectIdHex(hex string) bson.ObjectID {
	b, _ := bson.ObjectIDFromHex(hex)
	return b
}

type TodoTransformer struct{}

// Transformer is the method that implements the custom transformation logic.
func (t TodoTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	return nil
}
func GetNotNilMap(item interface{}) map[string]interface{} {
	notEmptyMap1 := make(map[string]interface{})
	werror.PanicError(mergo.Map(&notEmptyMap1, item))
	notEmptyMap2 := make(map[string]interface{})
	copier.CopyWithOption(&notEmptyMap2, &notEmptyMap1, copier.Option{IgnoreEmpty: true})
	return notEmptyMap2
}
