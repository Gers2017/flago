package flago

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	args_copy := copy(args_to_parse)

	for i, v := range args_copy {
		args_copy[i] = clean(v)
	}

	for i, arg := range args_copy {
		flag_name, flag_value := parseOneFlag(arg, args_copy, i, fs.Style)
		is_valid_flagname := fs.isFlagName(flag_name)

		if !is_valid_flagname {
			continue
		}

		next_arg, _ := getArg(args_copy, i+1)

		if isHelpValue(next_arg) && is_valid_flagname {
			fs.IsHelp = true
			fs.setFlagAsParsed(flag_name)
			continue
		}

		flag := fs.Flags[flag_name]
		err := parseType(fs, flag, flag_name, flag_value)
		if err != nil {
			return err
		}

		fs.setFlagAsParsed(flag_name)
	}

	fs.Parsed = true
	return nil
}

func parseOneFlag(arg string, args_copy []string, i int, style ParseStyle) (flag_name string, flag_value string) {
	if style == MODERN_STYLE {
		flag_name = arg
		nextValue, _ := getArg(args_copy, i+1)
		flag_value = nextValue
	} else if style == UNIX_STYLE {
		flag_name, flag_value = extractValues(arg)
	}
	return
}

func parseType(fs *FlagSet, flag *Flag, flag_name, flag_value string) error {
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

	return nil
}
