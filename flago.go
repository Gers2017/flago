package flago

func (fs *FlagSet) ParseFlags(args_to_parse []string) error {
	args_copy := copySlice(args_to_parse)

	// remove flag prefix from args
	for i, v := range args_copy {
		args_copy[i] = removeFlagPrefix(v)
	}

	// get the tokens from args
	tokens := getTokens(args_copy)

	// filter by flag names
	only_flags := make([]token, 0)
	for _, token := range tokens {
		if fs.isFlag(token.name) {
			only_flags = append(only_flags, token)
		}
	}

	tokens = only_flags

	// parse tokens into flag
	for _, token := range tokens {
		name := token.name
		value := token.value
		flag := fs.Flags[name]

		if flag.Datatype != "bool" && fs.isFlag(value) && value != "help" {
			return newInvalidFlagAsValueError(name, value)
		}

		switch flag.Datatype {
		case "bool":
			flag.Value = true
		case "string":
			flag.Value = value
		case "int":
			v, err := parseInt(value)
			if err != nil {
				return newParseTypeError(value, "int")
			}

			flag.Value = v

		case "float":
			v, err := parseFloat(value)
			if err != nil {
				return newParseTypeError(value, "float")
			}

			flag.Value = v

		default:
			return newUnknownDataTypeError(string(flag.Datatype), flag.Name)
		}
		fs.setAsParsed(name)
	}

	return nil
}

type token = struct {
	name  string
	value string
}

func getTokens(args []string) []token {
	tokens := make([]token, 0)
	for i := range args {
		tokens = append(tokens, token{name: getArg(args, i), value: getArg(args, i+1)})
	}

	return tokens
}
