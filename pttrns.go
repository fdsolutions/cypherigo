package cypherigo

import (
	"bytes"
	"fmt"
	"github.com/lann/builder"
	"strings"
)

type prop map[string]interface{}

type nodeData struct {
	Ident      string
	WithLabels []string
	WithProps  []prop
}

// NodeBuilder builds Cypher Node
type nodeBuilder builder.Builder

var Node = builder.Register(nodeBuilder{}, nodeData{}).(nodeBuilder)

func (d nodeData) ToCypher() (string, error) {
	fmt.Printf("%#v", d)
	lbls := d.WithLabels
	l := len(d.WithLabels)
	symbols := make([]string, l)
	for _, lbl := range lbls {
		sym := fmt.Sprintf(":%s", bytes.Title([]byte(lbl)))
		symbols = append(symbols, sym)
	}
	return fmt.Sprintf("(%s%s)", d.Ident, strings.Join(symbols, "")), nil
}

func (b nodeBuilder) Ident(name string) nodeBuilder {
	return builder.Set(b, "Ident", name).(nodeBuilder)
}

func (b nodeBuilder) WithLabels(lbls ...string) nodeBuilder {
	return builder.Set(b, "WithLabels", lbls).(nodeBuilder)
}

func (b nodeBuilder) ToCypher() (string, error) {
	d := builder.GetStruct(b).(nodeData)
	return d.ToCypher()
}
