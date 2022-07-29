package flago

import (
	"fmt"
	"strings"
	"testing"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func TestCli(t *testing.T) {
	args := split("cliname get --all")
	get := NewFlagSet("get").SetBool("all", false)
	add := NewFlagSet("add")

	cli := NewCli()

	cli.Handle(get, func(fs *FlagSet) error {
		if !fs.Bool("all") {
			return fmt.Errorf("Flag \"--all\" should be true")
		}

		return nil
	})

	cli.Handle(add, func(fs *FlagSet) error {
		if fs.Bool("help") {
			fmt.Println("print add help")
		}

		return fmt.Errorf("add subcommand should not be accessed")
	})

	if err := cli.Execute(args); err != nil {
		t.Error(err)
	}
}

func TestParsing(t *testing.T) {
	type FlagTest struct {
		name     string
		got      any
		expected any
	}

	args := split("all int 404 float 6.28 name lazy-dog")
	get := NewFlagSet("get").
		SetBool("all", false).
		SetInt("int", 0).
		SetFloat("float", 0.0).
		SetStr("name", "")

	err := get.ParseFlags(args)
	if err != nil {
		t.Error(err)
	}

	f_tests := []FlagTest{
		{"all", get.Bool("all"), true},
		{"int", get.Int("int"), 404},
		{"float", get.Float("float"), 6.28},
		{"name", get.Str("name"), "lazy-dog"},
	}

	for _, f_test := range f_tests {
		if f_test.got != f_test.expected {
			t.Errorf("Flag [%s] should equal \"%v\" got \"%v\"", f_test.name, f_test.expected, f_test.got)
		}
	}
}

func TestIsFlag(t *testing.T) {
	set := NewFlagSet("set")
	set.SetInt("score", 0)
	set.SetStr("title", "")
	set.SetFloat("radian", 0.0)

	flag_names := split("score title radian")

	for _, name := range flag_names {
		if !set.hasFlag(name) {
			t.Errorf("%s flag should be registered", name)
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

type pair struct {
	key   string
	value string
}

func (p *pair) compare(b *pair) bool {
	return (p.key == b.key && p.value == b.value)
}
func TestIterator(t *testing.T) {
	test_pairs := []pair{
		{"a", "1"},
		{"b", "2"},
		{"c", "3"},
	}

	fs := NewFlagSet("test")
	fs.SetInt("a", 0)
	fs.SetInt("b", 0)
	fs.SetInt("c", 0)

	iter := newFlagIterator(split("a 1 b 2 c 3"))

	pairs := []pair{}
	for !iter.is_empty() {
		key, ok := iter.next()
		if !ok {
			t.Error("Iterator attempted to get next \"key\", which doesn't exits")
		}

		if !fs.hasFlag(key) {
			continue
		}

		value, ok := iter.next()
		if !ok {
			t.Error("Iterator attempted to get next \"value\", which doesn't exits")
		}

		pairs = append(pairs, pair{key, value})
	}

	if len(test_pairs) != len(pairs) {
		t.Errorf("Missing pairs from %v, expected: %v", pairs, test_pairs)
	}

	for i, expected := range test_pairs {
		got := pairs[i]
		if !expected.compare(&got) {
			t.Errorf("Expected: %v Got: %v", expected, got)
		}
	}
}

func TestMissingValueError(t *testing.T) {
	args := split("int 404 name")
	get := NewFlagSet("get")
	get.SetInt("int", 0)
	get.SetStr("name", "")

	err := get.ParseFlags(args)
	if err == nil {
		t.Errorf("Should return a MissingValueError")
	}
}
