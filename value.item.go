package simex

// ("item": slice | [slice | null] )?
type valueItem struct {
	set   bool
	value []*jsonSlice
	array bool
}

func (v *valueItem) pull(container map[string]interface{}) *Error {
	v.set = false
	x, ok := container[propItem.name()]
	if !ok {
		return nil
	}
	if x == nil {
		return &Error{propItem.error(), propItem.location()}
	}

	v.set = true
	v.value = []*jsonSlice{}
	if item, ok := x.(map[string]interface{}); ok {
		s := &jsonSlice{}
		v.array = false
		if err := s.fromJSON(item, "item"); err != nil {
			return err.prepend(propItem.location())
		}
		v.value = append(v.value, s)
		return nil
	} else if array, ok := x.([]interface{}); ok {
		v.array = true
		for index, i := range array {
			if i == nil {
				v.value = append(v.value, nil)
			} else if item, ok := i.(map[string]interface{}); ok {
				s := &jsonSlice{}
				if err := s.fromJSON(item, "item"); err != nil {
					return err.prepend(propItem.locationAt(index))
				}
				v.value = append(v.value, s)
			} else {
				return &Error{propItem.error(), propItem.locationAt(index)}
			}
		}
		return nil
	}
	return &Error{propItem.error(), propItem.location()}
}

func (v *valueItem) push(container map[string]interface{}) {
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
		container[propItem.name()] = values
	} else {
		container[propItem.name()] = v.value[0].toJSON()
	}
}
