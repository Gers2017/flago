package flago

type DataTypeName string

const (
	INT    DataTypeName = "int"
	FLOAT  DataTypeName = "float"
	BOOL   DataTypeName = "bool"
	STRING DataTypeName = "string"
)

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

func (fs *FlagSet) isFlag(name string) bool {
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

func (fs *FlagSet) Int(name string, init int) {
	fs.addFlag(name, NewFlag(name, init, INT))
}

func (fs *FlagSet) Float(name string, init float64) {
	fs.addFlag(name, NewFlag(name, init, FLOAT))
}

func (fs *FlagSet) Bool(name string, init bool) {
	fs.addFlag(name, NewFlag(name, init, BOOL))
}

func (fs *FlagSet) Str(name string, init string) {
	fs.addFlag(name, NewFlag(name, init, STRING))
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

func (fs *FlagSet) GetInt(key string) int {
	f := fs.Flags[key]
	return f.ToInt()
}

func (fs *FlagSet) GetFloat(key string) float64 {
	f := fs.Flags[key]
	return tryGetType[float64](f.Value)
}

func (fs *FlagSet) GetBool(key string) bool {
	f := fs.Flags[key]
	return f.ToBool()
}

func (fs *FlagSet) GetStr(key string) string {
	f := fs.Flags[key]
	return f.ToStr()
}
