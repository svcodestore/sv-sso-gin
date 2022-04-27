package utils

import (
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func ToRpcStruct(data interface{}) *structpb.Struct {
	b, e := json.Marshal(data)
	if e != nil {
		return nil
	}
	var m map[string]interface{}
	e = json.Unmarshal(b, &m)
	if e != nil {
		return nil
	}
	s, err := structpb.NewStruct(m)
	if err != nil {
		return nil
	}
	return s
}
