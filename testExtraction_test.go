package simex_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/n-go/simex"
)

// [syntax, in, out,location?]

func TestExtraction(t *testing.T) {
	var items [][]json.RawMessage
	err := json.Unmarshal([]byte(testExtractionJSON), &items)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		var e simex.Expression
		var input string
		var output []byte
		var location *string

		if err := json.Unmarshal(item[1], &input); err != nil {
			t.Error(err)
			continue
		}
		output = item[2]
		if len(item) > 3 {
			var value string
			if err := json.Unmarshal(item[3], &value); err != nil {
				t.Error(err)
				continue
			}
			location = &value
		}
		if err := json.Unmarshal(item[0], &e); err != nil {
			t.Error(err)
			continue
		}
		result, err := e.Extract(input, nil, nil)
		if err != nil {
			if location != nil {
				if err.Location() != *location {
					t.Error("error location does not match.\nExpected: ", *location, "\nActual: ", err.Location(), "\n", string(item[0]))
					continue
				}
			} else {
				t.Error(err, "\n", string(item[0]))
				continue
			}
		}
		actual, er := json.Marshal(result)
		if er != nil {
			t.Error(er)
			continue
		}
		if !bytes.Equal(actual, output) {
			t.Error("result mismatch.\nExpected: ", string(output), "\nActual: ", string(actual), "\n", string(item[0]))
		}
	}
}
