package simex

type valueConvert struct {
	jsonString
}

func (v *valueConvert) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propConvert)
}

func (v *valueConvert) push(container map[string]interface{}) {
	v.pushProp(container, propConvert)
}

func (v *valueConvert) process(input string, convert Convert) (interface{}, *Error) {
	if !v.set {
		return input, nil
	}
	value, err := convert(input, v.value)
	if err != nil {
		return nil, &Error{err.Error(), propConvert.location()}
	}
	return value, nil
}
