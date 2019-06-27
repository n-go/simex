package simex

import "strings"

type valueSuffix struct {
	jsonStrings
}

func (v *valueSuffix) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propSuffix)
}

func (v *valueSuffix) push(container map[string]interface{}) {
	v.pushProp(container, propSuffix)
}

func (v *valueSuffix) process(input string, backward bool) (string, *Error) {
	str := input
	suffixed := false
	for _, suffix := range v.value {
		if len(suffix) <= 0 {
			suffixed = true
			break
		}
		if backward {
			start := strings.LastIndex(str, suffix)
			if start >= 0 {
				str = str[start+len(suffix):]
				suffixed = true
				break
			}
		} else {
			end := strings.Index(str, suffix)
			if end >= 0 {
				str = str[0:end]
				suffixed = true
				break
			}
		}
	}
	if !suffixed && len(v.value) > 0 {
		return "", &Error{msgExtract, propSuffix.location()}
	}
	return str, nil
}
