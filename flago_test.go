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

const (
	SET_CMD_HELP    = "USAGE: set [all | score | radian | title]\nAvailable commands: all, score, radian, title, help"
	SET_ALL_HELP    = "<Prints set all help>"
	SET_SCORE_HELP  = "<Prints set score help>"
	SET_RADIAN_HELP = "<Prints set radian help>"
	SET_TITLE_HELP  = "<Prints set title help>"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func TestIsHelpOnFlagSet(t *testing.T) {
	set := NewFlagSet("set", SET_CMD_HELP)
	set.SetStyle(MODERN_STYLE)
	set.ParseFlags(split("set help"))
	if !set.IsHelp {
		t.Errorf("IsHelp on set flagset should be %v. Instead got: %v", true, set.IsHelp)
	}
}

func TestModernParsing(t *testing.T) {
	args := split("set all score 300 radian 6.28 title lazy-dog help")
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Bool("all", false, SET_ALL_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 0, SET_RADIAN_HELP)
	set.Str("title", "", SET_TITLE_HELP)
	set.SetStyle(MODERN_STYLE)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	if !set.IsHelp {
		t.Error("Error at parsing help modifier")
	}

	f_tests := []flagTest{
		{"all", set.GetBool("all"), true},
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"title", set.GetStr("title"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag [%s] should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestUnixParsing(t *testing.T) {
	args := split("set all score=300 radian=6.28 title=lazy-dog help")
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Bool("all", false, SET_ALL_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 0, SET_RADIAN_HELP)
	set.Str("title", "", SET_TITLE_HELP)
	set.SetStyle(UNIX_STYLE)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	if !set.IsHelp {
		t.Error("Error at parsing help modifier")
	}

	f_tests := []flagTest{
		{"all", set.GetBool("all"), true},
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"title", set.GetStr("title"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag [%s] should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestExtractValues(t *testing.T) {
	extract_tests := []struct {
		flag       string
		flag_name  string
		flag_value string
	}{
		{"title=The-lazy-dog", "title", "The-lazy-dog"},
		{"score=300", "score", "300"},
		{"title=The_lazy_fox", "title", "The_lazy_fox"},
		{"title=the=lazy=parrot=2", "title", "the=lazy=parrot=2"},
		{"title_wrong", "title_wrong", ""},
	}

	for _, test := range extract_tests {
		flag_name, flag_value := extractValues(test.flag)

		if flag_name != test.flag_name {
			t.Errorf("Wrong extracted flag name. Expected: \"%v\" Got: \"%v\"", test.flag_name, flag_name)
		}

		if flag_value != test.flag_value {
			t.Errorf("Wrong extracted flag value. Expected: \"%v\" Got: \"%v\"", test.flag_value, flag_value)
		}
	}
}

func TestIsFlag(t *testing.T) {
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 0, SET_RADIAN_HELP)
	set.Str("title", "", SET_TITLE_HELP)

	flag_names := split("score radian title")

	for _, flag_name := range flag_names {
		if !set.isFlagName(flag_name) {
			t.Fatalf("%s flag should be regsitered", flag_name)
		}
	}
}

func TestValidateFlagValue(t *testing.T) {
	my_flag := "score"
	other_flag := "size"
	set := NewFlagSet("set", "Help for set flagset")
	set.Int(my_flag, 0, "Help text for my_flag!")
	set.Float(other_flag, 39.89, "Help text for other_flag!")

	if err := set.validateFlagValue(my_flag, other_flag); err == nil {
		t.Errorf("Invalid value [%s]. Flag cannot be a value", other_flag)
	}

	if err := set.validateFlagValue(my_flag, ""); err == nil {
		t.Errorf("Invalid flag value for %s. Flag cannot be empty", my_flag)
	}
}

func TestInvalidDataType(t *testing.T) {
	args := []string{"foo"}
	set := NewFlagSet("set", "")
	set.addFlag("foo", NewFlag("foo", false, "nil", ""))

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
