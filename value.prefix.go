package simex

import "strings"

type valuePrefix struct {
	jsonStrings
}

func (v *valuePrefix) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propPrefix)
}

func (v *valuePrefix) push(container map[string]interface{}) {
	v.pushProp(container, propPrefix)
}

func (v *valuePrefix) process(input string, backward bool) (string, *Error) {
	str := input
	for index, prefix := range v.value {
		if len(prefix) <= 0 {
			continue
		}
		if backward {
			end := strings.LastIndex(str, prefix)
			if end >= 0 {
				str = str[0:end]
			} else {
				var location string
				if v.array {
					location = propPrefix.locationAt(index)
				} else {
					location = propPrefix.location()
				}
				return "", &Error{msgExtract, location}
			}
		} else {
			start := strings.Index(str, prefix)
			if start >= 0 {
				str = str[start+len(prefix):]
			} else {
				var location string
				if v.array {
					location = propPrefix.locationAt(index)
				} else {
					location = propPrefix.location()
				}
				return "", &Error{msgExtract, location}
			}
		}
	}
	return str, nil
}
