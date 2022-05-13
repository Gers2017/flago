package flago

import (
	"errors"
	"fmt"
)

func newParseError(value, datatype string) error {
	return errors.New(fmt.Sprintf("Error at parsing value \"%s\" to datatype %s\n", value, datatype))
}

func newUnexpectedDataTypeError(datatype, flagName string) error {
	return errors.New(fmt.Sprintf("Unexpected datatype \"%s\" for flag of name %s", datatype, flagName))
}

func newEmptyFlagValueError(flag string) error {
	return errors.New(fmt.Sprintf("Invalid value for flag \"%s\". Flag value cannot be empty\n", flag))
}

type invalidFlagAsValueError struct {
	flag  string
	value string
}

func newInvalidFlagAsValueError(flag, value string) *invalidFlagAsValueError {
	return &invalidFlagAsValueError{flag, value}
}

func (err *invalidFlagAsValueError) Error() string {
	return fmt.Sprintf("Invalid value \"%s\" for flag: %s", err.value, err.flag)
}
