package simex

type jsonStrings struct {
	set   bool
	value []string
	array bool
}

func (v *jsonStrings) pullProp(container map[string]interface{}, p prop) *Error {
	v.set = false
	v.value = []string{}
	v.array = false

	if x, ok := container[p.name()]; ok {
		v.set = true
		if x != nil {
			if array, ok := x.([]interface{}); ok {
				v.array = true
				for index, item := range array {
					if i, ok := item.(string); ok {
						v.value = append(v.value, i)
					} else {
						return &Error{p.error(), p.locationAt(index)}
					}
				}
				return nil
			} else if i, ok := x.(string); ok {
				v.value = append(v.value, i)
				return nil
			}
		}
		return &Error{p.error(), p.location()}
	}
	return nil
}

func (v *jsonStrings) pushProp(container map[string]interface{}, p prop) {
	if v.set {
		if v.array {
			container[p.name()] = v.value
		} else {
			container[p.name()] = v.value[0]
		}
	}
}
