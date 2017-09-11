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

	TypeDecl struct {
		node *ast.GenDecl
	}
)
func (d *FuncDecl) asthDeclNode() ast.Decl { return d.node }
func (d *GenDecl) asthDeclNode() ast.Decl  { return d.node }
func (d *TypeDecl) asthDeclNode() ast.Decl { return d.node }
