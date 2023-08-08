package module

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"strings"
)

func GoDocReader(config *Config) (error, []TypeSpec) {
	fmt.Println("---------GODOC---------")

	filename := config.FileName
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return err, nil
	}

	// We make sure the module information is setup
	config.SetModule()

	packageImportPath := config.Directories.Module + "/" + config.Directories.IntermediaryPath
	middlewareImportPath := config.getPathWithSeparator(packageImportPath, config.Directories.MiddlewareDirectory, "/")
	presenterImportPath := config.getPathWithSeparator(packageImportPath, config.Directories.PresenterDirectory, "/")
	handlerImportPath := config.getPathWithSeparator(packageImportPath, config.Directories.HandlerDirectory, "/")

	// Extract the package documentation
	pkg, err := doc.NewFromFiles(fset, []*ast.File{file}, packageImportPath)

	// Extract the struct information
	types := make([]TypeSpec, 0)

	fmt.Println("ImportPath:", pkg.ImportPath)
	fmt.Println("pkg Doc:", pkg.Doc)
	for _, typ := range pkg.Types {
		fmt.Println(FindPackagesForType(typ.Decl))

		fmt.Println("Type:", typ.Name)
		fmt.Println("Doc:", strings.Replace(typ.Doc, "\n", "\\n", -1))
		//switch tp := typ.Decl.Specs[0].(type) {
		//case *ast.Package:
		//}
		if typ.Decl != nil && typ.Decl.Specs != nil {
			//if spec, ok := typ.(*ast.Package); ok {
			//
			//}
			if spec, ok := typ.Decl.Specs[0].(*ast.TypeSpec); ok {
				if s, ok := spec.Type.(*ast.StructType); ok {
					ts := TypeSpec{
						TypeSpec:        s,
						Fields:          make([]Field, 0),
						Ignored:         strings.Contains(strings.ToLower(typ.Doc), "@ignore"),
						APIPath:         config.GetAPIPath(config.APIPath),
						Module:          pkg.ImportPath,
						PrimaryKeys:     make([]Field, 0),
						PrimaryKeyCount: 0,
						Struct: StructSpec{
							APIName:              toSnakeCase(spec.Name.String()),
							Name:                 toPascalCase(spec.Name.String()),
							VarName:              toCamelCase(spec.Name.String()),
							Package:              pkg.Name,
							PackageImportPath:    packageImportPath,
							MiddlewareImportPath: middlewareImportPath,
							PresenterImportPath:  presenterImportPath,
							HandlerImportPath:    handlerImportPath,
						},
					}

					//structs = append(structs, spec)
					fmt.Println("Struct:", spec.Name)
					fmt.Println("Comments:", spec.Comment.Text())
					fmt.Println("Fields:")
					for _, field := range s.Fields.List {
						fn := ""
						if len(field.Names) > 0 {
							fn = field.Names[0].Name
						}

						fmt.Println("-----------------------------")
						fmt.Println(field.Doc.Text(), strings.Contains(strings.ToLower(field.Doc.Text()), "#pk"))
						fmt.Println("-----------------------------")

						tempField := Field{
							Field: field,
							Name:  toPascalCase(fn),
							// TODO: read the tags
							APIName:    toSnakeCase(fn),
							VarName:    toCamelCase(fn),
							PrimaryKey: strings.Contains(strings.ToLower(field.Doc.Text()), "#pk"),
						}

						tempField.Type = fmt.Sprint(field.Type)

						ts.Fields = append(ts.Fields, tempField)

						if tempField.PrimaryKey {
							ts.PrimaryKeys = append(ts.PrimaryKeys, tempField)
							ts.PrimaryKeyCount++
						}

						fmt.Printf("- %s \n", field.Type)
						fmt.Println("\t-",
							field.Names[0],
							field.Type,
							field.Tag.Value,
							strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
							strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
						)

					}

					types = append(types, ts)
				}
			}
		}

	}

	fmt.Println("Package: ", pkg.Name)
	fmt.Println("Package Doc: ", pkg.Doc)
	fmt.Println("Package Imports: ", pkg.Imports)
	fmt.Println("Package ImportPath: ", pkg.ImportPath)
	fmt.Printf("%+v\n", pkg)

	//for _, s := range structs {
	//
	//	fmt.Println("Struct:", s.Name)
	//	fmt.Println("Doc:", s.Doc.Text())
	//	fmt.Println("Comments:", s.Comment.Text())
	//	fmt.Println("Fields:")
	//	for _, field := range s.Type.(*ast.StructType).Fields.List {
	//
	//		fmt.Println("-",
	//			field.Names[0],
	//			field.Type,
	//			field.Tag.Value,
	//			strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
	//			strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
	//		)
	//	}
	//}
	fmt.Println("---------GODOC---------")

	// Extract the package documentation
	return nil, types
}

func extractStructInfo(filename string) error {
	return nil
}
