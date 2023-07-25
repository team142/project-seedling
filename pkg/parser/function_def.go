package parser

import "go/ast"

type ParserFunction func(file string) (error, *ast.Package)
