package main

import (
	"fmt"
	"os"
)

/*
							0   1  flags...
	os.Args = tudu get --help -A -T=joe --number 42 --percentage 0.685

	now the problem of flags:
	flags := args[2:] = --help -A --title "joe" --number 42 --percentage 0.685
*/

const (
	ROOT_HELP       = "<Prints root help>"
	GET_ALL_HELP    = "<Prints get all help>"
	GET_SCORE_HELP  = "<Prints get score help>"
	GET_RADIAN_HELP = "<Prints get radian help>"
	GET_TITLE_HELP  = "<Prints get title help>"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		return
	}

	action := args[1]
	args_to_parse := args[1:]

	get := NewFlagSet("get")
	get.Bool("all", false)
	get.Str("title", "")
	get.Bool("help", false)
	get.Int("score", 0)
	get.Float("radian", 3.1416)
	get.SetStyle(MODERN)

	if action == "get" {
		err := get.ParseFlags(args_to_parse)
		if err != nil {
			fmt.Println(err)
			return
		}

		isHelp := get.GetBool("help")

		if get.HasFlag("all") {

			if isHelp {
				fmt.Println(GET_ALL_HELP)
				return
			}

			fmt.Println("Get all todos!")

		} else if get.HasFlag("title") {

			if isHelp {
				fmt.Println(GET_TITLE_HELP)
				return
			}

			fmt.Println("Getting todo by title:", get.GetStr("title"))

		} else if get.HasFlag("score") {

			if isHelp {
				fmt.Println(GET_SCORE_HELP)
				return
			}
			score := get.GetInt("score")
			fmt.Println("Score:", score)

		} else if get.HasFlag("radian") {

			if isHelp {
				fmt.Println(GET_RADIAN_HELP)
				return
			}

			fmt.Println("Radian is:", get.GetFloat("radian"))

		} else if isHelp {
			fmt.Println(ROOT_HELP)
		}
	}
}

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	fs.Parsed = true
	args_copy := Copy(args_to_parse)

	// clean args
	for i, v := range args_copy {
		args_copy[i] = Clean(v)
	}

	for i, arg := range args_copy {
		var flag_name string
		var f_value string

		if fs.Style == MODERN {
			// --- MODERN STYLE

			flag_name = arg
			if !fs.IsFlagName(flag_name) { // Skip non flag args
				continue
			}
			f_value = GetNextValue(args_copy, i)

			// ----
		} else if fs.Style == UNIX {
			// --- UNIX STYLE

			flag_name, f_value = ExtractValues(arg)
			if !fs.IsFlagName(flag_name) { // Skip non flag args
				continue
			}

			// ----
		}

		f_err := fs._validateFlagValue(flag_name, f_value)
		is_skip_parse := f_value == "help"

		f := fs.Flags[flag_name]

		switch f.Datatype {
		case "bool":
			f.Value = true
		case "string":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			f.Value = f_value
		case "int":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			value, err := ParseInt(f_value)
			if err != nil {
				return NewParseError(f_value, "int")
			}

			f.Value = value
		case "float":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			value, err := ParseFloat(f_value)
			if err != nil {
				return NewParseError(f_value, "float")
			}

			f.Value = value
		default:
			return NewUnexpectedDataTypeError(f.Datatype, f.Name)
		}

		fs.ParsedFlags[flag_name] = true
	}

	fmt.Println("---- ---- ----")
	for k, f := range fs.Flags {
		fmt.Printf("[%s] -> %v - %T\n", k, f.Value, f.Value)
	}
	fmt.Println("---- ---- ----")

	return nil
}
