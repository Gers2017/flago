package flago

type DataTypeName string

const (
	INT    DataTypeName = "int"
	FLOAT  DataTypeName = "float"
	BOOL   DataTypeName = "bool"
	STRING DataTypeName = "string"
)

type HandlerFunc = func(fs *FlagSet) error

type Cli struct {
	flagsets   map[string]*FlagSet
	handlers   map[string]HandlerFunc
	subcommand string
}

func NewCli() *Cli {
	return &Cli{
		flagsets:   make(map[string]*FlagSet),
		handlers:   make(map[string]func(fs *FlagSet) error),
		subcommand: "",
	}
}

func (cli *Cli) Execute(args []string) error {
	if len(args) <= 2 {
		return newCliError(cli)
	}

	subcommand := args[1]
	fs, exits := cli.flagsets[subcommand]
	if !exits {
		return newCliError(cli)
	}

	if err := fs.ParseFlags(args[2:]); err != nil {
		return err
	}

	action, ok := cli.handlers[subcommand]
	if !ok {
		return newCliError(cli)
	}

	return action(fs)
}

func (cli *Cli) Handle(flagset *FlagSet, handler HandlerFunc) {
	if !flagset.hasFlag("bool") {
		flagset.SetSwitch("help")
	}

	cli.flagsets[flagset.Name] = flagset
	cli.handlers[flagset.Name] = handler
}

func (cli *Cli) flagsetKeys() []string {
	keys := make([]string, 0, len(cli.flagsets))
	for k := range cli.flagsets {
		keys = append(keys, k)
	}
	return keys
}

type FlagDataType interface {
	int | float64 | string | bool
}

type FlagSet struct {
	Name        string
	Flags       map[string]*Flag
	ParsedFlags map[string]bool
}

type Flag struct {
	Name     string
	Value    any
	Datatype DataTypeName
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		Name:        name,
		Flags:       make(map[string]*Flag),
		ParsedFlags: make(map[string]bool),
	}
}

func NewFlag[V FlagDataType](name string, value V, datatype DataTypeName) *Flag {
	return &Flag{
		Name:     name,
		Value:    value,
		Datatype: datatype,
	}
}

func (fs *FlagSet) hasFlag(name string) bool {
	_, ok := fs.Flags[name]
	return ok
}

func (fs *FlagSet) IsParsed(name string) bool {
	_, ok := fs.ParsedFlags[name]
	return ok
}

func (fs *FlagSet) setAsParsed(name string) {
	fs.ParsedFlags[name] = true
}

func (fs *FlagSet) addFlag(name string, f *Flag) {
	fs.Flags[name] = f
}

func (fs *FlagSet) SetInt(name string, init int) *FlagSet {
	fs.addFlag(name, NewFlag(name, init, INT))
	return fs
}

func (fs *FlagSet) SetFloat(name string, init float64) *FlagSet {
	fs.addFlag(name, NewFlag(name, init, FLOAT))
	return fs
}

func (fs *FlagSet) SetBool(name string, init bool) *FlagSet {
	fs.addFlag(name, NewFlag(name, init, BOOL))
	return fs
}

func (fs *FlagSet) SetSwitch(name string) *FlagSet {
	fs.addFlag(name, NewFlag(name, false, BOOL))
	return fs
}

func (fs *FlagSet) SetStr(name string, init string) *FlagSet {
	fs.addFlag(name, NewFlag(name, init, STRING))
	return fs
}

func tryGetType[T FlagDataType](v any) T {
	t, ok := v.(T)
	if !ok {
		var init T
		return init
	}
	return t
}

func (fs *FlagSet) GetFlag(name string) (*Flag, bool) {
	f, ok := fs.Flags[name]
	return f, ok
}

func (f *Flag) ToInt() int {
	return tryGetType[int](f.Value)
}

func (f *Flag) ToFloat() float64 {
	return tryGetType[float64](f.Value)
}

func (f *Flag) ToBool() bool {
	return tryGetType[bool](f.Value)
}

func (f *Flag) ToStr() string {
	return tryGetType[string](f.Value)
}

func (fs *FlagSet) Int(key string) int {
	f := fs.Flags[key]
	return f.ToInt()
}

func (fs *FlagSet) Float(key string) float64 {
	f := fs.Flags[key]
	return tryGetType[float64](f.Value)
}

func (fs *FlagSet) Bool(key string) bool {
	f := fs.Flags[key]
	return f.ToBool()
}

func (fs *FlagSet) Str(key string) string {
	f := fs.Flags[key]
	return f.ToStr()
}
