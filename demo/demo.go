package main

import (
	"fmt"
	"os"

	flago "github.com/Gers2017/flago"
)

func main() {
	potato := flago.NewFlagSet("potato")

	potato.Str("set-name", "", "Gives a name to your potato", func(f *flago.Flag) {
		if potato.IsHelp {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Potato's name:", f.ToStr())

	})

	potato.Bool("angry", false, "Toggles the agression of the potato", func(f *flago.Flag) {
		if potato.IsHelp {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Is your potato angry?", f.ToBool())

	})

	potato.Int("shape", 0, "Rates the shape of the potato", func(f *flago.Flag) {
		if potato.IsHelp {
			fmt.Println(f.HelpText)
			return
		}

		fmt.Println("Shape rating:", f.ToInt())

	})

	potato.SetHelpText(fmt.Sprintf("Potato commands: %v\n", potato.GetFlagNames()))

	args := os.Args[1:]

	if err := potato.ParseFlags(args); err != nil {
		fmt.Printf("Invalid potato arguments: %s\n", err)
		return
	}

	potato.Execute()

}
