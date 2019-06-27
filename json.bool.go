package simex

type jsonBool struct {
	set   bool
	value bool
}

func (v *jsonBool) pullProp(container map[string]interface{}, p prop) *Error {
	v.set = false
	v.value = false

	if x, ok := container[p.name()]; ok {
		v.set = true
		if x != nil {
			if v.value, ok = x.(bool); ok {
				return nil
			}
		}
		return &Error{p.error(), p.location()}
	}
	return nil
}

func (v *jsonBool) pushProp(container map[string]interface{}, p prop) {
	if v.set {
		container[p.name()] = v.value
	}
}
