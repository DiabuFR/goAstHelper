package asth

import (
	"go/ast"
)

type (
	Decl interface {
		asthDeclNode() ast.Decl
	}

	FuncDecl struct {
		node *ast.FuncDecl
	}

	GenDecl struct {
		node *ast.GenDecl
	}
)

func NewFuncDecl(name string) *FuncDecl {
	return &FuncDecl{&ast.FuncDecl{Name: ast.NewIdent(name)}}
}

func (d *FuncDecl) asthDeclNode() ast.Decl { return d.node }
func (d *GenDecl) asthDeclNode() ast.Decl  { return d.node }
