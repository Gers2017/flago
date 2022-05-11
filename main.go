package main

import (
	"fmt"
	"os"
)

// args := strings.Split("tudu get --help -A -T=joe --number 42 --percentage 0.685", " ")
// args := strings.Split("gh clone repo --help --all --title joe --number 42 --percentage 0.685", " ")

func main() {
	args := os.Args
	if len(args) <= 1 {
		return
	}

	action := args[1]

	// myInt := get.Int(name, value, usage)

	get := NewFlagSet("get") // TODO: use an enum/const to select datatype, reflect - like a japanese mother sucker
	get.AddFlag("all", "all A", false, "bool")
	get.AddFlag("title", "title T", "", "string")
	get.AddFlag("age", "age AG", 0, "int")
	get.AddFlag("sad", "sad GZ", 0.0, "float")
	get.AddFlag("help", "help h", false, "bool") // return the value from a pointer

	add := NewFlagSet("add")
	add.AddFlag("item", "item IT", "", "string")
	add.AddFlag("help", "help h", false, "bool")

	if action == "get" {
		err := get.Parse(args[1:])
		fmt.Println(err)
	} else if action == "add" {
		add.Parse(args[1:])
	}
}
