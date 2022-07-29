package flago

import (
	"fmt"
)

func newParseTypeError(value, datatype string) error {
	return fmt.Errorf("Error at parsing value \"%s\" to datatype %s\n", value, datatype)
}

func newUnknownDataTypeError(datatype, name string) error {
	return fmt.Errorf("Unexpected datatype \"%s\" for flag of name %s", datatype, name)
}

func newMissingValueError(name string, i int) error {
	return fmt.Errorf("The key: %s is trying to access missing item at index: %d", name, i)
}

func newCliError(cli *Cli) error {
	return fmt.Errorf("Available Subcommands: %v", cli.flagsetKeys())
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
