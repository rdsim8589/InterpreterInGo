package token

type TokenType string

type Token struct {
	Type    TokenType //string type allows multiple types *boiler plate*
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //specifies unknown token
	EOF     = "EOF"     //tell parser when to stop

	//	Identifiers	+	literals
	IDENT = "IDENT" //	add,	foobar,	x,	y,	...
	INT   = "INT"   //	1343456

	//	Operators
	ASSIGN = "="
	PLUS   = "+"

	//	Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//	Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
