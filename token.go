package dataptr

type TokenType string

const (
	TokenInteger     TokenType = "integer"
	TokenName        TokenType = "name"
	TokenEqual       TokenType = "="
	TokenSlash       TokenType = "/"
	TokenEndOfStream TokenType = "EOF"
)

type Token struct {
	Type     TokenType
	Position int
	Column   int
	Line     int
	Capture  string
}
