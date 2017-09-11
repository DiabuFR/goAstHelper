package asth

import (
	"go/ast"
	"go/token"
)

type (
	Statement interface {
		asthStatement() ast.Stmt
	}

	Block struct {
		node *ast.BlockStmt
	}

	Assign struct {
		node *ast.AssignStmt
	}

	Return struct {
		node *ast.ReturnStmt
	}

	Call struct {
		node *ast.ExprStmt
	}
)

func (s *Block) asthStatement() ast.Stmt  { return s.node }
func (s *Assign) asthStatement() ast.Stmt { return s.node }
func (s *Return) asthStatement() ast.Stmt { return s.node }
func (s *Call) asthStatement() ast.Stmt   { return s.node }
func (s *Call) asthRValue() ast.Expr      { return s.node.X }

func NewBlock(stmts ...Statement) *Block {
	b := &Block{&ast.BlockStmt{List: []ast.Stmt{}}}
	for _, s := range stmts {
		b.node.List = append(b.node.List, s.asthStatement())
	}
	return b
}

func NewAssign(l Lvalue, r Rvalue) *Assign {
	return &Assign{
		node: &ast.AssignStmt{
			Lhs: []ast.Expr{l.asthLValue()},
			Rhs: []ast.Expr{r.asthRValue()},
			Tok: token.DEFINE,
		},
	}
}

func NewReturn(e Rvalue) *Return {
	return &Return{
		node: &ast.ReturnStmt{
			Results: []ast.Expr{e.asthRValue()},
		},
	}
}

func NewCall(receiver Rvalue, fctName string) *Call {
	var fct ast.Expr = ast.NewIdent(fctName)
	if receiver != nil {
		fct = &ast.SelectorExpr{
			X:   receiver.asthRValue(),
			Sel: ast.NewIdent(fctName),
		}
	}
	return &Call{
		node: &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun:  fct,
				Args: []ast.Expr{},
			},
		},
	}
}

func (c *Call) WithArgs(params ...Rvalue) *Call {
	expr := c.node.X.(*ast.CallExpr)
	for _, p := range params {
		expr.Args = append(expr.Args, p.asthRValue())
	}
	return c
}
