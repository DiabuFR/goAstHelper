package asth

import (
	"fmt"
	"go/ast"
	"go/token"
)

type (
	Rvalue interface {
		asthRValue() ast.Expr
	}
	Lvalue interface {
		Rvalue
		asthLValue() ast.Expr
	}

	BaseRvalue struct {
		expr ast.Expr
	}
	BaseLvalue struct {
		expr ast.Expr
	}

	VarAssignSpec struct {
		spec *ast.ValueSpec
	}
)

func (v *BaseRvalue) asthRValue() ast.Expr { return v.expr }
func (v *BaseLvalue) asthRValue() ast.Expr { return v.expr }
func (v *BaseLvalue) asthLValue() ast.Expr { return v.expr }

func NewVarAssignSpec(name string, typ Type) *VarAssignSpec {
	if typ == nil {
		panic("Type param must be set in call to NewVarAssignSpec")
	}
	ident := ast.NewIdent(name)
	ident.Obj = &ast.Object{
		Name: name,
		Type: typ,
		Kind: ast.Var,
	}
	return &VarAssignSpec{&ast.ValueSpec{
		Names: []*ast.Ident{ident},
		Type:  typ.asthType(),
	}}
}

func (s *VarAssignSpec) WithValue(val Rvalue) *VarAssignSpec {
	s.spec.Values = []ast.Expr{val.asthRValue()}
	return s
}

// Helper functions

func NewVarDecl(specs ...*VarAssignSpec) *GenDecl {
	if len(specs) == 0 {
		fmt.Println("Empty NewVarDecl")
		return nil
	}
	pos := token.NoPos
	if len(specs) > 1 {
		pos = 1 // We just need something valid (!=0)
	}
	d := &GenDecl{
		&ast.GenDecl{
			Tok:    token.VAR,
			Specs:  []ast.Spec{},
			Lparen: pos,
			Rparen: pos,
		},
	}
	for _, s := range specs {
		d.node.Specs = append(d.node.Specs, s.spec)
	}
	return d
}

func NewConstDecl(specs ...*VarAssignSpec) *GenDecl {
	pos := token.NoPos
	if len(specs) > 1 {
		pos = 1 // We just need something valid (!=0)
	}
	d := &GenDecl{
		&ast.GenDecl{
			Tok:    token.CONST,
			Specs:  []ast.Spec{},
			Lparen: pos,
			Rparen: pos,
		},
	}
	for _, s := range specs {
		d.node.Specs = append(d.node.Specs, s.spec)
	}
	return d
}
