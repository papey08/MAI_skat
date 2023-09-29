package chich_ast

import (
	"strings"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
)

type Block []Node

func NewBlock() Block {
	return []Node{}
}

func (b *Block) String() string {
	var nodes []string
	for _, n := range *b {
		nodes = append(nodes, n.String())
	}
	return strings.Join(nodes, "; ")
}

func (b *Block) Add(n Node) {
	*b = append(*b, n)
}

func (b *Block) Compile(c *chich_compiler.Compiler) (p int, err error) {
	for _, n := range *b {
		if p, err = n.Compile(c); err != nil {
			return
		}

		if _, isReturn := n.(Return); !isReturn {
			p = c.Emit(chich_code.OpPop)
		}
	}
	return
}

func (b *Block) IsConstExpression() bool {
	return false
}
