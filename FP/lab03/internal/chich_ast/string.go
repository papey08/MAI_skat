package chich_ast

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"

	"chich/internal/chich_code"
	"chich/internal/chich_compiler"
	"chich/internal/chich_obj"
)

type String struct {
	s      string
	parse  parseFn
	substr []Node
	pos    int
}

func NewString(file, s string, parse parseFn, pos int) (Node, error) {
	str, err := escape(s)
	if err != nil {
		return nil, err
	}

	i := newInterpolator(file, str, parse)
	nodes, str, err := i.nodes()
	return String{s: str, parse: parse, substr: nodes, pos: pos}, err
}

func (s String) String() string {
	return s.s
}

func (s String) Quoted() string {
	return strconv.Quote(s.s)
}

func (s String) Compile(c *chich_compiler.Compiler) (position int, err error) {
	if len(s.substr) == 0 {
		position = c.Emit(chich_code.OpConstant, c.AddConstant(chich_obj.NewString(s.s)))
		c.Bookmark(s.pos)
		return
	}

	for _, sub := range s.substr {
		if position, err = sub.Compile(c); err != nil {
			return
		}
		c.RemoveLast()
	}

	position = c.Emit(chich_code.OpInterpolate, c.AddConstant(chich_obj.NewString(s.s)), len(s.substr))
	c.Bookmark(s.pos)
	return
}

func (s String) IsConstExpression() bool {
	return len(s.substr) == 0
}

func escape(s string) (string, error) {
	var buf strings.Builder

	for i := 0; i < len(s); {
		r, width := utf8.DecodeRuneInString(s[i:])
		buf.WriteRune(r)
		i += width
	}

	return buf.String(), nil
}

const eof = -1

var errBadInterpolationSyntax = errors.New("bad interpolation syntax")

type interpolator struct {
	s          string
	file       string
	parse      parseFn
	pos        int
	width      int
	nblocks    int
	inQuotes   bool
	inBacktick bool
	strings.Builder
}

func newInterpolator(file, s string, parse parseFn) interpolator {
	return interpolator{s: s, file: file, parse: parse}
}

func (i *interpolator) next() (r rune) {
	if i.pos >= len(i.s) {
		i.width = 0
		return eof
	}

	r, i.width = utf8.DecodeRuneInString(i.s[i.pos:])
	i.pos += i.width
	return
}

func (i *interpolator) backup() {
	i.pos -= i.width
}

func (i *interpolator) peek() rune {
	r := i.next()
	i.backup()
	return r
}

func (i *interpolator) enterBlock() {
	i.nblocks++
}

func (i *interpolator) exitBlock() {
	i.nblocks--
}

func (i *interpolator) insideBlock() bool {
	return i.nblocks > 0
}

func (i *interpolator) quotes() {
	i.inQuotes = !i.inQuotes
}

func (i *interpolator) backtick() {
	i.inBacktick = !i.inBacktick
}

func (i *interpolator) insideString() bool {
	return i.inQuotes || i.inBacktick
}

func (i *interpolator) acceptUntil(start, end rune) (string, error) {
	var buf strings.Builder

loop:
	for r := i.next(); ; r = i.next() {
		switch r {
		case eof:
			return "", errBadInterpolationSyntax

		case '"':
			i.quotes()

		case '`':
			i.backtick()

		case start:
			if !i.insideString() {
				i.enterBlock()
			}

		case end:
			if !i.insideString() {
				if !i.insideBlock() {
					break loop
				}
				i.exitBlock()
			}
		}

		buf.WriteRune(r)
	}

	return buf.String(), nil
}

func (i *interpolator) nodes() ([]Node, string, error) {
	var nodes []Node

	for r := i.next(); r != eof; r = i.next() {
		if r == '{' {
			if i.peek() == '{' {
				i.next()
				goto tail
			}

			s, err := i.acceptUntil('{', '}')
			if err != nil {
				return []Node{}, "", err
			} else if s == "" {
				continue
			}

			tree, errs := i.parse(i.file, s)
			if len(errs) > 0 {
				return []Node{}, "", i.parserError(errs)
			}

			nodes = append(nodes, tree)
			i.WriteString("%v")
			continue
		} else if r == '}' {
			if i.peek() != '}' {
				return []Node{}, "", errBadInterpolationSyntax
			}
			i.next()
		}

	tail:
		i.WriteRune(r)
	}

	return nodes, i.String(), nil
}

func (i *interpolator) parserError(errs []error) error {
	var buf strings.Builder

	buf.WriteString("interpolation errors:\n")
	for _, e := range errs {
		buf.WriteString(e.Error())
		buf.WriteByte('\n')
	}

	return errors.New(buf.String())
}
