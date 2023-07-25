package modules

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"strings"
)

func GoDocReader(filename string, config Config) (error, []TypeSpec) {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err, nil
	}

	// Extract the package documentation
	pkg, err := doc.NewFromFiles(fset, []*ast.File{file}, "")

	// Extract the struct information
	var structs []*TypeSpec

	for _, typ := range pkg.Types {
		fmt.Println("Type:", typ.Name)
		fmt.Println("Doc:", strings.Replace(typ.Doc, "\n", "\\n", -1))

		if typ.Decl != nil && typ.Decl.Specs != nil {
			if spec, ok := typ.Decl.Specs[0].(*ast.TypeSpec); ok {
				if s, ok := spec.Type.(*ast.StructType); ok {
					//structs = append(structs, spec)
					fmt.Println("Struct:", spec.Name)
					fmt.Println("Comments:", spec.Comment.Text())
					fmt.Println("Fields:")
					for _, field := range s.Fields.List {
						fmt.Println("-",
							field.Names[0],
							field.Type,
							field.Tag.Value,
							strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
							strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
						)
					}
				}
			}
		}

	}

	fmt.Println("Package: ", pkg.Name)
	fmt.Println("Package Doc: ", pkg.Doc)

	for _, s := range structs {
		fmt.Println("Struct:", s.Name)
		fmt.Println("Doc:", s.Doc.Text())
		fmt.Println("Comments:", s.Comment.Text())
		fmt.Println("Fields:")
		for _, field := range s.Type.(*ast.StructType).Fields.List {
			fmt.Println("-",
				field.Names[0],
				field.Type,
				field.Tag.Value,
				strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
				strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
			)
		}
	}

	// Extract the package documentation
	return nil, nil
}

func extractStructInfo(filename string) error {

}
