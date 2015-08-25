package cypherigo

// Cypherizer is an interface that wraps the ToCypher method.
// Every type needs to implement this method to generate a cypher string.
//
// ToCypher returns the Cypher representation of the cypherizer and an
// error if something bad happens.
type Cypherizer interface {
	ToCypher() (string, error)
}

// Tokens
//
type token int

const (
	comma token = iota
	openParenthesis
	closeParenthesis
	hyphen
	openBracket
	closeBracket
	openBrace
	closeBrace
	space
	matchClause
)

// tokenNames
var tokenNames = [...]string{
	comma:            ",",
	openParenthesis:  "(",
	closeParenthesis: ")",
	hyphen:           "-",
	openBracket:      "[",
	closeBracket:     "]",
	openBrace:        "{",
	closeBrace:       "}",
	space:            " ",
	// Clauses
	matchClause: "MATCH",
}

func tokName(t token) string {
	return tokenNames[t]
}
