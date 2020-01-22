package bml_test

import (
	"strings"
	"testing"

	"github.com/bketelsen/bml"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok bml.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: bml.EOF},
		{s: `#`, tok: bml.ILLEGAL, lit: `#`},
		{s: ` `, tok: bml.WS, lit: " "},
		{s: "\t", tok: bml.WS, lit: "\t"},
		{s: "\n", tok: bml.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: bml.ASTERISK, lit: "*"},

		{s: `<foo />`, tok: bml.BLOCK, lit: "<foo />"},

		// Identifiers
		{s: `foo`, tok: bml.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: bml.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: bml.FROM, lit: "FROM"},
		{s: `SELECT`, tok: bml.SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := bml.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}