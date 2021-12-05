// Package encoding_decoding
//Proto generate command
//go:generate protoc --go_out=. --go_opt=paths=source_relative Example.proto
//Easyjson generate command
//go:generate easyjson -no_std_marshalers Example.go
package encoding_decoding

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/mailru/easyjson"
)

//StandardProtoMarshaling ...
func StandardProtoMarshaling(data *ProtoExample) []byte {
	marshaled, err := proto.Marshal(data)
	if err != nil {
		failed("proto marshaling")
	}
	return marshaled
}

//StandardProtoUnmarshalling ...
func StandardProtoUnmarshalling(marshaled []byte, toUnmarshal *ProtoExample) {
	if err := proto.Unmarshal(marshaled, toUnmarshal); err != nil {
		failed("proto unmarshalling")
	}
}

//StandardProtoMarshalingUnmarshalling ...
func StandardProtoMarshalingUnmarshalling(data, result *ProtoExample) {
	marshaled, err := proto.Marshal(data)
	if err != nil {
		failed("proto marshaling")
	}

	if err := proto.Unmarshal(marshaled, result); err != nil {
		failed("proto unmarshalling")
	}
}

//StandardEasyJsonMarshaling ...
func StandardEasyJsonMarshaling(toMarshal *Example) []byte {
	marshaled, err := easyjson.Marshal(toMarshal)
	if err != nil {
		failed("easyjson marshaling")
	}
	return marshaled
}

//StandardEasyJsonUnmarshalling ...
func StandardEasyJsonUnmarshalling(marshaled []byte, toUnmarshal *Example) {
	if err := easyjson.Unmarshal(marshaled, toUnmarshal); err != nil {
		failed("easyjson unmarshalling")
	}
}

//StandardEasyJsonMarshalingUnmarshalling ...
func StandardEasyJsonMarshalingUnmarshalling(data, result *Example) {
	marshaled, err := easyjson.Marshal(data)
	if err != nil {
		failed("easyjson marshaling")
	}
	if err := easyjson.Unmarshal(marshaled, result); err != nil {
		failed("easyjson unmarshalling")
	}
}

//StandardJsonMarshaling ...
func StandardJsonMarshaling(toMarshal *Example) []byte {
	marshaled, err := json.Marshal(toMarshal)
	if err != nil {
		failed("json marshaling")
	}
	return marshaled
}

//StandardJsonUnmarshalling ...
func StandardJsonUnmarshalling(marshaled []byte, toUnmarshal *Example) {
	if err := json.Unmarshal(marshaled, toUnmarshal); err != nil {
		failed("json unmarshalling")
	}
}

//StandardJsonMarshalingUnmarshalling ...
func StandardJsonMarshalingUnmarshalling(data, result *Example) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		failed("json marshaling")
	}
	if err := json.Unmarshal(marshaled, result); err != nil {
		failed("json unmarshalling")
	}
}
