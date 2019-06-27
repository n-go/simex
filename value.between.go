package simex

type valueBetween struct {
	set   bool
	value []jsonBetween
	array bool
}

func (v *valueBetween) pull(container map[string]interface{}) *Error {
	v.set = false
	between, ok := container[propBetween.name()]
	if !ok {
		return nil
	}
	if between == nil {
		return &Error{propBetween.error(), propBetween.location()}
	}

	v.set = true
	if object, ok := between.(map[string]interface{}); ok {
		var b jsonBetween
		err := b.fromJSON(object)
		if err != nil {
			return err.prepend(propBetween.location())
		}
		v.array = false
		v.value = append(v.value, b)
		return nil
	}
	if array, ok := between.([]interface{}); ok {
		v.array = true
		for index, item := range array {
			if object, ok := item.(map[string]interface{}); ok {
				var b jsonBetween
				err := b.fromJSON(object)
				if err != nil {
					return err.prepend(propBetween.locationAt(index))
				}
				v.value = append(v.value, b)
			} else {
				return &Error{propBetween.error(), propBetween.locationAt(index)}
			}
		}
		return nil
	}
	return &Error{propBetween.error(), propBetween.location()}
}

func (v *valueBetween) push(container map[string]interface{}) {
	if !v.set {
		return
	}
	if v.array {
		values := make([]map[string]interface{}, 0)
		for _, b := range v.value {
			values = append(values, b.toJSON())
		}
		container[propBetween.name()] = values
	} else {
		container[propBetween.name()] = v.value[0].toJSON()
	}
}

func (v *valueBetween) process(input string) (string, *Error) {
	var err *Error
	var str = input
	for index, b := range v.value {
		str, err = b.process(str)
		if err != nil {
			if v.array {
				err.prepend(propBetween.locationAt(index))
			} else {
				err.prepend(propBetween.location())
			}
			return "", err
		}
	}
	return str, nil
}
