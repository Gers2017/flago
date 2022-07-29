package flago

func (fs *FlagSet) ParseFlags(args []string) error {
	// remove flag prefix from args
	for i, v := range args {
		args[i] = removeFlagPrefix(v)
	}

	iter := newFlagIterator(args)
	for !iter.is_empty() {
		key, ok := iter.next()
		if !fs.hasFlag(key) || !ok {
			continue
		}

		flag := fs.Flags[key]
		data_type := flag.Datatype

		if data_type == "bool" {
			flag.Value = true
		} else {
			value, ok := iter.next()
			if !ok {
				return newMissingValueError(key, iter.index)
			}

			if fs.hasFlag(value) && value != "help" {
				return newInvalidFlagAsValueError(key, value)
			}

			switch data_type {
			case "string":
				flag.Value = value

			case "int":
				int_value, err := parseInt(value)
				if err != nil {
					return newParseTypeError(value, "int")
				}

				flag.Value = int_value

			case "float":
				float_value, err := parseFloat(value)
				if err != nil {
					return newParseTypeError(value, "float")
				}

				flag.Value = float_value

			default:
				return newUnknownDataTypeError(string(data_type), flag.Name)
			}
		}

		fs.setAsParsed(key)
	}

	return nil
}

type flagIterator struct {
	args  []string
	index int
	max   int
}

func newFlagIterator(args []string) flagIterator {
	index := 0
	max := len(args)
	return flagIterator{args, index, max}
}

func (fi *flagIterator) next() (string, bool) {
	if fi.is_empty() {
		return "", false
	}

	s := fi.args[fi.index]
	fi.index += 1
	return s, true
}

func (fi *flagIterator) is_empty() bool {
	return fi.index >= fi.max
}
