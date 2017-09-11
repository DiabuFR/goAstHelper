package asth

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type (
	Package struct {
		name  *ast.Ident
		files map[string]*File
	}
)

func NewPackage(name string) *Package {
	return &Package{
		name:  ast.NewIdent(name),
		files: map[string]*File{},
	}
}

func (p *Package) NewFile(name string) *File {
	file := &File{
		Name: name,
		node: &ast.File{
			Name: p.name,
		},
		objects: map[string]*ast.Object{},
	}
	p.files[name] = file
	return file
}

func (p *Package) DefinedObject(name string) Lvalue {
	for _, f := range p.files {
		if o := f.DefinedObject(name); o != nil {
			return o
		}
	}
	return nil
}

func (p *Package) WriteFiles(outDir string, cfg printer.Config) error {
	fset := token.NewFileSet()
	files := p.files

	for k, v := range files {
		path := fmt.Sprintf("%s/%s", outDir, k)
		out, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error opening %s file: %s", k, err)
		}
		if err := cfg.Fprint(out, fset, v.Get()); err != nil {
			return fmt.Errorf("Error writing %s file: %s", k, err)
		}
	}
	return nil
}
