package asth

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type File struct {
	Name string

	node *ast.File

	importDecl *ast.GenDecl
	importDoc  *ast.CommentGroup

	objects map[string]*ast.Object
}

func (f *File) DefinedObject(name string) Lvalue {
	o, ok := f.objects[name]
	if !ok {
		return nil
	}

	return &BaseLvalue{expr: ast.NewIdent(o.Name)}
}

func (f *File) ListObjects() []string {
	objs := []string{}

	for k, v := range f.objects {
		typ := v.Type.(Type)
		objs = append(objs, fmt.Sprintf("%s: %s\t(Type: %s)", f.Name, k, typ.asthTypeName()))
	}
	return objs
}

func (f *File) defineObject(obj *ast.Object) {
	f.objects[obj.Name] = obj
}

func (f *File) WithImportDoc(lines ...string) *File {
	list := []*ast.Comment{}

	for _, l := range lines {
		l = "///" + l
		list = append(list, &ast.Comment{Text: l})
	}
	f.importDoc = &ast.CommentGroup{List: list}
	if f.importDecl != nil {
		f.importDecl.Doc = f.importDoc
	}
	return f
}

func (f *File) Get() *ast.File {
	return f.node
}

func (f *File) AddImport(imp *ImportSpec) {
	if f.importDecl == nil {
		f.importDecl = &ast.GenDecl{
			Tok:    token.IMPORT,
			Specs:  []ast.Spec{},
			Lparen: 1, // We just need something valid (!=0)
			Rparen: 1, // We just need something valid (!=0)
		}
		if f.importDoc != nil {
			f.importDecl.Doc = f.importDoc
		}
		f.node.Decls = append([]ast.Decl{f.importDecl}, f.node.Decls...)
	}
	f.importDecl.Specs = append(f.importDecl.Specs, imp.spec)
	f.node.Imports = append(f.node.Imports, imp.spec)
}
func (f *File) AddImports(imp ...*ImportSpec) {
	for _, i := range imp {
		f.AddImport(i)
	}
	ast.SortImports(token.NewFileSet(), f.node)
}

func (f *File) addDecl(decl Decl) {
	switch d := decl.(type) {
	case *TypeDecl:
		for _, t := range d.node.Specs {
			tspec, ok := t.(*ast.TypeSpec)
			if !ok {
				panic("TYPE GenDecl contains a non TypeValue spec.")
			}
			f.defineObject(tspec.Name.Obj)
		}
	case *GenDecl:
		switch d.node.Tok {
		case token.VAR | token.CONST:
			for _, v := range d.node.Specs {
				vspec, ok := v.(*ast.ValueSpec)
				if !ok {
					panic("VAR/CONST GenDecl contains a non ValueSpec spec.")
				}
				for _, name := range vspec.Names {
					if strings.Contains(name.Name, ".") {
						// Ignore name containing a .,  they're not defined in the file scope
						continue
					}
					f.defineObject(name.Obj)
				}
			}
		}
	}

	f.node.Decls = append(f.node.Decls, decl.asthDeclNode())
}
func (f *File) AddDecls(decls ...Decl) {
	for _, d := range decls {
		if d == nil {
			continue
		}
		f.addDecl(d)
	}
}
