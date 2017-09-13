package asth

import (
	"go/ast"
	"go/token"
)

type (
	If struct {
		node *ast.IfStmt
	}

	For struct {
		node *ast.ForStmt
	}

	Range struct {
		node *ast.RangeStmt
	}
)

func (s *If) asthStatement() ast.Stmt    { return s.node }
func (s *For) asthStatement() ast.Stmt   { return s.node }
func (s *Range) asthStatement() ast.Stmt { return s.node }

/// IF

func NewIf(cond Expr) *If {
	i := &If{
		node: &ast.IfStmt{},
	}
	if cond != nil {
		i.node.Cond = cond.asthExpr()
	}
	return i
}

func (s *If) WithInit(init Statement) *If {
	s.node.Init = init.asthStatement()
	return s
}

func (s *If) WithBody(stmts ...Statement) *If {
	s.node.Body = NewBlock(stmts...).node
	return s
}

func (s *If) WithBodyBlk(body *Block) *If {
	s.node.Body = body.node
	return s
}

func (s *If) WithElse(body *Block) *If {
	s.node.Else = body.node
	return s
}

func (s *If) WithElseIf(i *If) *If {
	s.node.Else = i.asthStatement()
	return s
}

/// FOR

func NewFor(cond Expr) *For {
	i := &For{
		node: &ast.ForStmt{},
	}
	if cond != nil {
		i.node.Cond = cond.asthExpr()
	}
	return i
}

func (s *For) WithInit(init Statement) *For {
	s.node.Init = init.asthStatement()
	return s
}

func (s *For) WithPost(post Statement) *For {
	s.node.Post = post.asthStatement()
	return s
}

func (s *For) WithBody(body *Block) *For {
	s.node.Body = body.node
	return s
}

/// RANGE

func NewForRange(iterator Rvalue) *Range {
	i := &Range{
		node: &ast.RangeStmt{
			X: iterator.asthRValue(),
		},
	}
	return i
}

func (s *Range) WithKey(k string, isDefinition bool) *Range {
	s.node.Key = ast.NewIdent(k)
	if isDefinition {
		s.node.Tok = token.DEFINE
	} else {
		s.node.Tok = token.ASSIGN
	}
	return s
}

func (s *Range) WithValue(v string) *Range {
	s.node.Value = ast.NewIdent(v)
	return s
}

func (s *Range) WithBody(body *Block) *Range {
	s.node.Body = body.node
	return s
}
