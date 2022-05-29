# Flago
### 🐉 Simple and Flexible Command line flag parser

## Why so many strings? Isn't that error-prone?
1) Well that's on golang's generics, an other option is to a pointer to the flag value. With a few of flags is not a big deal but with more than a few it becomes messy.

2) Yes I agree! You can mess up a name. 
Good thing that if you try to access a flag by a name that doesn't exist you'' get an error at runtime.

## Simple example
```go
get := flago.NewFlagSet("get")
get.Bool("all", false)
get.Bool("help", false)

if get.IsParsed("all") {
    if get.GetBool("help") {
        fmt.Println("Some helpful help message")
        return
    }
    // Do something...
    fmt.Println(todos)
}
```

If you had a string flag you could use the IsParsed method to know if the flag is present as a valid flag (Returns false if ParseFlags returns an error)

It is recommended to use this method to check first if a flag  exist/was parsed correctly
```go
get.IsParsed("your-flag-name") // -> bool
```

## Demo
A complete example can be found [here](./demo/demo.go)

```go
f,_:= get.GetFlag("all")
all := f.ToBool()
// same but shorter
all := get.GetBool("all")
```
