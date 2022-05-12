package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
							0   1  flags...
	os.Args = tudu get --help -A -T=joe --number 42 --percentage 0.685

	now the problem of flags:
	flags := args[2:] = --help -A --title "joe" --number 42 --percentage 0.685
*/

const (
	ROOT_HELP      = "<Prints root help>"
	GET_ALL_HELP   = "<Prints get all help>"
	GET_SCORE_HELP = "<Prints get score help>"
	GET_PI_HELP    = "<Prints get pi help>"
	GET_TITLE_HELP = "<Prints get title help>"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		return
	}

	action := args[1]
	args_to_parse := args[2:]

	get := NewFlagSet("get")
	get.Bool("all", false)
	get.Str("title", "")
	get.Bool("help", false)
	get.Int("score", 0)
	get.Float("pi", 3.1416)

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

			fmt.Println("Score:", get.GetInt("score"))

		} else if get.HasFlag("pi") {

			if isHelp {
				fmt.Println(GET_PI_HELP)
				return
			}

			fmt.Println("Pi is:", get.GetFloat("pi"))

		} else if isHelp {
			fmt.Println(ROOT_HELP)
		}
	}
}

func (fs *FlagSet) ParseFlags(_args []string) error {
	argscp := Copy(_args)

	// clean args
	for i, v := range argscp {
		argscp[i] = Clean(v)
	}

	for i, arg := range argscp {
		if !fs.IsFlagName(arg) { // Ignore unknown flags
			continue
		}

		flag_name := arg

		f := fs.Flags[flag_name]

		switch f.Datatype {
		case "bool":
			f.Value = true
		case "string":
			f_value, err := fs.GetNextValue(_args, i, flag_name)
			if err != nil {
				return err
			}
			f.Value = f_value
		case "int":
			f_value, err := fs.GetNextValue(_args, i, flag_name)
			if err != nil {
				return err
			}
			value, _ := strconv.Atoi(f_value)
			f.Value = value
		case "float":
			f_value, err := fs.GetNextValue(_args, i, flag_name)
			if err != nil {
				return err
			}
			value, _ := strconv.ParseFloat(f_value, 64)
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
