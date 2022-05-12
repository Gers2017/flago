package main

type ParseStyle int

const (
	MODERN ParseStyle = 1
	UNIX   ParseStyle = 2
)

type FlagSet struct {
	Name        string
	Parsed      bool
	ParsedFlags map[string]bool
	Flags       map[string]*Flag
}

type Flag struct {
	Name     string
	Value    any
	Datatype string
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		Name:        name,
		Parsed:      false,
		ParsedFlags: make(map[string]bool),
		Flags:       make(map[string]*Flag),
	}
}

func NewFlag(name string, defaultValue any, datatype string) *Flag {
	return &Flag{
		Name:     name,
		Value:    defaultValue,
		Datatype: datatype,
	}
}

func (fs *FlagSet) IsFlagName(name string) bool {
	_, ok := fs.Flags[name]
	return ok
}

func (fs *FlagSet) HasFlag(name string) bool {
	_, ok := fs.ParsedFlags[name]
	return ok
}

func (fs *FlagSet) GetNextValue(args_to_parse []string, current_index int, flag_name string) (string, error) {
	f_value, _ := GetArg(args_to_parse, current_index+1)

	if f_value == "" {
		return "", NewEmptyFlagValueError(flag_name)
	}

	if fs.IsFlagName(f_value) {
		return "", NewInvalidFlagValueError(flag_name, f_value)
	}
	return f_value, nil
}

func (fs *FlagSet) _addFlag(name string, f *Flag) {
	fs.Flags[name] = f
}

func (fs *FlagSet) Int(name string, init int) {
	fs._addFlag(name, NewFlag(name, init, "int"))
}

func (fs *FlagSet) Float(name string, init float64) {
	fs._addFlag(name, NewFlag(name, init, "float"))
}

func (fs *FlagSet) Bool(name string, init bool) {
	fs._addFlag(name, NewFlag(name, init, "bool"))
}

func (fs *FlagSet) Str(name string, init string) {
	fs._addFlag(name, NewFlag(name, init, "string"))
}

func TryGetType[T any](v any) T {
	t, ok := v.(T)
	if !ok {
		var def T
		return def
	}
	return t
}

func (fs *FlagSet) GetBool(key string) bool {
	f, ok := fs.Flags[key]
	if !ok {
		return false
	}

	return TryGetType[bool](f.Value)
}

func (fs *FlagSet) GetInt(key string) int {
	f, ok := fs.Flags[key]
	if !ok {
		return 0
	}
	return TryGetType[int](f.Value)
}

func (fs *FlagSet) GetFloat(key string) float64 {
	f, ok := fs.Flags[key]
	if !ok {
		return 0
	}
	return TryGetType[float64](f.Value)
}

func (fs *FlagSet) GetStr(key string) string {
	f, ok := fs.Flags[key]
	if !ok {
		return ""
	}
	return TryGetType[string](f.Value)
}
