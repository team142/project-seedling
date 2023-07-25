package generator

import "go/ast"

type GeneratorFunction func(pkg *ast.Package, templates map[string]string) error
