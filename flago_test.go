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

func TestHelpParsing(t *testing.T) {
	args := split("all radian help")
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Bool("all", false, SET_ALL_HELP)
	set.Float("radian", 0.0, SET_RADIAN_HELP)
	set.SetStyle(MODERN)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	flag_names := split("all radian")

	for _, flag_name := range flag_names {
		if !set.HasFlag(flag_name) {
			t.Errorf("set flagset should have the flag \"%s\"", flag_name)
		}
	}

	if !set.IsHelp {
		t.Errorf("Help for \"%s\" flag should be true", set.Name)
	}
}

func TestModernParsing(t *testing.T) {
	args := split("score 300 radian 6.28 title lazy-dog help")
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Bool("all", false, SET_ALL_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 0, SET_RADIAN_HELP)
	set.Str("title", "", SET_TITLE_HELP)
	set.SetStyle(MODERN)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	if !set.IsHelp {
		t.Error("Error at parsing help modifier")
	}

	f_tests := []flagTest{
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"title", set.GetStr("title"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag %s should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestUnixParsing(t *testing.T) {
	args := split("nil score=300 radian=6.28 title=lazy-dog help")
	set := NewFlagSet("set", SET_CMD_HELP)
	set.Bool("all", false, SET_ALL_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 0, SET_RADIAN_HELP)
	set.Str("title", "", SET_TITLE_HELP)
	set.SetStyle(UNIX)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	if !set.IsHelp {
		t.Error("Error at parsing help modifier")
	}

	f_tests := []flagTest{
		{"score", set.GetInt("score"), 300},
		{"radian", set.GetFloat("radian"), 6.28},
		{"title", set.GetStr("title"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Wrong value for flag %s Expected: \"%v\" Got: \"%v\"", f_test.name, f_test.expected, f_test.got)
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
