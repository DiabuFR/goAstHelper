package asth

import "go/ast"

func wrapWithQuotes(i string) string {
	return "\"" + i + "\""
}

func pathToSelectorExpr(elms ...string) ast.Expr {
	if len(elms) == 0 {
		return nil
	} else if len(elms) == 1 {
		return ast.NewIdent(elms[0])
	}

	return &ast.SelectorExpr{
		X:   pathToSelectorExpr(elms[0 : len(elms)-1]...),
		Sel: ast.NewIdent(elms[len(elms)-1]),
	}
}