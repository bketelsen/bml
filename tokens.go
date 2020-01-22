package bml

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF // End of File
	WS // white space
	DELIMOPEN // block delimiter
	DELIMCLOSE

	// Literals
	IDENT // identifiers, names

	// Blocks
	BLOCK
	DOCUMENT // 
	CHAPTER    // 
	PARAGRAPH
	LIST 
	LISTELEMENT
	ADMONITION
	NOTE
	QUOTE
	CAPTION
	RAW // raw type=html - embed HTML code in output

	// Media
	IMAGE
	AUDIO
	VIDEO


	IMPORT // import file=myfile.bml
	INPUT // prompted input
	OUTPUT // resulting output


	// INLINE Nodes

	BOLD
	ITALIC
	CODE // inline source code
	LINK // href
	XREF // cross reference
	NEWLINE 
	SPACE // &nbsp;


	COMMENT 
	CONSTANT
	


	// Keywords
	SELECT
	FROM

	
	ASTERISK
	COMMA
	
)