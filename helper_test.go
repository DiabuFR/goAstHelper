package asth

import (
	"fmt"
	"go/ast"
	"strings"
	"testing"
)

func Test_PathToSelectorExpr(t *testing.T) {
	input := []string{
		"hello",
		"fmt.Println",
		"struct.field.method",
	}

	for _, v := range input {
		expr := pathToSelectorExpr(strings.Split(v, ".")...)
		path := unwrapSelectorExpr(expr)

		if path != v {
			t.Fatalf("Wrong path: got `%s` but we expected `%s`", path, v)
		}
	}
}

func unwrapSelectorExpr(e ast.Expr) string {
	if e == nil {
		return "nil"
	}
	switch i := e.(type) {
	case *ast.Ident:
		return i.Name
	case *ast.SelectorExpr:
		switch j := i.X.(type) {
		case *ast.Ident:
			return j.Name
		case *ast.SelectorExpr:
			return unwrapSelectorExpr(i) + "." + j.Sel.Name
		default:
			return fmt.Sprintf("[INVALID_%s]", j)
		}
	}
	return fmt.Sprintf("[INVALID_%s]", e)
}
