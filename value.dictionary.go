package simex

import (
	"strconv"
	"strings"
)

type valueObject struct {
	set     bool
	members map[string]valueMember
}

func (v *valueObject) pull(container map[string]interface{}) *Error {
	v.set = false
	object, ok := container[propObject.name()]
	if !ok {
		return nil
	}
	if object == nil {
		return &Error{propObject.error(), propObject.location()}
	}
	objectContainer, ok := object.(map[string]interface{})
	if !ok {
		return &Error{propObject.error(), propObject.location()}
	}

	v.set = true
	v.members = map[string]valueMember{}
	for name := range objectContainer {
		var member valueMember
		if err := member.pull(objectContainer, name); err != nil {
			return err.prepend(propObject.location())
		}
		v.members[name] = member
	}
	return nil
}

func (v *valueObject) push(container map[string]interface{}) {
	if !v.set {
		return
	}
	c := make(map[string]interface{})
	for _, member := range v.members {
		member.push(c)
	}
	container[propObject.name()] = c
}

func (v *valueObject) process(input string, plugin Plugin, convert Convert) (map[string]interface{}, *Error) {
	results := make(map[string]interface{})
	for _, member := range v.members {
		var errs []*Error
		for index, slice := range member.value {
			if slice == nil {
				errs = nil
				break
			}
			result, err := slice.process(input, plugin, convert)
			if err != nil {
				if member.array {
					err.prepend("[" + strconv.Itoa(index) + "]")
				}
				err.prepend(propObject.location() + "[\"" + member.name + "\"]")
				errs = append(errs, err)
			} else {
				results[member.name] = result
				errs = nil
				break
			}
		}
		if errs != nil && len(errs) > 0 {
			if member.array && len(errs) > 1 {
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
