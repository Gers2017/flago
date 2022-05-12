package flago

import (
	"errors"
	"fmt"
)

func NewParseError(value, datatype string) error {
	return errors.New(fmt.Sprintf("Error at parsing value \"%s\" to datatype %s\n", value, datatype))
}

func NewUnexpectedDataTypeError(datatype, flagName string) error {
	return errors.New(fmt.Sprintf("Unexpected datatype \"%s\" for flag of name %s", datatype, flagName))
}

func NewEmptyFlagValueError(flag string) error {
	return errors.New(fmt.Sprintf("Invalid value for flag \"%s\". Flag value cannot be empty\n", flag))
}

type InvalidFlagAsValueError struct {
	flag  string
	value string
}

func NewInvalidFlagAsValueError(flag, value string) *InvalidFlagAsValueError {
	return &InvalidFlagAsValueError{flag, value}
}

func (err *InvalidFlagAsValueError) Error() string {
	return fmt.Sprintf("Invalid value \"%s\" for flag: %s", err.value, err.flag)
}
