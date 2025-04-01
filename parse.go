package dataptr

import (
	"fmt"
	"strconv"
)

func Parse(str string) (DataPointer, error) {
	lexer := NewLexer(str)
	return parse(lexer)
}

func parse(lexer Lexer) (DataPointer, error) {
	var segments []Segment
	for {
		// can be a single segment
		segment, err := parseSegment(lexer)
		if err != nil {
			return DataPointer{}, err
		}
		segments = append(segments, segment)

		// or multiple
		ok, err := eat(lexer, TokenSlash)
		if err != nil {
			return DataPointer{}, err
		}
		if !ok {
			break
		}
	}
	return DataPointer{
		Segments: segments,
	}, nil
}

func parseSegment(lexer Lexer) (Segment, error) {
	tok, err := lexer.Peek()
	if err != nil {
		return nil, err
	}

	// we have an integer
	if tok.Type == TokenInteger {
		i, err := parseInteger(lexer)
		if err != nil {
			return nil, err
		}
		return Index{
			Index: i,
		}, nil
	}

	// otherwise this is a name
	name, err := parseName(lexer)
	if err != nil {
		return nil, err
	}

	// if no equal sign, this is an element
	ok, err := eat(lexer, TokenEqual)
	if err != nil {
		return nil, err
	}
	if !ok {
		return Element{
			Name: name,
		}, nil
	}

	// this is a constraint 'key=value'
	value, err := parseName(lexer)
	if err != nil {
		return nil, err
	}
	return Constraint{
		Key:   name,
		Value: value,
	}, nil
}

func parseInteger(lexer Lexer) (int, error) {
	tok, err := expect(lexer, TokenInteger)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(tok.Capture, 0, 0)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func parseName(lexer Lexer) (string, error) {
	tok, err := expect(lexer, TokenName)
	if err != nil {
		return "", err
	}
	return tok.Capture, nil
}

func eat(lexer Lexer, ty TokenType) (bool, error) {
	p, err := lexer.Peek()
	if err != nil {
		return false, err
	}
	if p.Type != ty {
		return false, nil
	}
	_, err = lexer.Next()
	if err != nil {
		return false, err
	}
	return true, nil
}

func expect(lexer Lexer, ty TokenType) (*Token, error) {
	tok, err := lexer.Next()
	if err != nil {
		return nil, err
	}
	if tok.Type != ty {
		return nil, parseError(tok, TokenName)
	}
	return tok, nil
}

func parseError(tok *Token, expected TokenType) error {
	return fmt.Errorf("error parsing at line: %d column: %d. Expected: %s found: %s", tok.Line, tok.Column, expected, tok.Type)
}
