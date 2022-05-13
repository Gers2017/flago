package flago

import (
	"strings"
	"testing"
)

func TestModernParsing(t *testing.T) {
	args := strings.Split("get score 300", " ")
	set := NewFlagSet("get")
	set.Int("score", 0)
	set.SetStyle(MODERN)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}
	if !set.HasFlag("score") {
		t.Errorf("HasFlag(score) should equal true. Got %v", set.HasFlag("score"))
	}
	if set.GetInt("score") != 300 {
		for k, f := range set.ParsedFlags {
			t.Log(k, f)
		}
		t.Errorf("flag score should equal 300. Got %v", set.GetInt("score"))
	}
}

func TestUnixParsing(t *testing.T) {
	args := strings.Split("get score=300", " ")
	set := NewFlagSet("get")
	set.Int("score", 0)
	set.SetStyle(UNIX)

	err := set.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	if !set.HasFlag("score") {
		t.Errorf("HasFlag(score) should equal true. Got %v", set.HasFlag("score"))
	}

	if set.GetInt("score") != 300 {
		t.Errorf("flag score should equal 300. Got %v", set.GetInt("score"))
	}
}

func TestExtractValues(t *testing.T) {
	my_unix_flag := "title=The lazy dog"
	f_name, f_value := extractValues(my_unix_flag)

	if f_name != "title" {
		t.Errorf("Flag name should be title got \"%s\"", f_name)
	}

	if f_value == "" {
		t.Error("Flag value shouldn't be empty")
	}
}

func TestIsFlag(t *testing.T) {
	get := NewFlagSet("get")
	get.Int("score", 0)

	if !get.isFlagName("score") {
		t.Fatal("Score flag should be reconized")
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
