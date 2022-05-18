# Flago Experiment

## Unstable flago api
It mainly uses callback functions to setup the sub command actions of a given flag set.
- Drops _UNIX style_ parsing in favor of a better developer experience
- Tests aren't working yet
- It would probably rewrite the entire FlagSet - Flag structure and the ParseFlags function
- The Possibility of using pointer inside the actions is there
- Using multiple maps for the basic types is under testing


```go
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
```

### Example with pointers
```go
name := potato.Str("name", "", "Gives a name to your potato", func(f *flago.Flag) {
  if potato.IsHelp {
    fmt.Println(f.HelpText)
    return
  }

  fmt.Println("Potato's name:", name)

})
```

### Example of map-for-type approach
```go
potato.Str("name", "", "Gives a name to your potato", func(f *flago.Flag) {
  if potato.IsHelp {
    fmt.Println(f.HelpText)
    return
  }

  name := f.Value // string type provided by the flag (generic flag)

  fmt.Println("Potato's name:", name)

})
```

### Example of pointers and direct action call
```go
my_name_flag := potato.Str("name", "", "Gives a name to your potato", func(f *flago.Flag) {
	if potato.IsHelp {
		fmt.Println(f.HelpText)
		return
	}

	name := f.Value

	fmt.Println("Potato's name:", name)

})

my_name_flag.Action(my_name_flag)
```