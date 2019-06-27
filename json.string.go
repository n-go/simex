package simex

type jsonString struct {
	set   bool
	value string
}

func (v *jsonString) pullProp(container map[string]interface{}, p prop) *Error {
	v.set = false
	v.value = ""

	if x, ok := container[p.name()]; ok {
		v.set = true
		if x != nil {
			if v.value, ok = x.(string); ok {
				return nil
			}
		}
		return &Error{p.error(), p.location()}
	}
	return nil
}

func (v *jsonString) pushProp(container map[string]interface{}, p prop) {
	if v.set {
		container[p.name()] = v.value
	}
}
