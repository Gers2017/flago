package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type FlagSet struct {
	Name   string
	Parsed bool
	Args   []string
	Flags  map[string]*Flag
}

type Flag struct {
	Name         string
	Value        any
	DefaultValue any
	Datatype     string
}

func NewParseError(value, datatype string) error {
	return errors.New(fmt.Sprintf("Error at parsing %s to %s\n", value, datatype))
}

func NewEmptyFlagValueError(f string) error {
	return errors.New(fmt.Sprintf("Empty value after flag %s\n", f))
}

var parseError = errors.New("parse error")

func NewFlag(name string, defaultValue any, datatype string) *Flag {
	return &Flag{Name: name, Value: defaultValue, DefaultValue: defaultValue, Datatype: datatype}
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{Name: name, Parsed: false, Args: make([]string, 0), Flags: make(map[string]*Flag)}
}

func (f *FlagSet) AddFlag(name string, variants string, defaultValue any, datatype string) {
	flag := NewFlag(name, defaultValue, datatype)
	for _, variant := range strings.Split(variants, " ") {
		f.Flags[variant] = flag
	}
}

func (f *FlagSet) GetFlag(k string) (*Flag, bool) {
	v, ok := f.Flags[k]
	return v, ok
}

func (f *FlagSet) Parse(args []string) error {
	f.Parsed = false
	f.Args = args

	for _, arg := range args {
		if !IsFlag(arg) {
			continue
		}

		arg = Clean(arg)
		f_arg, f_value := ExtractValues(arg)

		f, exits := f.GetFlag(f_arg)

		if !exits {
			continue
		}

		switch f.Datatype {
		case "bool":
			f.Value = true

		case "string":
			if f_value == "" {
				return NewEmptyFlagValueError(f_arg)
			}

			f.Value = f_value

		case "int":
			v, err := strconv.ParseInt(f_value, 0, 64)

			if f_value == "" {
				return NewEmptyFlagValueError(f_arg)
			}

			if err != nil {
				return NewParseError(f_value, f.Datatype)
			}

			f.Value = v

		case "float":
			v, err := strconv.ParseFloat(f_value, 64)

			if f_value == "" {
				return NewEmptyFlagValueError(f_arg)
			}

			if err != nil {
				return NewParseError(f_value, f.Datatype)
			}

			f.Value = v
		}
	}

	for k, f := range f.Flags {
		fmt.Printf("[%s] -> %v - %T\n", k, f.Value, f.Value)
	}

	return nil
}
