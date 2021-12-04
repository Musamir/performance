package encoding_decoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//╻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╻
//┃       		  UNIT TESTS				┃
//╹━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╹

//TestStandardJsonEncoding ...
func TestStandardJsonEncoding(t *testing.T) {
	example := &ExampleData
	expected := jsonEncoded
	actualEncoded := StandardJsonEncoding(example)
	assert.Equal(t, expected, actualEncoded, "not equal!!!")
}

//TestStandardJsonDecoding ...
func TestStandardJsonDecoding(t *testing.T) {
	encodedExample := jsonEncoded
	expected := jsonDecoded
	actualDecoded := &Example{}
	StandardJsonDecoding(encodedExample, actualDecoded)
	assert.Equal(t, expected, *actualDecoded, "not equal!!!")
}

//TestStandardJsonEncodingDecoding ...
func TestStandardJsonEncodingDecoding(t *testing.T) {
	example := &ExampleData
	expected := ExampleData
	actual := &Example{}
	StandardJsonEncodingDecoding(example, actual)
	assert.Equal(t, expected, *actual, "not equal!!!")
}

//TestStandardGobEncoding ...
func TestStandardGobEncoding(t *testing.T) {
	example := &ExampleData
	expected := gobEncoded
	actualEncoded := StandardGobEncoding(example)
	assert.Equal(t, expected, actualEncoded, "not equal!!!")
}

//TestStandardGobDecoding ...
func TestStandardGobDecoding(t *testing.T) {
	encodedExample := gobEncoded
	expected := gobDecoded
	actualDecoded := &Example{}
	StandardGobDecoding(encodedExample, actualDecoded)
	assert.Equal(t, expected, *actualDecoded, "not equal!!!")
}

//TestStandardGobEncodingDecoding ...
func TestStandardGobEncodingDecoding(t *testing.T) {
	example := &ExampleData
	expected := ExampleData
	actual := &Example{}
	StandardGobEncodingDecoding(example, actual)
	assert.Equal(t, expected, *actual, "not equal")
}

//╻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╻
//┃       		  BENCHMARKS				┃
//╹━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╹

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

//BenchmarkStandardJsonDecoding ...
func BenchmarkStandardJsonDecoding(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardJsonDecoding(jsonEncoded, &result)
	}
}

//BenchmarkStandardGobDecoding ...
func BenchmarkStandardGobDecoding(b *testing.B) {
	var ex Example
	for i := 0; i < b.N; i++ {
		StandardGobDecoding(gobEncoded, &ex)
	}
}

//BenchmarkStandardJsonEncodingDecoding ...
func BenchmarkStandardJsonEncodingDecoding(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardJsonEncodingDecoding(&ExampleData, &result)
	}
}

//BenchmarkStandardGobEncodingDecoding ...
func BenchmarkStandardGobEncodingDecoding(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardGobEncodingDecoding(&ExampleData, &result)
	}
}
