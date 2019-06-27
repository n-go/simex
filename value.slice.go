package simex

import (
	"strings"
)

// ("slice": slice | [slice] )

type valueSlice struct {
	set   bool
	value []jsonSlice
	array bool
}

func (v *valueSlice) pull(container map[string]interface{}) *Error {
	v.set = false
	x, ok := container[propSlice.name()]
	if !ok {
		return nil
	}
	if x == nil {
		return &Error{propSlice.error(), propSlice.location()}
	}

	v.set = true
	v.value = []jsonSlice{}
	if item, ok := x.(map[string]interface{}); ok {
		var s jsonSlice
		v.array = false
		if err := s.fromJSON(item, "slice"); err != nil {
			return err.prepend(propSlice.location())
		}
		v.value = append(v.value, s)
		return nil
	} else if array, ok := x.([]interface{}); ok {
		v.array = true
		for index, i := range array {
			if i == nil {
				return &Error{propSlice.error(), propSlice.locationAt(index)}
			}
			if item, ok := i.(map[string]interface{}); ok {
				var s jsonSlice
				if err := s.fromJSON(item, "slice"); err != nil {
					return err.prepend(propSlice.locationAt(index))
				}
				v.value = append(v.value, s)
			} else {
				return &Error{propSlice.error(), propSlice.locationAt(index)}
			}
		}
		return nil
	}
	return &Error{propSlice.error(), propSlice.location()}
}

func (v *valueSlice) push(container map[string]interface{}) {
	if !v.set {
		return
	}
	if v.array {
		values := make([]map[string]interface{}, 0)
		for _, s := range v.value {
			values = append(values, s.toJSON())
		}
		container[propSlice.name()] = values
	} else {
		container[propSlice.name()] = v.value[0].toJSON()
	}
}

func (v *valueSlice) process(input string, plugin Plugin, convert Convert) (interface{}, *Error) {
	if !v.set {
		return input, nil
	}
	errs := []*Error{}
	for index, slice := range v.value {
		result, err := slice.process(input, plugin, convert)
		if err == nil {
			return result, nil
		}
		if v.array {
			err.prepend(propSlice.locationAt(index))
		} else {
			err.prepend(propSlice.location())
		}
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		err := errs[0]
		if v.array {
			locations := []string{}
			for _, err := range errs {
				locations = append(locations, err.location)
			}
			err = &Error{msgExtract, strings.Join(locations, "\n")}
		}
		return nil, err
	}
	return input, nil
}
