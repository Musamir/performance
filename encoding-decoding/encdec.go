//go:generate protoc --go_out=. --go_opt=paths=source_relative Example.proto
package encoding_decoding

import (
	"bytes"

	"encoding/gob"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

//Example ...
type Example struct {
	A int32
	B string
	C []int32
	D []string
}

//StandardJsonEncoding ...
func StandardJsonEncoding(toEnc *Example) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(*toEnc); err != nil {
		failed("json encoding", err)
	}
	return buf.Bytes()
}

//StandardJsonDecoding ...
func StandardJsonDecoding(encoded []byte, toDec *Example) {
	buf := bytes.NewBuffer(encoded)
	dec := json.NewDecoder(buf)
	if err := dec.Decode(toDec); err != nil {
		failed("json decoding", err)
	}
}

//StandardJsonEncodingDecoding ...
func StandardJsonEncodingDecoding(data, result *Example) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(*data); err != nil {
		failed("json encoding", err)
	}
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(result); err != nil {
		failed("json decoding", err)
	}
}

//StandardGobEncoding ...
func StandardGobEncoding(toEnc *Example) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*toEnc); err != nil {
		failed("gob encoding", err)
	}
	return buf.Bytes()
}

//StandardGobDecoding ...
func StandardGobDecoding(encoded []byte, toDec *Example) {
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(toDec); err != nil {
		failed("gob decoding", err)
	}
}

//StandardGobEncodingDecoding ...
func StandardGobEncodingDecoding(data, result *Example) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*data); err != nil {
		failed("gob encoding", err)
	}
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(result); err != nil {
		failed("gob decoding", err)
	}
}

//StandardProtoMarshaling ...
func StandardProtoMarshaling(data *ProtoExample) []byte {
	marshaled, err := proto.Marshal(data)
	if err != nil {
		failed("proto marshaling", err)
	}
	return marshaled
}

//StandardProtoUnmarshalling ...
func StandardProtoUnmarshalling(marshaled []byte, toUnmarshal *ProtoExample) {
	if err := proto.Unmarshal(marshaled, toUnmarshal); err != nil {
		failed("proto unmarshalling", err)
	}
}

//StandardProtoMarshalingUnmarshalling ...
func StandardProtoMarshalingUnmarshalling(data, result *ProtoExample) {
	marshaled, err := proto.Marshal(data)
	if err != nil {
		failed("proto marshaling", err)
	}

	if err := proto.Unmarshal(marshaled, result); err != nil {
		failed("proto unmarshalling", err)
	}
}

func failed(str string, err error) {
	panic(str + " [an error occurred] " + err.Error())
}
