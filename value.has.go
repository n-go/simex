package simex

import (
	"strings"
)

type valueHas struct {
	jsonString
}

func (v *valueHas) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propHas)
}

func (v *valueHas) push(container map[string]interface{}) {
	v.pushProp(container, propHas)
}

func (v *valueHas) validate(input string) *Error {
	if len(v.value) > 0 && !strings.Contains(input, v.value) {
		return &Error{msgExtract, propHas.location()}
	}
	return nil
}
