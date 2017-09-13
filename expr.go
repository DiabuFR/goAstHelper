package asth

import (
	"go/ast"
	"go/token"
)

type (
	Operator token.Token

	Expr interface {
		asthExpr() ast.Expr
	}

	BinOp struct {
		node *ast.BinaryExpr
	}
)

const (
	ADD = Operator(token.ADD)
	SUB = Operator(token.SUB)
	MUL = Operator(token.MUL)
	QUO = Operator(token.QUO)
	REM = Operator(token.REM)

	AND     = Operator(token.AND)
	OR      = Operator(token.OR)
	XOR     = Operator(token.XOR)
	SHL     = Operator(token.SHL)
	SHR     = Operator(token.SHR)
	AND_NOT = Operator(token.AND_NOT)

	ADD_ASSIGN = Operator(token.ADD_ASSIGN)
	SUB_ASSIGN = Operator(token.SUB_ASSIGN)
	MUL_ASSIGN = Operator(token.MUL_ASSIGN)
	QUO_ASSIGN = Operator(token.QUO_ASSIGN)
	REM_ASSIGN = Operator(token.REM_ASSIGN)

	AND_ASSIGN     = Operator(token.AND_ASSIGN)
	OR_ASSIGN      = Operator(token.OR_ASSIGN)
	XOR_ASSIGN     = Operator(token.XOR_ASSIGN)
	SHL_ASSIGN     = Operator(token.SHL_ASSIGN)
	SHR_ASSIGN     = Operator(token.SHR_ASSIGN)
	AND_NOT_ASSIGN = Operator(token.AND_NOT_ASSIGN)

	LAND   = Operator(token.LAND)
	LOR    = Operator(token.LOR)
	ARROW  = Operator(token.ARROW)
	INC    = Operator(token.INC)
	DEC    = Operator(token.DEC)

	EQL    = Operator(token.EQL)
	LSS    = Operator(token.LSS)
	GTR    = Operator(token.GTR)
	ASSIGN = Operator(token.ASSIGN)
	NOT    = Operator(token.NOT)

	NEQ    = Operator(token.NEQ)
	LEQ    = Operator(token.LEQ)
	GEQ    = Operator(token.GEQ)
)

func (s *BinOp) asthExpr() ast.Expr { return s.node }

func NewBinOP(op Operator, op1 Rvalue, op2 Rvalue) *BinOp {
	return&BinOp{
		node: &ast.BinaryExpr{
			Op: token.Token(op),
			X: op1.asthRValue(),
			Y: op2.asthRValue(),
		},
	}
}
