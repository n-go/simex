package simex

import (
	"log"
	"strings"
)

type valueArray struct {
	set       bool
	separator valueSeparator
	omit      valueOmit
	item      valueItem
}

func (v *valueArray) pull(container map[string]interface{}) *Error {
	v.set = false
	array, ok := container[propArray.name()]
	if !ok {
		return nil
	}
	if array == nil {
		return &Error{propArray.error(), propArray.location()}
	}
	arrayContainer, ok := array.(map[string]interface{})
	if !ok {
		return &Error{propArray.error(), propArray.location()}
	}

	v.set = true
	if err := v.separator.pull(arrayContainer); err != nil {
		return err.prepend(propArray.location())
	}
	if err := v.omit.pull(arrayContainer); err != nil {
		return err.prepend(propArray.location())
	}
	if err := v.item.pull(arrayContainer); err != nil {
		return err.prepend(propArray.location())
	}
	for name := range arrayContainer {
		if name != propSeparator.name() &&
			name != propOmit.name() &&
			name != propItem.name() {
			log.Println(`redundant property "` + name + `" in array`)
		}
	}
	return nil
}

func (v *valueArray) push(container map[string]interface{}) {
	if !v.set {
		return
	}
	c := make(map[string]interface{})
	v.separator.push(c)
	v.omit.push(c)
	v.item.push(c)
	container[propArray.name()] = c
}

func (v *valueArray) process(input string, plugin Plugin, convert Convert) ([]interface{}, *Error) {
	ps := v.separator.split(input)
	results := make([]interface{}, 0)
	for _, p := range ps {
		if len(p) == 0 && v.omit.value {
			continue
		}
		if !v.item.set {
			results = append(results, p)
			continue
		}
		var errs []*Error
		for index, slice := range v.item.value {
			if slice == nil {
				errs = nil
				break
			}
			result, err := slice.process(p, plugin, convert)
			if err != nil {
				if v.item.array {
					err.prepend(propArray.location() + propItem.locationAt(index))
				} else {
					err.prepend(propArray.location() + propItem.location())
				}
				errs = append(errs, err)
			} else {
				results = append(results, result)
				errs = nil
				break
			}
		}
		if errs != nil && len(errs) > 0 {
			if v.item.array && len(errs) > 1 {
				locations := []string{}
				for _, err := range errs {
					locations = append(locations, err.location)
				}
				return nil, &Error{msgExtract, strings.Join(locations, "\n")}
			}
			return nil, errs[0]
		}
	}
	return results, nil
}
