package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
	GET_TITLE_HELP = "<Prints get title help>"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		return
	}
	ls := []string{"all", "non", "mom", "score", "help"}
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] < ls[j]
	})

	action := args[1]
	args_to_parse := args[2:]

	get := NewFlagSet("get")
	get.Int("score", 0)
	get.Bool("all", false)
	get.Str("title", "")
	get.Bool("help", false)
	get._addFlag("bubu", NewFlag("bubu", 9.0, "float"))

	if action == "get" {
		err := get.ParseFlags(args_to_parse)
		if err != nil {
			log.Fatalln(err)
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

		f := fs.Flags[arg]
		fs.ParsedFlags[arg] = true

		switch f.Datatype {
		case "bool":
			f.Value = true
		case "string":
			f_value, _ := GetArg(_args, i+1)
			if fs.IsFlagName(f_value) {
				return NewInvalidFlagValueError(arg, f_value)
			}
			f.Value = f_value
		case "int":
			f_value, _ := GetArg(_args, i+1)
			if fs.IsFlagName(f_value) {
				return NewInvalidFlagValueError(arg, f_value)
			}
			value, _ := strconv.Atoi(f_value)
			f.Value = value
		default:
			return NewUnexpectedDataTypeError(f.Datatype, f.Name)
		}
	}

	fmt.Println("---- ---- ----")
	for k, f := range fs.Flags {
		fmt.Printf("[%s] -> %v - %T\n", k, f.Value, f.Value)
	}
	fmt.Println("---- ---- ----")

	return nil
}
