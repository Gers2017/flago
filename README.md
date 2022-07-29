# Flago
### 🐉 Simple and Flexible Command line flag parser
<br>

[![Stars](https://img.shields.io/github/stars/Gers2017/flago.svg?style=for-the-badge)](https://github.com/Gers2017/flago/stargazers)
[![Issues](https://img.shields.io/github/issues/Gers2017/flago.svg?style=for-the-badge)](https://github.com/Gers2017/flago/issues)
[![Contributors](https://img.shields.io/github/contributors/Gers2017/flago?style=for-the-badge)](https://github.com/Gers2017/flago/graphs/contributors)
[![LICENSE](https://img.shields.io/github/license/Gers2017/flago.svg?style=for-the-badge)](./LICENSE)
[![Pkg.go](https://img.shields.io/badge/reference-12c9c0?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/Gers2017/flago)


## Install ⭐️

```sh
go get github.com/Gers2017/flago
```

## Basic Usage 🔥
Import flago
```go
import (
    "github.com/Gers2017/flago"
)
```

Populate the flagset
```go
get := flago.NewFlagSet("get")
get.Bool("all", false)
get.Switch("verbose") // Same as get.Bool("verbose", false)
```
### Builder
```go
get := flago.NewFlagSet("get").
    Bool("all", false).
    Switch("verbose") // Same as get.Bool("verbose", false)
```
### Using the Cli struct
```go
cli := flago.NewCli()
cli.Handle(get, func(fs *flago.FlagSet) error {
    HandleFlagset(fs) // do something with parsed get
    return nil
})

if err := cli.Execute(os.Args); err != nil {
    log.Fatal(err)
}
```

```go 
func HandleFlagset(fs *flago.FlagSet) {
    if fs.Bool("help") {
		fmt.Println("Some helpful help message")
		return
	}

    // Do something...
    fmt.Println(todos)
}
```
### Without the Cli struct
Parse the arguments into flags
```go
// os.Args = []string{ "cli", "all", "help" }
if err := get.ParseFlags(os.Args[1:]); err != nil {
    log.Fatal(err)
    return
}
```

Then use the parsed flagset
```go
if get.IsParsed("all") {
    if get.Bool("help") {
        fmt.Println("Some helpful help message")
        return
    }
    // Do something...
    fmt.Println(todos)
}
```

## Demo 🐲
A complete example can be found **[here](./demo/demo.go)**

### New FlagSet
```go
get := flago.NewFlagSet("get")
```

### Add flags
```go
get.SetSwitch("verbose")
get.Int("x", 0)
```

### Check if a flag was parsed
It's highly recommended to use this method to check first if a flag was parsed correctly.
```go
get.IsParsed("your-flag-name")
```
The `FlagSet.[Bool, Int, Float, Str]` methods are just a shortcut for:

### Get values
```go
verbose := get.Bool("verbose")
x := get.Int("x")
```

If the flag name inside the getter method is not registered in the flagset, you'll get an error at runtime.
```go
wrong := get.Bool("some-invalid-flag")
```

## About the API

### Why so many strings? Isn't that error-prone?
1) The `FlagSet.[Bool, Int, Float, Str]` method can raise an error at runtine (use `FlagSet.IsParsed` to avoid this)

2)  A note on golang's generics
    Behind the scenes flago uses maps + generics + type parsing
    The `Flag struct` contains a `Value` property of type `any`.
    
    Because if we try to use generics we'd need to declare a map for every type of flag inside Flagset, and `Flags map[string]*Flag` wouldn't work anymore leading to repeated code.

    The flag module in the standard library solves this by using pointers to the underliying values.