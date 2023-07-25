package modules

import "go/ast"

type TypeSpec struct {
	TypeSpec *ast.TypeSpec
	Fields   []Field
	Ignored  bool
	APIPath  string
}

type Field struct {
	Field *ast.Field
	Rules Rule
	Type  string
}

type Rule struct {
	Field *ast.Field
}
