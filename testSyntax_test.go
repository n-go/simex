package simex_test

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/n-go/simex"
)

// [syntax,location?]

func TestSyntax(t *testing.T) {
	var items [][]json.RawMessage
	err := json.Unmarshal([]byte(testSyntaxJSON), &items)
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range items {
		var e simex.Expression
		var location *string

		if len(item) > 1 {
			var value string
			if err := json.Unmarshal(item[1], &value); err != nil {
				t.Error(err)
				continue
			}
			location = &value
		}
		if err := json.Unmarshal(item[0], &e); err != nil {
			if location != nil {
				if !strings.HasSuffix(err.Error(), *location) {
					t.Error("error location does not match.\nMessage: ", err, "\nExpected: ", *location)
				}
			} else {
				t.Error(err)
			}
			continue
		}
		if location != nil {
			t.Error("should fail with location: " + *location + "\n" + string(item[0]))
			continue
		}

		data, err := json.Marshal(&e)
		if err != nil {
			t.Error(err)
			continue
		}

		if !bytes.Equal(data, item[0]) {
			t.Error("should sucessful stringify.\nExpected:\n" + string(item[0]) + "\n\nActual:\n" + string(data))
		}
	}
}
