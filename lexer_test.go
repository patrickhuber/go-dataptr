package dataptr_test

import (
	"testing"

	"github.com/patrickhuber/go-dataptr"
)

func TestLexer(t *testing.T) {
	type test struct {
		name   string
		str    string
		tokens []dataptr.TokenType
	}
	tests := []test{
		{"name", "name", []dataptr.TokenType{dataptr.TokenName}},
		{"name and path", "name/parent/child", []dataptr.TokenType{
			dataptr.TokenName,
			dataptr.TokenSlash,
			dataptr.TokenName,
			dataptr.TokenSlash,
			dataptr.TokenName,
		}},
		{"condition", "name/key=value", []dataptr.TokenType{
			dataptr.TokenName,
			dataptr.TokenSlash,
			dataptr.TokenName,
			dataptr.TokenEqual,
			dataptr.TokenName,
		}},
		{"underscore", "name_with_underscore", []dataptr.TokenType{
			dataptr.TokenName,
		}},
		{"integer", "123", []dataptr.TokenType{
			dataptr.TokenInteger,
		}},
		{"integer with name", "123/name", []dataptr.TokenType{
			dataptr.TokenInteger,
			dataptr.TokenSlash,
			dataptr.TokenName,
		}},
		{"dash", "name/-", []dataptr.TokenType{
			dataptr.TokenName,
			dataptr.TokenSlash,
			dataptr.TokenDash,
		}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testLexer(t, test.str, test.tokens)
		})
	}
}

func testLexer(t *testing.T, str string, tokens []dataptr.TokenType) {
	lexer := dataptr.NewLexer(str)
	var i = 0
	for i = 0; i < len(tokens); i++ {
		actual, err := lexer.Next()

		if err != nil {
			t.Fatalf("expected err to be nil")
		}
		if actual == nil {
			t.Fatalf("expected actual result to not be nil")
		}
		expected := tokens[i]
		if expected != actual.Type {
			t.Fatalf("expected %s but found %s", expected, actual.Type)
		}
	}
	if len(tokens) != i {
		t.Fatalf("expected token count of %d but found %d", len(tokens), i)
	}
}
