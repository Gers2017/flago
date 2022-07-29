package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/Gers2017/flago"
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
	get := flago.NewFlagSet("get").
		SetSwitch("all"). // SetSwitch shorthand for SetBool(<flagname>, false)
		SetBool("reverse", false).
		SetStr("sort-by", "title").
		SetStr("by-title", "").
		SetBool("primary", false).
		SetSwitch("help") // Optional, cli adds a help flag if none was provided

	args := os.Args

	todos := []Todo{
		{"foo", "describes foo", 3, false},
		{"bar", "describes bar", 8, false},
		{"yeet", "describes yeet", 12, false},
		{"boo", "describes boo", 9, true},
	}

	cli := flago.NewCli()
	cli.Handle(get, func(fs *flago.FlagSet) error {
		HandleGet(fs, todos)
		return nil
	})

	if err := cli.Execute(args); err != nil {
		fmt.Println(err)
	}
}

func HandleGet(fs *flago.FlagSet, todos []Todo) {
	isAll := fs.IsParsed("all")
	isByTitle := fs.IsParsed("by-title")
	isPrimary := fs.IsParsed("primary")

	if isAll {
		HandleAll(fs, todos)
	} else if isByTitle {
		HandleByTitle(fs, todos)
	} else if isPrimary {
		HandleByPrimary(fs, todos)
	} else {
		fmt.Println(GET_CMD_HELP)
	}
}

func HandleAll(fs *flago.FlagSet, todos []Todo) {
	if fs.Bool("help") {
		fmt.Println(GET_ALL_HELP)
		return
	}

	is_reverse := fs.Bool("reverse")
	sort_by := fs.Str("sort-by")

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
	if fs.Bool("help") {
		fmt.Println(GET_BY_TITLE_HELP)
		return
	}

	title := strings.ToLower(fs.Str("by-title"))
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
	if fs.Bool("help") {
		fmt.Println(GET_PRIMARY_HELP)
		return
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].priority > todos[j].priority // highest priority
	})

	fmt.Println(todos[0].String())
}
