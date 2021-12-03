package encoding_decoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//ExampleData ...
var ExampleData = Example{
	A: 10,
	B: "20",
	C: []int32{1, 2, 3, 4, 5},
	D: []string{"12", "23", "34", "45", "56"},
}

//TestStandardEncodingDecoding ...
func TestStandardEncodingDecoding(t *testing.T) {
	assertions := assert.New(t)
	expect := ExampleData
	{ //json testing
		resultJson := StandardJsonEncoding(&ExampleData)
		var actualJsonResult Example
		StandardJsonDecoding(resultJson, &actualJsonResult)
		assertions.Equal(expect, actualJsonResult, "json result not equal!!!")

		result := &Example{}
		StandardJsonEncodingDecoding(&ExampleData, result)
		assertions.Equal(expect, *result, "json not equal StandardJsonEncodingDecoding")
	}
	{ //gob testing
		resultGob := StandardGobEncoding(&ExampleData)
		var actualGobResult Example
		StandardGobDecoding(resultGob, &actualGobResult)
		assertions.Equal(expect, actualGobResult, "gob result not equal!!!")

		result := &Example{}
		StandardGobEncodingDecoding(&ExampleData, result)
		assertions.Equal(expect, *result, "gob not equal StandardGobEncodingDecoding")
	}
	{ //proto testing
		ProtoData := &ProtoExample{A: ExampleData.A, B: ExampleData.B, C: ExampleData.C, D: ExampleData.D}
		marshaled := StandardProtoMarshaling(ProtoData)

		UnmarshalResult := &ProtoExample{}
		StandardProtoUnmarshalling(marshaled, UnmarshalResult)
		actualResult := Example{A: UnmarshalResult.A, B: UnmarshalResult.B, C: UnmarshalResult.C, D: UnmarshalResult.D}
		assertions.Equal(expect, actualResult, "proto not equal!!!")

		protoResult := &ProtoExample{}

		StandardProtoMarshalingUnmarshalling(ProtoData, protoResult)
		result := Example{A: protoResult.A, B: protoResult.B, C: protoResult.C, D: protoResult.D}
		assertions.Equal(expect, result, "proto not equal StandardProtoMarshalingUnmarshalling")
	}
}

//BenchmarkStandardJsonEncoding ...
func BenchmarkStandardJsonEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardJsonEncoding(&ExampleData)
	}
}

//BenchmarkStandardGobEncoding ...
func BenchmarkStandardGobEncoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardGobEncoding(&ExampleData)
	}
}

func BenchmarkStandardProtoMarshaling(b *testing.B) {
	ProtoData := &ProtoExample{
		A: ExampleData.A,
		B: ExampleData.B,
		C: ExampleData.C,
		D: ExampleData.D,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardProtoMarshaling(ProtoData)
	}
}

//BenchmarkStandardJsonDecoding ...
func BenchmarkStandardJsonDecoding(b *testing.B) {
	var ex1 Example
	encodedBytes := StandardJsonEncoding(&ExampleData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardJsonDecoding(encodedBytes, &ex1)
	}
}

//BenchmarkStandardGobDecoding ...
func BenchmarkStandardGobDecoding(b *testing.B) {
	var ex Example
	encodedBytes := StandardGobEncoding(&ExampleData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardGobDecoding(encodedBytes, &ex)
	}
}

//BenchmarkStandardProtoUnmarshalling ...
func BenchmarkStandardProtoUnmarshalling(b *testing.B) {
	var ex ProtoExample
	ProtoData := &ProtoExample{
		A: ExampleData.A,
		B: ExampleData.B,
		C: ExampleData.C,
		D: ExampleData.D,
	}
	marshaled := StandardProtoMarshaling(ProtoData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardProtoUnmarshalling(marshaled, &ex)
	}
}

//BenchmarkStandardJsonEncodingDecoding ...
func BenchmarkStandardJsonEncodingDecoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardJsonEncodingDecoding(&ExampleData, &Example{})
	}
}

//BenchmarkStandardGobEncodingDecoding ...
func BenchmarkStandardGobEncodingDecoding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardGobEncodingDecoding(&ExampleData, &Example{})
	}
}

func BenchmarkStandardProtoMarshalingUnmarshalling(b *testing.B) {
	ProtoData := &ProtoExample{
		A: ExampleData.A,
		B: ExampleData.B,
		C: ExampleData.C,
		D: ExampleData.D,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StandardProtoMarshalingUnmarshalling(ProtoData, &ProtoExample{})
	}
}
