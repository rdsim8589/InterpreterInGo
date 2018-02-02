package token

type TokenType string

type Token struct {
	Type    TokenType //string type allows multiple types *boiler plate*
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL" //specifies unknown token
	EOF     = "EOF"     //tell parser when to stop

	//	Identifiers	+	literals
	IDENT = "IDENT" //	add,	foobar,	x,	y,	...
	INT   = "INT"   //	1343456

	//	Operators
	EQ       = "=="
	NOT_EQ   = "!="
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	//	Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//	Keywords
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)
