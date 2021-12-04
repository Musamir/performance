package encoding_decoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//╻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╻
//┃       		  UNIT TESTS				┃
//╹━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╹

//TestStandardProtoMarshaling ...
func TestStandardProtoMarshaling(t *testing.T) {
	example := &ExampleProtoData
	expected := protoMarshaled
	actualMarshaled := StandardProtoMarshaling(example)
	assert.Equal(t, expected, actualMarshaled, "not equal!!!")
}

//TestStandardProtoUnmarshalling ...
func TestStandardProtoUnmarshalling(t *testing.T) {
	marshaledExample := protoMarshaled
	expected := protoUnmarshaled
	actualProtoUnmarshaled := &ProtoExample{}
	StandardProtoUnmarshalling(marshaledExample, actualProtoUnmarshaled)
	actualUnmarshaled := toExample(actualProtoUnmarshaled)
	assert.Equal(t, expected, *actualUnmarshaled, "not equal!!!")
}

//TestStandardProtoMarshalingUnmarshalling ...
func TestStandardProtoMarshalingUnmarshalling(t *testing.T) {
	example := &ExampleProtoData
	expected := ExampleData
	actualProto := &ProtoExample{}
	StandardProtoMarshalingUnmarshalling(example, actualProto)
	actual := toExample(actualProto)
	assert.Equal(t, expected, *actual, "not equal!!!")
}

//TestStandardEasyJsonMarshaling ...
func TestStandardEasyJsonMarshaling(t *testing.T) {
	example := &ExampleData
	expected := jsonMarshaled
	actualMarshaled := StandardEasyJsonMarshaling(example)
	assert.Equal(t, expected, actualMarshaled, "not equal!!!")
}

//TestStandardEasyJsonUnmarshalling ...
func TestStandardEasyJsonUnmarshalling(t *testing.T) {
	marshaledExample := jsonMarshaled
	expected := jsonUnmarshaled
	actualUnmarshaled := &Example{}
	StandardEasyJsonUnmarshalling(marshaledExample, actualUnmarshaled)
	assert.Equal(t, expected, *actualUnmarshaled, "not equal!!!")
}

//TestStandardEasyJsonMarshalingUnmarshalling ...
func TestStandardEasyJsonMarshalingUnmarshalling(t *testing.T) {
	example := &ExampleData
	expected := ExampleData
	actual := &Example{}
	StandardEasyJsonMarshalingUnmarshalling(example, actual)
	assert.Equal(t, expected, *actual, "not equal!!!")
}

//TestStandartJsonMarshaling ...
func TestStandartJsonMarshaling(t *testing.T) {
	example := &ExampleData
	expected := jsonMarshaled
	actualMarshaled := StandardJsonMarshaling(example)
	assert.Equal(t, expected, actualMarshaled, "not equal!!!")
}

//TestStandartJsonUnmarshalling ...
func TestStandartJsonUnmarshalling(t *testing.T) {
	marshaledExample := jsonMarshaled
	expected := jsonUnmarshaled
	actualUnmarshaled := &Example{}
	StandardJsonUnmarshalling(marshaledExample, actualUnmarshaled)
	assert.Equal(t, expected, *actualUnmarshaled, "not equal!!!")
}

//TestStandartJsonMarshalingUnmarshalling ...
func TestStandartJsonMarshalingUnmarshalling(t *testing.T) {
	example := &ExampleData
	expected := ExampleData
	actual := &Example{}
	StandardJsonMarshalingUnmarshalling(example, actual)
	assert.Equal(t, expected, *actual, "not equal!!!")
}

//╻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╻
//┃       		  BENCHMARKS				┃
//╹━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╹

//BenchmarkStandardProtoMarshaling ...
func BenchmarkStandardProtoMarshaling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardProtoMarshaling(&ExampleProtoData)
	}
}

//BenchmarkStandardEasyJsonMarshaling ...
func BenchmarkStandardEasyJsonMarshaling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardEasyJsonMarshaling(&ExampleData)
	}
}

func BenchmarkStandardJsonMarshaling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardJsonMarshaling(&ExampleData)
	}
}

//BenchmarkStandardProtoUnmarshalling ...
func BenchmarkStandardProtoUnmarshalling(b *testing.B) {
	var result ProtoExample
	for i := 0; i < b.N; i++ {
		StandardProtoUnmarshalling(protoMarshaled, &result)
	}
}

func BenchmarkStandardEasyJsonUnmarshalling(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardEasyJsonUnmarshalling(jsonMarshaled, &result)
	}
}

func BenchmarkStandardJsonUnmarshalling(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardJsonUnmarshalling(jsonMarshaled, &result)
	}
}

func BenchmarkStandardProtoMarshalingUnmarshalling(b *testing.B) {
	var result ProtoExample
	for i := 0; i < b.N; i++ {
		StandardProtoMarshalingUnmarshalling(&ExampleProtoData, &result)
	}
}

func BenchmarkStandardEasyJsonMarshalingUnmarshalling(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardEasyJsonMarshalingUnmarshalling(&ExampleData, &result)
	}
}

func BenchmarkStandardJsonMarshalingUnmarshalling(b *testing.B) {
	var result Example
	for i := 0; i < b.N; i++ {
		StandardJsonMarshalingUnmarshalling(&ExampleData, &result)
	}
}
