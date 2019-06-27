package simex

type valueBackward struct {
	jsonBool
}

func (v *valueBackward) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propBackward)
}

func (v *valueBackward) push(container map[string]interface{}) {
	v.pushProp(container, propBackward)
}
