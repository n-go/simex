package simex

import (
	"strings"
)

type valueTrim struct {
	jsonBool
}

func (v *valueTrim) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propTrim)
}

func (v *valueTrim) push(container map[string]interface{}) {
	v.pushProp(container, propTrim)
}

func (v *valueTrim) process(input string) string {
	if v.value {
		return strings.TrimSpace(input)
	}
	return input
}
