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
get.Bool("help", false)
```

Parse the arguments into flags
```go
// os.Args = []string{ "cli", "all", "help" }
if err := get.ParseFlags(os.Args[1:]); err != nil {
    fmt.Println(err)
    return
}
```

Then use the updated flagset
```go
if get.IsParsed("all") {
    if get.GetBool("help") {
        fmt.Println("Some helpful help message")
        return
    }
    // Do something...
    fmt.Println(todos)
}
```


## Demo 🐲
A complete example can be found **[here](./demo/demo.go)**

## Extras

If you had a non-boolean flag you could use the IsParsed method to know if the flag is present as a valid flag.

It's highly recommended to use this method to check first if a flag was parsed correctly.
```go
get.IsParsed("your-flag-name")
```
The Get[Bool, Int, Float, Str] methods are just a shortcut for:
```go
f, _:= get.GetFlag("flag-name") // -> (f: Flag, exits: boolean)
all := f.ToBool() // ToInt(), ToFloat(), ToStr()
```

If the `flag-name` is not registered in the flagset, you'll get an error at runtime.
```go
// shortcut
all := get.GetBool("flag-name")
```

## About the API

### Why so many strings? Isn't that error-prone?
1) 
    Well that's on golang's generics.
    Behind the scenes flago is using maps + generics + type parsing
    The `Flag struct` contains a `Value` property of type `any`.
    
    Because if we try to use generics we'd need to declare a map for every type of flag inside Flagset, and `Flags map[string]*Flag` wouldn't work anymore leading to repeated code.

    The flag module in the standard library solves this by using pointers to the underliying values. This with a more than a few flags tends to get messy.

2) 
    Yes I agree! You can mess up a name.
    The good thing is that if you try to access a flag by a name that doesn't exist you'll get an error at runtime.