package encoding_decoding

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

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
