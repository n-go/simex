package simex

type valueOmit struct {
	jsonBool
}

func (v *valueOmit) pull(container map[string]interface{}) *Error {
	return v.pullProp(container, propOmit)
}

func (v *valueOmit) push(container map[string]interface{}) {
	v.pushProp(container, propOmit)
}
