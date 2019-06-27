package simex

import "strings"

type valueSeparator struct {
	jsonStrings
}

func (v *valueSeparator) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propSeparator)
}

func (v *valueSeparator) push(container map[string]interface{}) {
	v.pushProp(container, propSeparator)
}

func (v *valueSeparator) split(input string) []string {
	ps := []string{input}
	if v.set && len(v.value) > 0 {
		for _, s := range v.value {
			if len(s) > 0 {
				next := []string{}
				for _, p := range ps {
					next = append(next, strings.Split(p, s)...)
				}
				ps = next
			}
		}
	}
	return ps
}
