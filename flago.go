package flago

import "fmt"

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	fs.Parsed = true
	args_copy := Copy(args_to_parse)

	// clean args
	for i, v := range args_copy {
		args_copy[i] = Clean(v)
	}

	for i, arg := range args_copy {
		var flag_name string
		var f_value string

		if fs.Style == MODERN {
			// --- MODERN STYLE

			flag_name = arg
			if !fs.IsFlagName(flag_name) { // Skip non flag args
				continue
			}
			f_value = GetNextValue(args_copy, i)

			// ----
		} else if fs.Style == UNIX {
			// --- UNIX STYLE

			flag_name, f_value = ExtractValues(arg)
			if !fs.IsFlagName(flag_name) { // Skip non flag args
				continue
			}

			// ----
		}

		f_err := fs._validateFlagValue(flag_name, f_value)
		is_skip_parse := f_value == "help"

		f := fs.Flags[flag_name]

		switch f.Datatype {
		case "bool":
			f.Value = true
		case "string":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			f.Value = f_value
		case "int":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			value, err := ParseInt(f_value)
			if err != nil {
				return NewParseError(f_value, "int")
			}

			f.Value = value
		case "float":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			value, err := ParseFloat(f_value)
			if err != nil {
				return NewParseError(f_value, "float")
			}

			f.Value = value
		default:
			return NewUnexpectedDataTypeError(f.Datatype, f.Name)
		}

		fs.ParsedFlags[flag_name] = true
	}

	fmt.Println("---- ---- ----")
	for k, f := range fs.Flags {
		fmt.Printf("[%s] -> %v - %T\n", k, f.Value, f.Value)
	}
	fmt.Println("---- ---- ----")

	return nil
}
