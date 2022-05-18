package flago

import "fmt"

type ParseStyle string

const (
	MODERN_STYLE ParseStyle = "MODERN"
	UNIX_STYLE   ParseStyle = "UNIX"
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

type ActionFunc func(f *Flag)

type Flag struct {
	Name     string
	Value    any
	Datatype DataTypeName
	HelpText string
	Action   ActionFunc
}

func NewFlagSet(name string) *FlagSet {
	return &FlagSet{
		Name:        name,
		Parsed:      false,
		ParsedFlags: make(map[string]bool),
		Flags:       make(map[string]*Flag),
		Style:       MODERN_STYLE,
		IsHelp:      false,
		HelpText:    "",
	}
}

func NewFlag[V FlagDataType](name string, value V, datatype DataTypeName, helptext string, action ActionFunc) *Flag {
	return &Flag{
		Name:     name,
		Value:    value,
		Datatype: datatype,
		HelpText: helptext,
		Action:   action,
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

func (fs *FlagSet) SetHelpText(helptext string) {
	fs.HelpText = helptext
}

func (fs *FlagSet) GetFlagNames() []string {
	keys := make([]string, 0, len(fs.Flags))
	for k := range fs.Flags {
		keys = append(keys, k)
	}
	return keys
}

func (fs *FlagSet) Execute() {
	parsed := make([]string, 0, len(fs.ParsedFlags))
	for k := range fs.ParsedFlags {
		parsed = append(parsed, k)
	}

	if len(parsed) == 0 {
		fmt.Println(fs.HelpText)
	}

	for k := range fs.ParsedFlags {
		if !fs.isFlagName(k) {
			continue
		}

		f, _ := fs.Flags[k]
		f.Action(f)
	}

}

func (fs *FlagSet) validateFlagValue(flag_name, flag_value string) error {
	if flag_value == "" {
		return newEmptyFlagValueError(flag_name)
	}

	// Checks if the flag_value is another flag (only in MODERN style)
	if fs.isFlagName(flag_value) && fs.Style == MODERN_STYLE {
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

func (fs *FlagSet) Int(name string, init int, helptext string, action ActionFunc) {
	fs.addFlag(name, NewFlag(name, init, INT, helptext, action))
}

func (fs *FlagSet) Float(name string, init float64, helptext string, action ActionFunc) {
	fs.addFlag(name, NewFlag(name, init, FLOAT, helptext, action))
}

func (fs *FlagSet) Bool(name string, init bool, helptext string, action ActionFunc) {
	fs.addFlag(name, NewFlag(name, init, BOOL, helptext, action))
}

func (fs *FlagSet) Str(name string, init string, helptext string, action ActionFunc) {
	fs.addFlag(name, NewFlag(name, init, STRING, helptext, action))
}

func tryGetType[T FlagDataType](v any) T {
	t, ok := v.(T)
	if !ok {
		var init T
		return init
	}
	return t
}

func newEmptyFlag() *Flag {
	return NewFlag("", false, BOOL, "", func(f *Flag) {})
}

func (fs *FlagSet) GetFlag(flag_name string) (*Flag, bool) {
	_, ok := fs.ParsedFlags[flag_name]
	if !fs.isFlagName(flag_name) || !ok {
		return newEmptyFlag(), false
	}
	f, ok := fs.Flags[flag_name]
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
	f, ok := fs.Flags[key]
	if !ok {
		return 0
	}
	return f.ToInt()
}

func (fs *FlagSet) GetFloat(key string) float64 {
	f, ok := fs.Flags[key]
	if !ok {
		return f.ToFloat()
	}
	return tryGetType[float64](f.Value)
}

func (fs *FlagSet) GetBool(key string) bool {
	f, ok := fs.GetFlag(key)
	if !ok {
		return false
	}
	return f.ToBool()
}

func (fs *FlagSet) GetStr(key string) string {
	f, ok := fs.GetFlag(key)
	if !ok {
		return ""
	}
	return f.ToStr()
}
