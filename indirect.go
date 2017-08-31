package asth

import (
	"go/ast"
	"go/token"
)

type (
	Pointer struct {
		expr *ast.StarExpr
	}

	AddressOfExpr struct {
		expr *ast.UnaryExpr
	}
)

func AddressOf(l Rvalue) *BaseRvalue {
	return &BaseRvalue{expr: &ast.UnaryExpr{
		Op: token.AND,
		X:  l.asthRValue(),
	}}
}

func Dereference(l Rvalue) *Pointer {
	return &Pointer{expr: &ast.StarExpr{X: l.asthRValue()}}
}

func PointerType(l *TypeRef) *TypeRef {
	return &TypeRef{node: &ast.StarExpr{X: l.node}}
}
