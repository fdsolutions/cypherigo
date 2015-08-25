package cypherigo

import (
	"bytes"
	///"github.com/hashicorp/go-multierror"
)

type pttrn interface {
	Node(name string) Cypherizer
}

type node struct {
	name string
}

func (n *node) ToCypher() (string, error) {
	cypher := &bytes.Buffer{}
	cypher.WriteString(tokName(openParenthesis))
	cypher.WriteString(n.name)
	cypher.WriteString(tokName(closeParenthesis))
	return cypher.String(), nil
}

func Node(name string) Cypherizer {
	return &node{name}
}
