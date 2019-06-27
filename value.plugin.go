package simex

type valuePlugin struct {
	jsonStrings
}

func (v *valuePlugin) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propPlugin)
}

func (v *valuePlugin) push(container map[string]interface{}) {
	v.pushProp(container, propPlugin)
}

func (v *valuePlugin) process(input string, plugin Plugin) (string, *Error) {
	if !v.set {
		return input, nil
	}
	str := input
	var err error
	for index, name := range v.value {
		str, err = plugin(str, name)
		if err != nil {
			var location string
			if v.array {
				location = propPlugin.locationAt(index)
			} else {
				location = propPlugin.location()
			}
			return "", &Error{err.Error(), location}
		}
	}
	return str, nil
}
