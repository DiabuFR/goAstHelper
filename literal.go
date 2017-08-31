package asth

import (
	"fmt"
	"go/ast"
	"go/token"
)

type (
	Literal interface {
		Rvalue
		asthLiteralExpr() ast.Expr
	}

	BasicLiteral struct {
		*BaseRvalue
	}

	StructLiteral struct {
		expr *ast.CompositeLit
	}
	MapLiteral struct {
		expr *ast.CompositeLit
	}
)

func NewStringLiteral(v string) *BasicLiteral {
	return &BasicLiteral{&BaseRvalue{&ast.BasicLit{Kind: token.STRING, Value: "\"" + v + "\""}}}
}
func NewIntLiteral(v int) *BasicLiteral {
	return &BasicLiteral{&BaseRvalue{&ast.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%d", v)}}}
}
func NewFloatLiteral(v float64) *BasicLiteral {
	return &BasicLiteral{&BaseRvalue{&ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprintf("%f", v)}}}
}
func (l *BasicLiteral) asthLiteralExpr() ast.Expr { return l.expr }

// MAP
func NewMapLiteral(keyType *TypeRef, valType *TypeRef) *MapLiteral {
	return &MapLiteral{
		&ast.CompositeLit{
			Type: &ast.MapType{Key: keyType.node, Value: valType.node},
		},
	}
}
func (l *MapLiteral) AddEntry(key Literal, val Literal) *MapLiteral {
	l.expr.Elts = append(l.expr.Elts, &ast.KeyValueExpr{
		Key:   key.asthLiteralExpr(),
		Value: val.asthLiteralExpr(),
	})
	return l
}
func (l *MapLiteral) AddEntries(es map[Literal]Literal) *MapLiteral {
	for k, v := range es {
		l.AddEntry(k, v)
	}
	return l
}

func (l *MapLiteral) asthLiteralExpr() ast.Expr { return l.expr }
func (l *MapLiteral) asthRValue() ast.Expr      { return l.expr }

// STRUCT

func NewStructLiteral() *StructLiteral {
	return &StructLiteral{&ast.CompositeLit{}}
}
func NewStructTypedLiteral(typ *TypeRef) *StructLiteral {
	return &StructLiteral{
		&ast.CompositeLit{
			Type: typ.node,
		},
	}
}
func (l *StructLiteral) AddField(f *StructFieldValue) *StructLiteral {
	l.expr.Elts = append(l.expr.Elts, f.expr)
	return l
}
func (l *StructLiteral) AddFields(fs ...*StructFieldValue) *StructLiteral {
	for _, f := range fs {
		l.AddField(f)
	}
	return l
}

func (l *StructLiteral) asthLiteralExpr() ast.Expr { return l.expr }
func (l *StructLiteral) asthRValue() ast.Expr      { return l.expr }

type StructFieldValue struct {
	expr ast.Expr
}

func NewStructFieldNamedValue(name string, val Rvalue) *StructFieldValue {
	return &StructFieldValue{
		&ast.KeyValueExpr{
			Key:   ast.NewIdent(name),
			Value: val.asthRValue(),
		},
	}
}

func NewStructFieldValue(lit Literal) *StructFieldValue {
	return &StructFieldValue{
		lit.asthLiteralExpr(),
	}
}
