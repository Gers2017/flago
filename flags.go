package flago

type ParseStyle string

const (
	MODERN ParseStyle = "MODERN"
	UNIX   ParseStyle = "UNIX"
)

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
	Parsed      bool
	ParsedFlags map[string]bool
	Flags       map[string]*Flag
	Style       ParseStyle
	IsHelp      bool
	HelpText    string
}

type Flag struct {
	Name     string
	Value    any
	Datatype DataTypeName
	HelpText string
}

func NewFlagSet(name string, helptext string) *FlagSet {
	return &FlagSet{
		Name:        name,
		Parsed:      false,
		ParsedFlags: make(map[string]bool),
		Flags:       make(map[string]*Flag),
		Style:       MODERN,
		IsHelp:      false,
		HelpText:    helptext,
	}
}

func NewFlag[V FlagDataType](name string, value V, datatype DataTypeName, helptext string) *Flag {
	return &Flag{
		Name:     name,
		Value:    value,
		Datatype: datatype,
		HelpText: helptext,
	}
}

func (fs *FlagSet) isFlagName(name string) bool {
	_, ok := fs.Flags[name]
	return ok
}

func (fs *FlagSet) HasFlag(name string) bool {
	_, ok := fs.ParsedFlags[name]
	return ok
}

func (fs *FlagSet) SetStyle(style ParseStyle) {
	fs.Style = style
}

func (fs *FlagSet) GetFlag(flag_name string) (*Flag, bool) { // MAYBE RAISE AN ERROR FOR ACCESS AN UNKNOWN FLAG (?) <---
	f, ok := fs.Flags[flag_name]
	return f, ok
}

func (fs *FlagSet) validateFlagValue(flag_name, flag_value string) error {
	if flag_value == "" {
		return newEmptyFlagValueError(flag_name)
	}

	// Checks if the flag_value is another flag (only in MODERN style)
	if fs.isFlagName(flag_value) && fs.Style == MODERN {
		return newInvalidFlagAsValueError(flag_name, flag_value)
	}

	return nil
}

func (fs *FlagSet) setFlagAsParsed(flag_name string) {
	fs.ParsedFlags[flag_name] = true
}

func (fs *FlagSet) addFlag(name string, f *Flag) {
	fs.Flags[name] = f
}

func (fs *FlagSet) Int(name string, init int, helptext string) {
	fs.addFlag(name, NewFlag(name, init, INT, helptext))
}

func (fs *FlagSet) Float(name string, init float64, helptext string) {
	fs.addFlag(name, NewFlag(name, init, FLOAT, helptext))
}

func (fs *FlagSet) Bool(name string, init bool, helptext string) {
	fs.addFlag(name, NewFlag(name, init, BOOL, helptext))
}

func (fs *FlagSet) Str(name string, init string, helptext string) {
	fs.addFlag(name, NewFlag(name, init, STRING, helptext))
}

func tryGetType[T any](v any) T {
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

	return tryGetType[bool](f.Value)
}

func (fs *FlagSet) GetInt(key string) int {
	f, ok := fs.Flags[key]
	if !ok {
		return 0
	}
	return tryGetType[int](f.Value)
}

func (fs *FlagSet) GetFloat(key string) float64 {
	f, ok := fs.Flags[key]
	if !ok {
		return 0
	}
	return tryGetType[float64](f.Value)
}

func (fs *FlagSet) GetStr(key string) string {
	f, ok := fs.Flags[key]
	if !ok {
		return ""
	}
	return tryGetType[string](f.Value)
}
