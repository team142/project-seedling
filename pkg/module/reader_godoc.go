package module

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"reflect"
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
	// TODO: This should come from a config file
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
					ts := config.GetDefaultTypeSpec()

					ts.TypeSpec = s
					ts.Fields = make([]Field, 0)
					ts.Ignored = strings.Contains(strings.ToLower(typ.Doc), "#ignore")
					ts.APIPath = config.GetAPIPath(config.APIPath)
					ts.Module = pkg.ImportPath
					ts.PrimaryKeys = make([]Field, 0)
					ts.PrimaryKeyCount = 0
					ts.Struct = StructSpec{
						APIName:              "unknown",
						Name:                 "unknown",
						VarName:              "unknown",
						Package:              pkg.Name,
						PackageImportPath:    packageImportPath,
						MiddlewareImportPath: middlewareImportPath,
						PresenterImportPath:  presenterImportPath,
						HandlerImportPath:    handlerImportPath,
					}

					if spec.Name != nil {
						ts.Struct.APIName = toSnakeCase(spec.Name.String())
						ts.Struct.Name = toPascalCase(spec.Name.String())
						ts.Struct.VarName = toCamelCase(spec.Name.String())
					}

					//structs = append(structs, spec)
					fmt.Println("Struct:", spec.Name)
					if spec.Comment != nil {
						ts.Auth = commentContains(spec.Comment.Text(), "#noauth")
						fmt.Println("Comments:", spec.Comment.Text())
					}
					fmt.Println("Fields:")
					if s.Fields != nil {
						for _, field := range s.Fields.List {
							fn := "unknown"
							if len(field.Names) > 0 {
								fn = field.Names[0].Name
							}

							tempField := Field{
								Field: field,
								Name:  toPascalCase(fn),
								// TODO: read the tags
								APIName: toSnakeCase(fn),
								VarName: toCamelCase(fn),
							}

							if field.Doc != nil {
								tempField.PrimaryKey = commentContains(field.Doc.Text(), "#pk\n")
							}

							if field.Tag != nil {
								tempField.Tag = reflect.StructTag(strings.Replace(field.Tag.Value, "`", "", -1))
							}

							fmt.Println("-----------------------------", fn)
							fmt.Println(field.Doc.Text())
							fmt.Println("\t- Tag: ", tempField.Tag)

							if val, ok := tempField.Tag.Lookup("json"); ok {
								if val == "-" {
									tempField.Ignore = true
								}
								fmt.Println("\t\t- json: ", val)
							} else {
								fmt.Println("\t\t-", val, ok, tempField.Tag)
							}

							fmt.Printf("\t- Ignored: %v \n", tempField.Ignore)
							fmt.Printf("\t- %s \n", field.Type)
							fmt.Println("\t-",
								field.Names[0],
								field.Type,
								strings.Replace(field.Comment.Text(), "\n", "\\n", -1),
								strings.Replace(field.Doc.Text(), "\n", "\\n", -1),
							)
							if field.Tag != nil {
								fmt.Println("\t- TAG: ",
									field.Tag.Value,
								)
							}

							fmt.Println("-----------------------------")

							tempField.Type = fmt.Sprint(field.Type)

							if !tempField.Ignore {
								ts.Fields = append(ts.Fields, tempField)
							}

							if tempField.PrimaryKey {
								ts.PrimaryKeys = append(ts.PrimaryKeys, tempField)
								ts.PrimaryKeyCount++
							}
						}
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

func getComment(doc, lookingFor string) string {
	return ""
}

func commentContains(doc, lookingFor string) bool {
	return strings.Contains(strings.ToLower(doc), strings.ToLower(lookingFor))
}
