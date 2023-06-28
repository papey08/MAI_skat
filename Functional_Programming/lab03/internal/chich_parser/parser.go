package chich_parser

import (
	"chich/internal/chich_ast/calc_ops"
	"chich/internal/chich_ast/logic_ops"
	"errors"
	"strconv"

	"chich/internal/chich_ast"
	"chich/internal/chich_err"
	"chich/internal/chich_item"
	"chich/internal/chich_lexer"
)

type Parser struct {
	items         chan chich_item.Item
	file          string
	input         string
	prefixParsers map[chich_item.Type]parsePrefixFn
	infixParsers  map[chich_item.Type]parseInfixFn
	cur           chich_item.Item
	peek          chich_item.Item
	errs          []error
	nestedLoops   uint
}

type (
	parsePrefixFn func() chich_ast.Node
	parseInfixFn  func(chich_ast.Node) chich_ast.Node
)

const (
	Lowest int = iota
	Assignment
	LogicalOr
	LogicalAnd
	Equality
	Relational
	Additive
	Multiplicative
	Call
	Index
)

var precedences = map[chich_item.Type]int{
	chich_item.Assign:        Assignment,
	chich_item.ModulusAssign: Assignment,
	chich_item.Or:            LogicalOr,
	chich_item.And:           LogicalAnd,
	chich_item.Equals:        Equality,
	chich_item.NotEquals:     Equality,
	chich_item.LT:            Relational,
	chich_item.GT:            Relational,
	chich_item.LTEQ:          Relational,
	chich_item.GTEQ:          Relational,
	chich_item.Plus:          Additive,
	chich_item.Minus:         Additive,
	chich_item.Modulus:       Multiplicative,
	chich_item.Slash:         Multiplicative,
	chich_item.Asterisk:      Multiplicative,
	chich_item.LParen:        Call,
	chich_item.LBracket:      Index,
}

func newParser(file, input string, items chan chich_item.Item) *Parser {
	p := &Parser{
		cur:           <-items,
		peek:          <-items,
		items:         items,
		file:          file,
		input:         input,
		prefixParsers: make(map[chich_item.Type]parsePrefixFn),
		infixParsers:  make(map[chich_item.Type]parseInfixFn),
	}
	p.registerPrefix(chich_item.Ident, p.parseIdentifier)
	p.registerPrefix(chich_item.Int, p.parseInteger)
	p.registerPrefix(chich_item.String, p.parseString)
	p.registerPrefix(chich_item.LParen, p.parseGroupedExpr)
	p.registerPrefix(chich_item.If, p.parseIfExpr)
	p.registerPrefix(chich_item.Function, p.parseFunction)
	p.registerPrefix(chich_item.LBracket, p.parseList)
	p.registerPrefix(chich_item.Error, p.parseError)

	p.registerInfix(chich_item.Equals, p.parseEquals)
	p.registerInfix(chich_item.NotEquals, p.parseNotEquals)
	p.registerInfix(chich_item.LT, p.parseLess)
	p.registerInfix(chich_item.GT, p.parseGreater)
	p.registerInfix(chich_item.LTEQ, p.parseLessEq)
	p.registerInfix(chich_item.GTEQ, p.parseGreaterEq)
	p.registerInfix(chich_item.And, p.parseAnd)
	p.registerInfix(chich_item.Or, p.parseOr)
	p.registerInfix(chich_item.Plus, p.parsePlus)
	p.registerInfix(chich_item.Minus, p.parseMinus)
	p.registerInfix(chich_item.Slash, p.parseSlash)
	p.registerInfix(chich_item.Asterisk, p.parseAsterisk)
	p.registerInfix(chich_item.Modulus, p.parseModulus)
	p.registerInfix(chich_item.Assign, p.parseAssign)
	p.registerInfix(chich_item.LParen, p.parseCall)

	return p
}

func (p *Parser) next() {
	p.cur = p.peek
	p.peek = <-p.items
}

func (p *Parser) errors() []error {
	return p.errs
}

func (p *Parser) errorf(s string, a ...any) {
	p.errs = append(p.errs, chich_err.New(p.file, p.input, p.cur.Pos, s, a...))
}

func (p *Parser) parse() chich_ast.Node {
	var block = chich_ast.NewBlock()

	for !p.cur.Is(chich_item.EOF) {
		if s := p.parseStatement(); s != nil {
			block.Add(s)
		}
		p.next()
	}
	return &block
}

func (p *Parser) parseStatement() chich_ast.Node {
	if p.cur.Is(chich_item.Return) {
		return p.parseReturn()
	}
	return p.parseExpr(Lowest)
}

func (p *Parser) parseReturn() chich_ast.Node {
	var ret chich_ast.Node

	p.next()
	if !p.cur.Is(chich_item.Semicolon) {
		ret = chich_ast.NewReturn(p.parseExpr(Lowest), p.cur.Pos)
	}

	if p.peek.Is(chich_item.Semicolon) {
		p.next()
	}
	return ret
}

func (p *Parser) parseExpr(precedence int) chich_ast.Node {
	if prefixFn, ok := p.prefixParsers[p.cur.Typ]; ok {
		leftExp := prefixFn()

		for !p.peek.Is(chich_item.Semicolon) && precedence < p.peekPrecedence() {
			if infixFn, ok := p.infixParsers[p.peek.Typ]; ok {
				p.next()
				leftExp = infixFn(leftExp)
			} else {
				break
			}
		}

		if p.peek.Is(chich_item.Semicolon) {
			p.next()
		}
		return leftExp
	}
	p.noParsePrefixFnError(p.cur.Typ)
	return nil
}

func (p *Parser) parseGroupedExpr() chich_ast.Node {
	p.next()
	exp := p.parseExpr(Lowest)
	if !p.expectPeek(chich_item.RParen) {
		return nil
	}
	return exp
}

func (p *Parser) parseBlock() chich_ast.Node {
	var block chich_ast.Block
	p.next()

	for !p.cur.Is(chich_item.RBrace) && !p.cur.Is(chich_item.EOF) {
		if s := p.parseStatement(); s != nil {
			block.Add(s)
		}
		p.next()
	}

	if !p.cur.Is(chich_item.RBrace) {
		p.peekError(chich_item.RBrace)
		return nil
	}

	return &block
}

func (p *Parser) parseIfExpr() chich_ast.Node {
	pos := p.cur.Pos
	p.next()
	cond := p.parseExpr(Lowest)

	if !p.expectPeek(chich_item.LBrace) {
		return nil
	}

	body := p.parseBlock()

	var alt chich_ast.Node
	if p.peek.Is(chich_item.Else) {
		p.next()

		if p.peek.Is(chich_item.If) {
			p.next()
			alt = p.parseIfExpr()
		} else {
			if !p.expectPeek(chich_item.LBrace) {
				return nil
			}
			alt = p.parseBlock()
		}
	}

	return chich_ast.NewIfExpr(cond, body, alt, pos)
}

func (p *Parser) parseList() chich_ast.Node {
	nodes := p.parseNodeList(chich_item.RBracket)
	return chich_ast.NewList(nodes...)
}

func (p *Parser) parseFunction() chich_ast.Node {
	pos := p.cur.Pos
	if !p.expectPeek(chich_item.LParen) {
		return nil
	}

	params := p.parseFunctionParams()
	if !p.expectPeek(chich_item.LBrace) {
		return nil
	}

	return chich_ast.NewFunction(params, p.parseBlock(), pos)
}

func (p *Parser) parseFunctionParams() []chich_ast.Identifier {
	var ret []chich_ast.Identifier

	if p.peek.Is(chich_item.RParen) {
		p.next()
		return ret
	}

	p.next()
	ret = append(ret, chich_ast.NewIdentifier(p.cur.Val, p.cur.Pos))

	for p.peek.Is(chich_item.Comma) {
		p.next()
		p.next()
		ret = append(ret, chich_ast.NewIdentifier(p.cur.Val, p.cur.Pos))
	}

	if !p.expectPeek(chich_item.RParen) {
		return nil
	}
	return ret
}

func (p *Parser) parseIdentifier() chich_ast.Node {
	return chich_ast.NewIdentifier(p.cur.Val, p.cur.Pos)
}

func (p *Parser) parseError() chich_ast.Node {
	p.errs = append(p.errs, errors.New(p.cur.Val))
	return nil
}

func (p *Parser) parseInteger() chich_ast.Node {
	i, err := strconv.ParseInt(p.cur.Val, 0, 64)
	if err != nil {
		p.errorf("unable to parse %q as integer", p.cur.Val)
		return nil
	}
	return chich_ast.NewInteger(i)
}

func (p *Parser) parseString() chich_ast.Node {
	s, err := chich_ast.NewString(p.file, p.cur.Val, Parse, p.cur.Pos)
	if err != nil {
		p.errorf(err.Error())
		return nil
	}
	return s
}

func (p *Parser) parsePlus(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return calc_ops.NewPlus(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseMinus(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return calc_ops.NewMinus(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseAsterisk(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return calc_ops.NewTimes(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseSlash(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return calc_ops.NewDivide(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseModulus(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return calc_ops.NewMod(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseEquals(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewEquals(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseNotEquals(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewNotEquals(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseLess(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewLess(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseGreater(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewGreater(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseLessEq(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewLessEq(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseGreaterEq(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewGreaterEq(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseAnd(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewAnd(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseOr(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	prec := p.precedence()
	p.next()
	return logic_ops.NewOr(left, p.parseExpr(prec), pos)
}

func (p *Parser) parseAssign(left chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	p.next()
	right := p.parseExpr(Lowest)

	i, leftIsIdentifier := left.(chich_ast.Identifier)
	fn, rightIsFunction := right.(chich_ast.Function)

	if leftIsIdentifier && rightIsFunction {
		fn.Name = i.String()
	}

	return chich_ast.NewAssign(left, right, pos)
}

func (p *Parser) parseCall(fn chich_ast.Node) chich_ast.Node {
	pos := p.cur.Pos
	return chich_ast.NewCall(fn, p.parseNodeList(chich_item.RParen), pos)
}

func (p *Parser) parsePair() [2]chich_ast.Node {
	l := p.parseExpr(Lowest)
	p.next()
	r := p.parseExpr(Lowest)

	return [2]chich_ast.Node{l, r}
}

func (p *Parser) parseNodeList(end chich_item.Type) []chich_ast.Node {
	return p.parseNodeSequence(chich_item.Comma, end)
}

func (p *Parser) parseNodeSequence(sep, end chich_item.Type) []chich_ast.Node {
	var seq []chich_ast.Node

	p.next()
	if p.cur.Is(end) {
		return seq
	}

	seq = append(seq, p.parseExpr(Lowest))

	for p.peek.Is(sep) {
		p.next()
		p.next()
		seq = append(seq, p.parseExpr(Lowest))
	}

	if !p.expectPeek(end) {
		return nil
	}
	return seq
}

func (p *Parser) expectPeek(t chich_item.Type) bool {
	if p.peek.Is(t) {
		p.next()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t chich_item.Type) {
	p.errorf("expected next chich_item to be %v, got %v instead", t, p.peek.Typ)
}

func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peek.Typ]; ok {
		return prec
	}
	return Lowest
}

func (p *Parser) precedence() int {
	if prec, ok := precedences[p.cur.Typ]; ok {
		return prec
	}
	return Lowest
}

func (p *Parser) registerPrefix(typ chich_item.Type, fn parsePrefixFn) {
	p.prefixParsers[typ] = fn
}

func (p *Parser) registerInfix(typ chich_item.Type, fn parseInfixFn) {
	p.infixParsers[typ] = fn
}

func (p *Parser) noParsePrefixFnError(t chich_item.Type) {
	p.errorf("no parse prefix function for %q found", t)
}

func Parse(file, input string) (prog chich_ast.Node, errs []error) {
	items := chich_lexer.Lex(input)
	p := newParser(file, input, items)
	return p.parse(), p.errors()
}
