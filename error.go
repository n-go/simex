package simex

import (
	"strings"
)

const msgExtract = "Provided input does not match the expression."

// Error represents error from building expression and extrating data.
type Error struct {
	message  string
	location string
}

func (e *Error) String() string {
	return e.message
}

// Location returns expression path where the error occurs.
func (e *Error) Location() string {
	return e.location
}

func (e *Error) Error() string {
	return e.message + " @ " + e.location
}

func (e *Error) prepend(location string) *Error {
	locs := strings.Split(e.location, "\n")
	if len(locs) > 1 {
		for i := 0; i < len(locs); i++ {
			locs[i] = location + locs[i]
		}
		e.location = strings.Join(locs, "\n")
	} else {
		e.location = location + e.location
	}
	return e
}

func memberMessage(name string) string {
	return "Property object[\"" + name + "\"] must be an object or an array of object and null."
}
