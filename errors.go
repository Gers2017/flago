package main

import (
	"errors"
	"fmt"
)

func NewParseError(value, datatype string) error {
	return errors.New(fmt.Sprintf("Error at parsing %s to %s\n", value, datatype))
}

func NewUnexpectedDataTypeError(datatype, flagName string) error {
	return errors.New(fmt.Sprintf("Unexpected datatype \"%s\" for flag of name %s", datatype, flagName))
}

type InvalidFlagValueError struct {
	flag  string
	value string
}

func NewInvalidFlagValueError(flag, value string) *InvalidFlagValueError {
	return &InvalidFlagValueError{flag, value}
}

func (err *InvalidFlagValueError) Error() string {
	return fmt.Sprintf("Invalid value for flag \"%s\". Value \"%s\" is a registered flag", err.flag, err.value)
}