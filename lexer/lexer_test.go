package lexer

import (
	"testing"

	"github.com/nlepage/monkey-interpreter/common"
	"github.com/nlepage/monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	expecteds := []token.Token{
		{token.ASSIGN, "=", common.Position{"test.monkey", 1, 1}},
		{token.PLUS, "+", common.Position{"test.monkey", 1, 2}},
		{token.LPAREN, "(", common.Position{"test.monkey", 1, 3}},
		{token.RPAREN, ")", common.Position{"test.monkey", 1, 4}},
		{token.LBRACE, "{", common.Position{"test.monkey", 1, 5}},
		{token.RBRACE, "}", common.Position{"test.monkey", 1, 6}},
		{token.COMMA, ",", common.Position{"test.monkey", 1, 7}},
		{token.SEMICOLON, ";", common.Position{"test.monkey", 1, 8}},
		{token.EOF, "", common.Position{"test.monkey", 1, 9}},
	}

	l, err := New("test.monkey")
	if err != nil {
		t.Fatal(err)
	}

	for _, expected := range expecteds {
		actual := l.NextToken()
		if actual != expected {
			t.Fatalf("Expected token %+v, got %+v", expected, actual)
		}
	}
}
