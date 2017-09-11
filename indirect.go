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

func AddressOf(l Rvalue) Rvalue {
	if l == nil {
		return Nil
	}
	return &BaseRvalue{expr: &ast.UnaryExpr{
		Op: token.AND,
		X:  l.asthRValue(),
	}}
}

func SliceOfOrNil(typ Type, in ...Rvalue) Rvalue {
	if len(in) == 0 {
		return Nil
	}
	return SliceOf(typ, in...)
}

func SliceOf(eltType Type, in ...Rvalue) Rvalue {
	lit := &ast.CompositeLit{
		Type: SliceType(eltType).asthType(),
		Elts: []ast.Expr{},
	}

	for _, v := range in {
		lit.Elts = append(lit.Elts, v.asthRValue())
	}
	return &BaseRvalue{expr: lit}
}

func Dereference(l Rvalue) *Pointer {
	return &Pointer{expr: &ast.StarExpr{X: l.asthRValue()}}
}

func PointerType(l Type) Type {
	return &TypeRef{node: &ast.StarExpr{X: l.asthType()}}
}

func SliceType(l Type) Type {
	return &TypeRef{node: &ast.ArrayType{Elt: l.asthType()}}
}

func ArrayType(l *TypeRef, len int) *TypeRef {
	return &TypeRef{node: &ast.ArrayType{Elt: l.node, Len: NewIntLiteral(len).expr}}
}
