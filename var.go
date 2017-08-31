package asth

import (
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

func NewVarAssignSpec(name string, typ *TypeRef) *VarAssignSpec {
	ident := ast.NewIdent(name)
	ident.Obj = &ast.Object{
		Name: name,
		Type: typ.node,
		Kind: ast.Var,
	}
	return &VarAssignSpec{&ast.ValueSpec{
		Names: []*ast.Ident{ident},
		Type:  typ.node,
	}}
}

func (s *VarAssignSpec) WithValue(val Rvalue) *VarAssignSpec {
	s.spec.Values = []ast.Expr{val.asthRValue()}
	return s
}

// Helper functions

func NewVarDecl(spec *VarAssignSpec) *GenDecl {
	return &GenDecl{
		&ast.GenDecl{
			Tok:   token.VAR,
			Specs: []ast.Spec{spec.spec},
		},
	}
}
func NewVarGroupDecl(specs ...*VarAssignSpec) *GenDecl {
	d := &GenDecl{
		&ast.GenDecl{
			Tok:    token.VAR,
			Specs:  []ast.Spec{},
			Lparen: 1, // We just need something valid (!=0)
			Rparen: 1, // We just need something valid (!=0)
		},
	}
	for _, s := range specs {
		d.node.Specs = append(d.node.Specs, s.spec)
	}
	return d
}
