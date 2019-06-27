package simex

import (
	"strconv"
)

type prop int

const (
	propArray prop = iota
	propBackward
	propBetween
	propBy
	propConvert
	propObject
	propHas
	propItem
	propOmit
	propPrefix
	propPlugin
	propRoot
	propSeparator
	propSlice
	propSuffix
	propTrim
	propWith
)

type propTuple struct {
	name  string
	error string
}

func (p prop) name() string {
	return pts[p].name
}

func (p prop) error() string {
	return pts[p].error
}

func (p prop) location() string {
	return "." + pts[p].name
}

func (p prop) locationAt(index int) string {
	return "." + pts[p].name + "[" + strconv.Itoa(index) + "]"
}

var pts = []propTuple{
	{"array", "Property \"array\" must be an object."},
	{"backward", "Property \"backward\" must be true or false."},
	{"between", "Property \"between\" must be an object or an array of object."},
	{"by", "Property \"by\" must be a string."},
	{"convert", "Property \"convert\" must be a string."},
	{"object", "Property \"object\" must be an object."},
	{"has", "Property \"has\" must be a string."},
	{"item", "Property \"item\" must be an object or an array of object and null."},
	{"omit", "Property \"omit\" must be true or false."},
	{"prefix", "Property \"prefix\" must be either a string or an array of strings."},
	{"plugin", "Property \"plugin\" must be a string or an array of string."},
	{"root", "Property \"root\" must be an object."},
	{"separator", "Property \"separator\" must be either a string or an array of strings."},
	{"slice", "Property \"slice\" must be an object or an array of object."},
	{"suffix", "Property \"suffix\" must be either a string or an array of strings."},
	{"trim", "Property \"trim\" must be boolean."},
	{"with", "Property \"with\" must be either a string or an array of strings."},
}
