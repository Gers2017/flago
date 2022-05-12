package flago

type ParseStyle string

const (
	MODERN ParseStyle = "MODERN"
	UNIX   ParseStyle = "UNIX"
)

type FlagSet struct {
	Name        string
	Parsed      bool
	ParsedFlags map[string]bool
	Flags       map[string]*Flag
	Style       ParseStyle
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
		Style:       MODERN,
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

func (fs *FlagSet) SetStyle(style ParseStyle) {
	fs.Style = style
}

func (fs *FlagSet) _validateFlagValue(f_name, f_value string) error {
	if f_value == "" {
		return NewEmptyFlagValueError(f_name)
	}

	// Checks if the f_value is another flag (except for the help flag) only in MODERN style
	if fs.IsFlagName(f_value) && fs.Style == MODERN && f_value != "help" {
		return NewInvalidFlagAsValueError(f_name, f_value)
	}

	return nil
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
