package asth

import (
	"go/ast"
	"go/token"
	"strings"
)

var (
	StringType  = NewTypeRef("string")
	IntType     = NewTypeRef("int")
	UIntType    = NewTypeRef("uint")
	UInt64Type  = NewTypeRef("uint64")
	BoolType    = NewTypeRef("bool")
	Float32Type = NewTypeRef("float32")
	Float64Type = NewTypeRef("float64")
)

type (
	Type interface {
		asthType() ast.Expr
	}

	TypeRef struct {
		node ast.Expr
	}

	TypeSpec struct {
		node *ast.TypeSpec
	}

	MapType struct {
		node *ast.MapType
	}

	StructType struct {
		node *ast.StructType
	}

	StructField struct {
		Name string
		Type Type
	}

	InterfaceType struct {
		node *ast.InterfaceType
	}
)

func (t *TypeRef) asthType() ast.Expr       { return t.node }
func (t *MapType) asthType() ast.Expr       { return t.node }
func (t *StructType) asthType() ast.Expr    { return t.node }
func (t *InterfaceType) asthType() ast.Expr { return t.node }

func NewTypeRef(path string) *TypeRef {
	p := pathToSelectorExpr(strings.Split(path, ".")...)
	return &TypeRef{p}
}

func NewTypeDecl(specs ...TypeSpec) *TypeDecl {
	pos := token.NoPos
	if len(specs) > 1 {
		pos = 1 // We just need something valid (!=0)
	}
	d := &TypeDecl{
		node: &ast.GenDecl{
			Tok:    token.TYPE,
			Specs:  []ast.Spec{},
			Lparen: pos,
			Rparen: pos,
		},
	}
	for _, s := range specs {
		d.node.Specs = append(d.node.Specs, s.node)
	}
	return d
}

func NewMapType(key Type, val Type) *MapType {
	return &MapType{node: &ast.MapType{Key: key.asthType(), Value: val.asthType()}}
}

func NewTypeSpec(name string, typ Type) TypeSpec {
	return TypeSpec{
		node: &ast.TypeSpec{
			Name: ast.NewIdent(name),
			Type: typ.asthType(),
		},
	}
}

func NewStructType() *StructType {
	return &StructType{node: &ast.StructType{Fields: &ast.FieldList{List: []*ast.Field{}}}}
}

func (t *StructType) WithField(name string, typ Type) *StructType {
	t.node.Fields.List = append(t.node.Fields.List, &ast.Field{Names: []*ast.Ident{ast.NewIdent(name)}, Type: typ.asthType()})
	return t
}

func (t *StructType) WithTaggedField(name string, typ Type, tags string) *StructType {
	t.node.Fields.List = append(t.node.Fields.List, &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(name)},
		Type:  typ.asthType(),
		Tag:   NewBackquoteStringLiteral(tags).expr.(*ast.BasicLit),
	})
	return t
}

func NewInterfaceType() *InterfaceType {
	return &InterfaceType{
		node: &ast.InterfaceType{
			Methods: new(ast.FieldList),
		},
	}
}

func (t *InterfaceType) WithMethod(decl *FuncDecl) *InterfaceType {
	if decl == nil {
		return t
	}
	t.node.Methods.List = append(t.node.Methods.List, &ast.Field{
		Names: []*ast.Ident{decl.node.Name},
		Type:  decl.node.Type,
	})

	return t
}
