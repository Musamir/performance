package encoding_decoding

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

//StandardJsonEncoding ...
func StandardJsonEncoding(toEnc *Example) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(*toEnc); err != nil {
		failed("json encoding")
	}
	return buf.Bytes()
}

//StandardJsonDecoding ...
func StandardJsonDecoding(encoded []byte, toDec *Example) {
	buf := bytes.NewBuffer(encoded)
	dec := json.NewDecoder(buf)
	if err := dec.Decode(toDec); err != nil {
		failed("json decoding")
	}
}

//StandardJsonEncodingDecoding ...
func StandardJsonEncodingDecoding(data, result *Example) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(*data); err != nil {
		failed("json encoding")
	}
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(result); err != nil {
		failed("json decoding")
	}
}

//StandardGobEncoding ...
func StandardGobEncoding(toEnc *Example) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*toEnc); err != nil {
		failed("gob encoding")
	}
	return buf.Bytes()
}

//StandardGobDecoding ...
func StandardGobDecoding(encoded []byte, toDec *Example) {
	buf := bytes.NewBuffer(encoded)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(toDec); err != nil {
		failed("gob decoding")
	}
}

//StandardGobEncodingDecoding ...
func StandardGobEncodingDecoding(data, result *Example) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*data); err != nil {
		failed("gob encoding")
	}
	dec := gob.NewDecoder(&buf)
	if err := dec.Decode(result); err != nil {
		failed("gob decoding")
	}
}

func failed(str string) {
	fmt.Println(" [an error occurred] " + str)
}
