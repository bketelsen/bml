package bml_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bketelsen/bml"
)

// Ensure the parser can parse strings into Statement ASTs.


func TestParser_ParseBlock(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *bml.Block
		err  string
	}{
		// Single field statement
		{
			s: `<document />`,
			stmt: &bml.Block{
				Node: bml.Node{
					Type:    bml.BLOCK,
				},

			},
		},

		
		
	}

	for i, tt := range tests {
		stmt, err := bml.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nsblock mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}