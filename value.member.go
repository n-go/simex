package simex

import (
	"strconv"
)

// ("item": slice | [slice | null] )?
type valueMember struct {
	set   bool
	name  string
	value []*jsonSlice
	array bool
}

func (v *valueMember) pull(container map[string]interface{}, name string) *Error {
	v.set = false
	v.name = name

	member, ok := container[name]
	if !ok {
		return nil
	}
	if member == nil {
		return &Error{memberMessage(name), "[\"" + name + "\"]"}
	}

	v.set = true
	v.value = []*jsonSlice{}
	if item, ok := member.(map[string]interface{}); ok {
		s := &jsonSlice{}
		v.array = false
		if err := s.fromJSON(item, "object[\""+name+"\"]"); err != nil {
			return err.prepend("[\"" + name + "\"]")
		}
		v.value = append(v.value, s)
		return nil
	} else if array, ok := member.([]interface{}); ok {
		v.array = true
		for index, i := range array {
			if i == nil {
				v.value = append(v.value, nil)
			} else if item, ok := i.(map[string]interface{}); ok {
				s := &jsonSlice{}
				if err := s.fromJSON(item, "object[\""+name+"\"]"); err != nil {
					return err.prepend("[\"" + name + "\"][" + strconv.Itoa(index) + "]")
				}
				v.value = append(v.value, s)
			} else {
				return &Error{memberMessage(name), "[\"" + name + "\"][" + strconv.Itoa(index) + "]"}
			}
		}
		return nil
	}
	return &Error{memberMessage(name), "[\"" + name + "\"]"}
}

func (v *valueMember) push(container map[string]interface{}) {
	if !v.set {
		return
	}
	if v.array {
		values := make([]map[string]interface{}, 0)
		for _, s := range v.value {
			if s != nil {
				values = append(values, s.toJSON())
			} else {
				values = append(values, nil)
			}
		}
		container[v.name] = values
	} else {
		container[v.name] = v.value[0].toJSON()
	}
}
