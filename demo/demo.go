package main

import (
	"fmt"
	"os"

	flago "github.com/Gers2017/flago"
)

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

	get := flago.NewFlagSet("get")
	get.Bool("all", false)
	get.Str("title", "")
	get.Bool("help", false)
	get.Int("score", 0)
	get.Float("radian", 3.1416)
	get.SetStyle(flago.MODERN)

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
