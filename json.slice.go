package simex

import (
	"log"
)

type jsonSlice struct {
	has     valueHas
	between valueBetween
	plugin  valuePlugin
	value   *interface{}
	slice   valueSlice
	array   valueArray
	object  valueObject
	convert valueConvert
}

func (v *jsonSlice) fromJSON(container map[string]interface{}, containerName string) *Error {
	var err *Error
	if err = v.has.pull(container); err != nil {
		return err
	}
	if err = v.between.pull(container); err != nil {
		return err
	}
	if err = v.plugin.pull(container); err != nil {
		return err
	}
	if value, ok := container["value"]; ok {
		v.value = &value
	}
	if err = v.slice.pull(container); err != nil {
		return err
	}
	if err = v.array.pull(container); err != nil {
		return err
	}
	if err = v.object.pull(container); err != nil {
		return err
	}
	if err = v.convert.pull(container); err != nil {
		return err
	}
	for name := range container {
		if name != propHas.name() &&
			name != propBetween.name() &&
			name != propPlugin.name() &&
			name != "value" &&
			name != propSlice.name() &&
			name != propArray.name() &&
			name != propObject.name() &&
			name != propConvert.name() {
			log.Println(`redundant property "` + name + `" in ` + containerName)
		}
	}
	return nil
}

func (v *jsonSlice) toJSON() map[string]interface{} {
	container := make(map[string]interface{})
	v.has.push(container)
	v.between.push(container)
	v.plugin.push(container)
	if v.value != nil {
		container["value"] = *v.value
	}
	v.slice.push(container)
	v.array.push(container)
	v.object.push(container)
	return container
}

func (v *jsonSlice) process(input string, plugin Plugin, convert Convert) (interface{}, *Error) {
	var str = input
	var err *Error
	if err = v.has.validate(str); err != nil {
		return nil, err
	}
	if str, err = v.between.process(str); err != nil {
		return nil, err
	}
	if v.plugin.set && plugin != nil {
		str, err = v.plugin.process(str, plugin)
		if err != nil {
			return nil, err
		}
	}
	if v.value != nil {
		return *v.value, nil
	} else if v.slice.set {
		return v.slice.process(str, plugin, convert)
	} else if v.array.set {
		return v.array.process(str, plugin, convert)
	} else if v.object.set {
		return v.object.process(str, plugin, convert)
	} else if v.convert.set && convert != nil {
		return v.convert.process(str, convert)
	}
	return str, nil
}
