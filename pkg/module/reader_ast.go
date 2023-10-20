package module

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func AstReader(conf *Config) (error, []TypeSpec) {
	fmt.Println("---------AST---------")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, conf.FileName, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		//case *ast.Comment:
		//	fmt.Println("ast.Comment", t.Text)
		case *ast.TypeSpec:
			fmt.Println("ast.TypeSpec")
			fmt.Printf("\t%+v\n", t)
			fmt.Println("Struct:", t.Name)
			fmt.Println("Doc:", t.Doc.Text())
			fmt.Println("Comments:", t.Comment.Text())
			fmt.Println("Fields:")
			for _, field := range t.Type.(*ast.StructType).Fields.List {
				fmt.Println("-",
					field.Names[0],
					field.Type,
					field.Tag.Value,
					strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
					strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
				)
			}
		case *ast.Package:
			fmt.Println("ast.Package")
			fmt.Println(t.Name)
		//case *ast.GenDecl:
		//	fmt.Printf("ast.GenDecl %+v\n", t.Doc.Text())
		//case *ast.StructType:
		//	fmt.Println("ast.StructType")
		//	fmt.Printf("%+v\n", t)
		//	for _, field := range t.Fields.List {
		//		fmt.Printf("\t%v\t%s\t%s\t%s\n",
		//			field.Names,
		//			field.Tag.Value,
		//			strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
		//			strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
		//		)
		//	}
		default:
			if t != nil {
				fmt.Printf("---\t%+v\n", t)
			}
		}
		return true
	})
	fmt.Println("---------AST---------")

	return nil, nil
}
