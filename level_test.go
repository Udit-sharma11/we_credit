package log

import (
	"github.com/MuratSs/assert"
	"testing"
)

func TestLevel_String(t *testing.T) {
	var actual string
	var assert = assert.With(t)

	actual = DEBUG.String()
	assert.That(actual).IsEqualTo("DEBUG")

	actual = INFO.String()
	assert.That(actual).IsEqualTo("INFO")

	actual = WARN.String()
	assert.That(actual).IsEqualTo("WARN")

	actual = ERROR.String()
	assert.That(actual).IsEqualTo("ERROR")

	assert.ThatPanics(func() { Level(-1).String() })
}

func TestParse(t *testing.T) {
	var actual Level
	var assert = assert.With(t)

	actual = Parse("DEBUG")
	assert.That(actual).IsEqualTo(DEBUG)

	actual = Parse("INFO")
	assert.That(actual).IsEqualTo(INFO)

	actual = Parse("WARN")
	assert.That(actual).IsEqualTo(WARN)

	actual = Parse("ERROR")
	assert.That(actual).IsEqualTo(ERROR)

	assert.ThatPanics(func() { Parse("foobar") })
}
