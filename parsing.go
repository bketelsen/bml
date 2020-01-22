package bml

import (
	"fmt"
	"io"
	"strings"
	"errors"
	
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct {
	Fields    []string
	TableName string
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (*Block, error) {


	 tok, lit := p.scanIgnoreWhitespace()
	 if tok != BLOCK {
		return nil, fmt.Errorf("found %q, expected BLOCK", lit)
	}


	bl, err := p.parseBlock(lit)
	fmt.Println(bl,err)
	if err != nil {
		return bl, err
	}


	// Return the successfully parsed statement.
	return bl, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) parseBlock(inner string) (b *Block, err error) {
	b = &Block{}
	b.Raw = inner
	s := NewScanner(strings.NewReader(inner[1:]))
	_,l := s.scanIdent()

	switch strings.ToUpper(l) {
	case "DOCUMENT":
		b.Type = DOCUMENT
	case "CHAPTER":
		b.Type = CHAPTER
	default:
		err = errors.New("Unknown block type")
	}

	return
}


// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }