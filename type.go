package asth

import (
	"go/ast"
	"strings"
)

type (
	TypeRef struct {
		node ast.Expr
	}
)

func NewTypeRef(path string) *TypeRef {
	p := pathToSelectorExpr(strings.Split(path, ".")...)
	return &TypeRef{p}
}
