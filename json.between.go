package simex

import "log"

type jsonBetween struct {
	backward valueBackward
	prefix   valuePrefix
	suffix   valueSuffix
	trim     valueTrim
}

func (v *jsonBetween) fromJSON(container map[string]interface{}) *Error {
	var err *Error
	if err = v.backward.pull(container); err != nil {
		return err
	}
	if err = v.prefix.pull(container); err != nil {
		return err
	}
	if err = v.suffix.pull(container); err != nil {
		return err
	}
	if err = v.trim.pull(container); err != nil {
		return err
	}
	for name := range container {
		if name != propBackward.name() &&
			name != propPrefix.name() &&
			name != propSuffix.name() &&
			name != propTrim.name() {
			log.Println(`redundant property "` + name + `" in between`)
		}
	}
	return nil
}

func (v *jsonBetween) toJSON() map[string]interface{} {
	container := make(map[string]interface{})
	v.backward.push(container)
	v.prefix.push(container)
	v.suffix.push(container)
	v.trim.push(container)
	return container
}

func (v *jsonBetween) process(input string) (string, *Error) {
	var str = input
	var err *Error
	str, err = v.prefix.process(str, v.backward.value)
	if err != nil {
		return "", err
	}
	str, err = v.suffix.process(str, v.backward.value)
	if err != nil {
		return "", err
	}
	return v.trim.process(str), nil
}
