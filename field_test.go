package log

import (
	"errors"
	"github.com/MuratSs/assert"
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	Assert := assert.With(t)

	field := Bool("name", true)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Bool)
	Assert.That(field.BoolValue).IsEqualTo(true)
}

func TestInt(t *testing.T) {
	Assert := assert.With(t)

	field := Int("name", -9876)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Int)
	Assert.That(field.IntValue).IsEqualTo(-9876)
}

func TestInt8(t *testing.T) {
	Assert := assert.With(t)

	field := Int8("name", -127)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Int8)
	Assert.That(field.IntValue).IsEqualTo(-127)
}

func TestInt16(t *testing.T) {
	Assert := assert.With(t)

	field := Int16("name", -32767)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Int16)
	Assert.That(field.IntValue).IsEqualTo(-32767)
}

func TestInt32(t *testing.T) {
	Assert := assert.With(t)

	field := Int32("name", -2147483647)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Int32)
	Assert.That(field.IntValue).IsEqualTo(-2147483647)
}

func TestInt64(t *testing.T) {
	Assert := assert.With(t)

	field := Int64("name", -9223372036854775807)

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.Int64)
	Assert.That(field.IntValue).IsEqualTo(-9223372036854775807)
}

func TestString(t *testing.T) {
	Assert := assert.With(t)

	field := String("name", "value")

	Assert.That(field.Name).IsEqualTo("name")
	Assert.That(field.Type).IsEqualTo(reflect.String)
	Assert.That(field.StringValue).IsEqualTo("value")
}

func TestErr(t *testing.T) {
	Assert := assert.With(t)

	field := Err(errors.New("sample error"))

	Assert.That(field.Name).IsEqualTo("error")
	Assert.That(field.Type).IsEqualTo(reflect.String)
	Assert.That(field.StringValue).IsEqualTo("sample error")
}

func TestField_Json(t *testing.T) {
	Assert := assert.With(t)
	var field Field

	field = Bool("name", true)
	Assert.That(field.Json()).IsEqualTo("\"name\":true")

	field = Int8("name", -127)
	Assert.That(field.Json()).IsEqualTo("\"name\":-127")

	field = Int16("name", -32767)
	Assert.That(field.Json()).IsEqualTo("\"name\":-32767")

	field = Int32("name", -2147483647)
	Assert.That(field.Json()).IsEqualTo("\"name\":-2147483647")

	field = Int64("name", -9223372036854775807)
	Assert.That(field.Json()).IsEqualTo("\"name\":-9223372036854775807")

	field = String("name", "value")
	Assert.That(field.Json()).IsEqualTo("\"name\":\"value\"")

	err := errors.New("error")
	field = Err(err)
	Assert.That(field.Json()).IsEqualTo("\"error\":\"error\"")
}

func TestField_String(t *testing.T) {
	Assert := assert.With(t)
	var field Field

	field = Bool("name", true)
	Assert.That(field.String()).IsEqualTo("name=true")

	field = Int8("name", -127)
	Assert.That(field.String()).IsEqualTo("name=-127")

	field = Int16("name", -32767)
	Assert.That(field.String()).IsEqualTo("name=-32767")

	field = Int32("name", -2147483647)
	Assert.That(field.String()).IsEqualTo("name=-2147483647")

	field = Int64("name", -9223372036854775807)
	Assert.That(field.String()).IsEqualTo("name=-9223372036854775807")

	field = String("name", "value")
	Assert.That(field.String()).IsEqualTo("name=value")

	err := errors.New("error")
	field = Err(err)
	Assert.That(field.String()).IsEqualTo("error=error")
}

func BenchmarkField_Json(b *testing.B) {
	var field Field

	err := errors.New("error message")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		field = Bool("name", true)
		field.Json()

		field = Int8("name", -127)
		field.Json()

		field = Int16("name", -32767)
		field.Json()

		field = Int32("name", -2147483647)
		field.Json()

		field = Int64("name", -9223372036854775807)
		field.Json()

		field = String("name", "value")
		field.Json()

		field = Err(err)
		field.Json()
	}
}

func BenchmarkField_String(b *testing.B) {
	var field Field

	err := errors.New("error message")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		field = Bool("name", true)
		field.String()

		field = Int8("name", -127)
		field.String()

		field = Int16("name", -32767)
		field.String()

		field = Int32("name", -2147483647)
		field.String()

		field = Int64("name", -9223372036854775807)
		field.String()

		field = String("name", "value")
		field.String()

		field = Err(err)
		field.String()
	}
}
