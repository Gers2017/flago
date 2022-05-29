package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	flago "github.com/Gers2017/flago"
)

const (
	GET_CMD_HELP      = "USAGE: get [ all | primary | by-title | help ]\nAvailable flags: all, primary, by-title, help"
	GET_ALL_HELP      = "Help for all flag: [ALL] [OPTIONS...]\nOptions: sort-by, reverse"
	GET_PRIMARY_HELP  = "Help for primary flag: [PRIMARY] [OPTIONS...]\nPrints the todo with highest priority"
	GET_BY_TITLE_HELP = "Help for by-title flag: [BY-TITLE] [OPTIONS...]\nPrints the todo with the matching title"
)

type Todo struct {
	title       string
	description string
	priority    int
	completed   bool
}

func (t *Todo) String() string {
	return fmt.Sprintf("[%s]%s -> %v\n%s", t.title, strings.Repeat("!", t.priority), t.completed, t.description)
}

func main() {
	get := flago.NewFlagSet("get")
	get.Bool("all", false)
	get.Bool("reverse", false)
	get.Str("sort-by", "title")
	get.Str("by-title", "")
	get.Bool("primary", false)
	get.Bool("help", false)

	args := os.Args
	if len(args) <= 2 {
		fmt.Println(GET_CMD_HELP)
		return
	}

	action := args[1]
	args_to_parse := args[2:]
	todos := []Todo{
		{"foo", "describes foo", 3, false},
		{"bar", "describes bar", 8, false},
		{"yeet", "describes yeet", 2, false},
		{"boo", "describes boo", 7, true},
	}

	if action != get.Name {
		fmt.Printf("Invalid action %s\n", action)
		fmt.Println(GET_CMD_HELP)
		return
	}

	if err := get.ParseFlags(args_to_parse); err != nil {
		fmt.Println(err)
		return
	}

	if get.IsParsed("all") {

		HandleAll(get, todos)

	} else if get.IsParsed("by-title") {

		HandleByTitle(get, todos)

	} else if get.IsParsed("primary") {

		HandleByPrimary(get, todos)

	} else {
		fmt.Println(GET_CMD_HELP)
	}
}

func HandleAll(fs *flago.FlagSet, todos []Todo) {
	if fs.GetBool("help") {
		fmt.Println(GET_ALL_HELP)
		return
	}

	is_reverse := fs.GetBool("reverse")
	sort_by := fs.GetStr("sort-by")

	if sort_by == "title" {
		sort.Slice(todos, func(i, j int) bool {
			if is_reverse {
				return todos[i].title > todos[j].title
			}
			return todos[i].title < todos[j].title
		})
	} else if sort_by == "priority" {
		sort.Slice(todos, func(i, j int) bool {
			if is_reverse {
				return todos[i].priority < todos[j].priority
			}
			return todos[i].priority > todos[j].priority // highest priority
		})
	}

	for _, x := range todos {
		println(x.String())
	}
}

func HandleByTitle(fs *flago.FlagSet, todos []Todo) {
	if fs.GetBool("help") {
		fmt.Println(GET_BY_TITLE_HELP)
		return
	}

	title := fs.GetStr("by-title")
	title = strings.ToLower(title)
	filtered := make([]Todo, 0)

	for _, todo := range todos {
		if strings.ToLower(todo.title) == title {
			filtered = append(filtered, todo)
		}
	}
	if len(filtered) > 0 {
		fmt.Println(filtered[0].String())
	} else {
		fmt.Printf("No title \"%s\" in todos\n", title)
	}
}

func HandleByPrimary(fs *flago.FlagSet, todos []Todo) {
	if fs.GetBool("help") {
		fmt.Println(GET_PRIMARY_HELP)
		return
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].priority > todos[j].priority // highest priority
	})

	fmt.Println(todos[0].String())
}
