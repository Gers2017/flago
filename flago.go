package flago

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	args_copy := copy(args_to_parse)

	// clean args
	for i, v := range args_copy {
		args_copy[i] = clean(v)
	}

	for i, arg := range args_copy {
		var flag_name string
		var flag_value string
		var is_help bool

		if fs.Style == MODERN_STYLE {

			flag_name, flag_value, is_help = parseFlagModernStyle(arg, args_copy, i)

		} else if fs.Style == UNIX_STYLE {

			flag_name, flag_value, is_help = parseFlagUnixStyle(arg)

		}

		if !fs.isFlagName(flag_name) && !is_help {
			continue // Skip if flag_name is not a flag and is not help
		}

		flag := fs.Flags[flag_name]

		if is_help {
			fs.IsHelp = true
			fs.setFlagAsParsed(flag_name)
			continue // skip checking flag_value and assigning values
		}

		f_err := fs.validateFlagValue(flag_name, flag_value)

		switch flag.Datatype {
		case "bool":
			flag.Value = true
		case "string":
			if f_err != nil {
				return f_err
			}

			flag.Value = flag_value
		case "int":
			if f_err != nil {
				return f_err
			}

			value, err := parseInt(flag_value)
			if err != nil {
				return newParseError(flag_value, "int")
			}

			flag.Value = value
		case "float":
			if f_err != nil {
				return f_err
			}

			value, err := parseFloat(flag_value)
			if err != nil {
				return newParseError(flag_value, "float")
			}

			flag.Value = value
		default:
			return newUnexpectedDataTypeError(string(flag.Datatype), flag.Name)
		}

		fs.setFlagAsParsed(flag_name)
	}

	fs.Parsed = true
	return nil
}

func parseFlagModernStyle(arg string, args_copy []string, i int) (flag_name string, flag_value string, is_help bool) {
	flag_name = arg
	flag_value = getNextValue(args_copy, i)
	is_help = isHelpValue(flag_value)
	return
}

func parseFlagUnixStyle(arg string) (flag_name string, flag_value string, is_help bool) {
	flag_name, flag_value = extractValues(arg)
	is_help = isHelpValue(flag_name)
	return
}
