package lexer

import (
	"bufio"
	"io"
	"os"
	"unicode/utf8"

	"github.com/nlepage/monkey-interpreter/common"
	"github.com/nlepage/monkey-interpreter/token"
)

// Lexer is a instance of a lexer used to lex
type Lexer struct {
	pos        common.Position
	in         *bufio.Reader
	line       string
	eof        bool
	ch         rune
	readColumn int
}

// New creates a new Lexer for a given file
func New(filename string) (*Lexer, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	l := &Lexer{
		common.Position{
			filename,
			0,
			0,
		},
		bufio.NewReader(f),
		"",
		false,
		0,
		1,
	}

	l.readCh()

	return l, nil
}

// NextToken gives the next token and advances the lexers position
func (l *Lexer) NextToken() token.Token {
	t := token.Token{0, string(l.ch), l.pos}

	switch l.ch {
	case '=':
		t.Type = token.ASSIGN
	case '+':
		t.Type = token.PLUS
	case '(':
		t.Type = token.LPAREN
	case ')':
		t.Type = token.RPAREN
	case '{':
		t.Type = token.LBRACE
	case '}':
		t.Type = token.RBRACE
	case ',':
		t.Type = token.COMMA
	case ';':
		t.Type = token.SEMICOLON
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	default:
		t.Type = token.ILLEGAL
	}

	l.readCh()

	return t
}

func (l *Lexer) readCh() {
	readIndex := l.readColumn - 1

	for readIndex >= len(l.line) && !l.eof {
		var err error
		l.line, err = l.in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				l.eof = true
			} else {
				panic(err)
			}
		}

		l.pos.Line++
		l.pos.Column = 0
		l.readColumn = 1
		readIndex = 0
	}

	var size int
	if readIndex >= len(l.line) && l.eof {
		l.ch = 0
	} else {
		l.ch, size = utf8.DecodeRuneInString(l.line[readIndex:])
	}

	l.pos.Column = l.readColumn
	l.readColumn += size
}
