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

func TestModernParsing(t *testing.T) {
	args := strings.Split("nil score 300 radian 6.28 name lazy-dog help", " ")
	set := NewFlagSet("set")
	set.Int("score", 0)
	set.Float("radian", 0)
	set.Str("name", "")
	set.Bool("help", false)
	set.SetStyle(MODERN)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	f_tests := []flagTest{
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"name", set.GetStr("name"), "lazy-dog"},
		{"help", set.GetBool("help"), true},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag %s should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestUnixParsing(t *testing.T) {
	args := strings.Split("nil score=300 radian=6.28 name=lazy-dog help", " ")
	set := NewFlagSet("set")
	set.Int("score", 0)
	set.Float("radian", 0)
	set.Str("name", "")
	set.Bool("help", false)
	set.SetStyle(UNIX)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	f_tests := []flagTest{
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"name", set.GetStr("name"), "lazy-dog"},
		{"help", set.GetBool("help"), true},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Wrong value for flag %s Expected: \"%v\" Got: \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestExtractValues(t *testing.T) {
	extract_tests := []struct {
		flag    string
		f_name  string
		f_value string
	}{
		{"title=The-lazy-dog", "title", "The-lazy-dog"},
		{"score=300", "score", "300"},
		{"title=The_lazy_fox", "title", "The_lazy_fox"},
		{"title=the=lazy=parrot=2", "title", "the=lazy=parrot=2"},
		{"title_wrong", "title_wrong", ""},
	}

	for _, test := range extract_tests {
		f_name, f_value := extractValues(test.flag)

		if f_name != test.f_name {
			t.Errorf("Wrong extracted flag name. Expected: \"%v\" Got: \"%v\"", test.f_name, f_name)
		}

		if f_value != test.f_value {
			t.Errorf("Wrong extracted flag value. Expected: \"%v\" Got: \"%v\"", test.f_value, f_value)
		}
	}
}

func TestIsFlag(t *testing.T) {
	get := NewFlagSet("get")
	get.Int("score", 0)
	get.Float("radian", 0)
	get.Str("name", "")
	get.Bool("help", false)

	flag_names := []string{"score", "radian", "name", "help"}

	for _, f_name := range flag_names {
		if !get.isFlagName(f_name) {
			t.Fatalf("%s flag should be regsitered", f_name)
		}
	}
}

func TestValidateFlagValue(t *testing.T) {
	my_flag := "score"
	other_flag := "size"
	get := NewFlagSet("get")
	get.Int(my_flag, 0)
	get.Float(other_flag, 39.89)

	if err := get.validateFlagValue(my_flag, other_flag); err == nil {
		t.Errorf("Invalid value [%s]. Flag cannot be a value", other_flag)
	}

	if err := get.validateFlagValue(my_flag, ""); err == nil {
		t.Errorf("Invalid flag value for %s. Flag cannot be empty", my_flag)
	}
}

func TestInvalidDataType(t *testing.T) {
	args := []string{"foo"}
	get := NewFlagSet("get")
	get.addFlag("foo", NewFlag("foo", false, "nil"))

	if err := get.ParseFlags(args); err == nil {
		t.Error("ParseFlags should return an UnexpectedDataTypeError")
	}
}
