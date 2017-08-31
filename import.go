package asth

import (
	"go/ast"
	"go/token"
)

type ImportSpec struct {
	spec *ast.ImportSpec
}

func NewImportSpec(path string) *ImportSpec {
	return &ImportSpec{&ast.ImportSpec{
		Path: &ast.BasicLit{Kind: token.STRING, Value: wrapWithQuotes(path)},
	}}
}

func (i *ImportSpec) WithName(name string) {
	i.spec.Name = ast.NewIdent(name)
}
