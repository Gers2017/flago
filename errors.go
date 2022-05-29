package flago

import (
	"errors"
	"fmt"
)

func newParseTypeError(value, datatype string) error {
	return errors.New(fmt.Sprintf("Error at parsing value \"%s\" to datatype %s\n", value, datatype))
}

func newUnknownDataTypeError(datatype, name string) error {
	return errors.New(fmt.Sprintf("Unexpected datatype \"%s\" for flag of name %s", datatype, name))
}

type invalidFlagAsValueError struct {
	name  string
	value string
}

func newInvalidFlagAsValueError(name, value string) *invalidFlagAsValueError {
	return &invalidFlagAsValueError{name, value}
}

func (err *invalidFlagAsValueError) Error() string {
	return fmt.Sprintf("Invalid value for flag: \"%s\", cannot use flag \"%s\" as value", err.name, err.value)
}
