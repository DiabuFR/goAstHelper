package asth

import "go/ast"

type (
	Param struct {
		Name string
		Type Type
	}
)

func (param *Param) astField() *ast.Field {
	p := &ast.Field{
		Type: param.Type.asthType(),
	}
	if param.Name != "" {
		p.Names = []*ast.Ident{ast.NewIdent(param.Name)}
	}
	return p
}

func NewFuncDecl(name string) *FuncDecl {
	return &FuncDecl{&ast.FuncDecl{
		Name: ast.NewIdent(name),
		Type: &ast.FuncType{},
	}}
}

func (f *FuncDecl) WithReturnValues(vals ...Param) *FuncDecl {
	for _, v := range vals {
		p := v.astField()
		if f.node.Type.Results == nil {
			f.node.Type.Results = new(ast.FieldList)
		}
		f.node.Type.Results.List = append(f.node.Type.Results.List, p)
	}
	return f
}
func (f *FuncDecl) WithReceiver(r Param) *FuncDecl {
	f.node.Recv = &ast.FieldList{
		List: []*ast.Field{r.astField()},
	}
	return f
}

func (f *FuncDecl) WithParams(vals ...Param) *FuncDecl {
	for _, v := range vals {
		p := &ast.Field{
			Type: v.Type.asthType(),
		}
		if v.Name != "" {
			p.Names = []*ast.Ident{ast.NewIdent(v.Name)}
		}
		if f.node.Type.Params == nil {
			f.node.Type.Params = &ast.FieldList{
				List: []*ast.Field{},
			}
		}
		f.node.Type.Params.List = append(f.node.Type.Params.List, p)
	}
	return f
}

func (f *FuncDecl) WithBody(stmts ...Statement) *FuncDecl {
	f.node.Body = NewBlock(stmts...).node
	return f
}
func (f *FuncDecl) WithBodyBlk(b *Block) *FuncDecl {
	f.node.Body = b.node
	return f
}
