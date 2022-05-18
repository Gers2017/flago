package main

import (
	"fmt"
	"os"

	flago "github.com/Gers2017/flago"
)

const (
	SET_CMD_HELP    = "USAGE: set [all | score | radian | title]\nAvailable commands: all, score, radian, title, help"
	SET_ALL_HELP    = "<Prints set all help>"
	SET_SCORE_HELP  = "<Prints set score help>"
	SET_RADIAN_HELP = "<Prints set radian help>"
	SET_TITLE_HELP  = "<Prints set title help>"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		return
	}

	action := args[1]
	args_to_parse := args[1:]

	set := flago.NewFlagSet("set", SET_CMD_HELP) // help flag is no longer needed
	set.Bool("all", false, SET_ALL_HELP)
	set.Str("title", "", SET_TITLE_HELP)
	set.Int("score", 0, SET_SCORE_HELP)
	set.Float("radian", 3.1416, SET_RADIAN_HELP)
	set.SetStyle(flago.MODERN_STYLE)

	if action != "set" {
		fmt.Printf("Invalid action %s\n", action)
		return
	}

	err := set.ParseFlags(args_to_parse)
	if err != nil {
		fmt.Println(err)
		return
	}

	isHelp := set.IsHelp

	fmt.Println("args:", args_to_parse)
	fmt.Println("isHelp:", isHelp)

	if set.HasFlag("all") {

		f, ok := set.GetFlag("all")

		if isHelp || !ok {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Set all todos:", f.ToBool())

	} else if set.HasFlag("title") {

		f, ok := set.GetFlag("title")

		if isHelp || !ok {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Getting todo by title:", f.ToStr())

	} else if set.HasFlag("score") {

		f, ok := set.GetFlag("score")

		if isHelp || !ok {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Score:", f.ToInt())

	} else if set.HasFlag("radian") {

		f, ok := set.GetFlag("radian")

		if isHelp || !ok {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Radian:", f.ToFloat())

	} else {
		fmt.Println(set.HelpText)
	}
}
