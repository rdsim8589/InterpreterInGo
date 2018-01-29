package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;` //backtik ignore special char

	// test cases
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input) //what is this method?

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType { //checking the token against the tests
			t.Fatalf("tests[%d]	-	tokentype	wrong.	expected=%q,	got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d]	-	literal	wrong.	expected=%q,	got=%q",
				i, tt.expectedLiteral, tok.Literal)

		}
	}
}
