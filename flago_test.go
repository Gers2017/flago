package flago

import (
	"strings"
	"testing"
)

type flagTest struct {
	name     string
	got      any
	expected any
}

func split(s string) []string {
	return strings.Split(s, " ")
}

func TestParsing(t *testing.T) {
	args := split("all int 404 float 6.28 name lazy-dog")
	get := NewFlagSet("get")
	get.Bool("all", false)
	get.Int("int", 0)
	get.Float("float", 0.0)
	get.Str("name", "")

	err := get.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	f_tests := []flagTest{
		{"all", get.GetBool("all"), true},
		{"int", get.GetInt("int"), 404},
		{"float", get.GetFloat("float"), 6.28},
		{"name", get.GetStr("name"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag [%s] should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestIsFlag(t *testing.T) {
	set := NewFlagSet("set")
	set.Int("score", 0)
	set.Str("title", "")
	set.Float("radian", 0.0)

	flag_names := split("score title radian")

	for _, name := range flag_names {
		if !set.isFlag(name) {
			t.Fatalf("%s flag should be regsitered", name)
		}
	}
}

func TestInvalidDataType(t *testing.T) {
	args := []string{"foo"}
	set := NewFlagSet("set")
	set.addFlag("foo", NewFlag("foo", false, "none"))

	if err := set.ParseFlags(args); err == nil {
		t.Error("ParseFlags should return an UnexpectedDataTypeError")
	}
}

func TestTryGetType(t *testing.T) {
	type TypeTest[T any] struct {
		value    T
		asInt    int
		asFloat  float64
		asBool   bool
		asString string
	}

	values := []TypeTest[any]{
		{"hello", 0, 0.0, false, "hello"}, // string
		{true, 0, 0.0, true, ""},          // bool
		{420, 420, 0.0, false, ""},        // int
		{33.56, 0, 33.56, false, ""},      // float
	}

	for _, v := range values {
		asInt := tryGetType[int](v.value)
		if asInt != v.asInt {
			t.Errorf("Expected: %d Got: %d", v.asInt, asInt)
		}

		asFloat := tryGetType[float64](v.value)
		if asFloat != v.asFloat {
			t.Errorf("Expected: %f Got: %f", v.asFloat, asFloat)
		}

		asBool := tryGetType[bool](v.value)
		if asBool != v.asBool {
			t.Errorf("Expected: %v Got: %v", v.asBool, asBool)
		}

		asString := tryGetType[string](v.value)
		if asString != v.asString {
			t.Errorf("Expected: %s Got: %s", v.asString, asString)
		}
	}
}
