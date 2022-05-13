package flago

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	fs.Parsed = true
	args_copy := copy(args_to_parse)

	// clean args
	for i, v := range args_copy {
		args_copy[i] = clean(v)
	}

	for i, arg := range args_copy {
		var flag_name string
		var f_value string

		if fs.Style == MODERN {
			// --- MODERN STYLE

			flag_name = arg
			if !fs.isFlagName(flag_name) { // Skip if is not a flag
				continue
			}
			f_value = getNextValue(args_copy, i)

			// ----
		} else if fs.Style == UNIX {
			// --- UNIX STYLE

			flag_name, f_value = extractValues(arg)
			if !fs.isFlagName(flag_name) { // Skip if is not a flag
				continue
			}

			// ----
		}

		f_err := fs.validateFlagValue(flag_name, f_value)
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

			value, err := parseInt(f_value)
			if err != nil {
				return newParseError(f_value, "int")
			}

			f.Value = value
		case "float":
			if f_err != nil {
				return f_err
			}

			if is_skip_parse {
				break
			}

			value, err := parseFloat(f_value)
			if err != nil {
				return newParseError(f_value, "float")
			}

			f.Value = value
		default:
			return newUnexpectedDataTypeError(f.Datatype, f.Name)
		}

		fs.ParsedFlags[flag_name] = true
	}

	return nil
}
