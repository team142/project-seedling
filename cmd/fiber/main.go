package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	input              = flag.String("i", "", "input struct(s) comma seperated. If this is blank all structs in this package will be processed")
	outputDir          = flag.String("o", "", "output directory default is .")
	generationFunction = func() {}
)

func main() {
	//flag.Parse()
	//if *input == "" {
	//	//if *input == "" || *outputDir == "" {
	//	flag.Usage()
	//	os.Exit(2)
	//}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "user.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			fmt.Println(t.Doc.Text())
		case *ast.StructType:
			for _, field := range t.Fields.List {
				fmt.Printf("%v\t%s\t%s\t%s\n", field.Names, field.Tag.Value, field.Comment.Text(), field.Doc.Text())
			}
		}
		return true
	})

	//fset := token.NewFileSet() // positions are relative to fset
	//
	//d, err := parser.ParseDir(fset, "./", nil, parser.ParseComments)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//for k, f := range d {
	//	fmt.Println("package", k)
	//	p := doc.New(f, "./", 0)
	//
	//	for _, t := range p.Types {
	//		fmt.Println("\ttype", t.Name)
	//		fmt.Println("\t\tdocs:", strings.Replace(t.Doc, "\n", "\\n", -1))
	//		for vI, variable := range t. {
	//			fmt.Printf("\t\t[%d]:%s\n", vI, variable.Decl.Doc)
	//		}
	//	}
	//}

	// 2. Inspect package and use type checker to infer imported types
	cfg := &packages.Config{
		Mode:  packages.NeedTypes | packages.NeedImports,
		Tests: false,
		ParseFile: func(fset *token.FileSet, parseFilename string, _ []byte) (*ast.File, error) {
			var src interface{}
			mode := parser.ParseComments // | parser.AllErrors
			file, err := parser.ParseFile(fset, parseFilename, src, mode)
			if file == nil {
				return nil, err
			}
			for _, decl := range file.Decls {
				if fd, ok := decl.(*ast.FuncDecl); ok {
					fd.Body = nil
				}
			}
			return file, nil
		},
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		failErr(fmt.Errorf("loading packages for inspection: %v", err))
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}
	for pkgIndex, pkg := range pkgs {
		fmt.Printf("Pakage index: %d\n", pkgIndex)
		fmt.Printf("ID: %s\n", pkg.ID)
		fmt.Printf("Name: %s\n", pkg.Name)
		fmt.Printf("GOPACKAGE: %s\n", os.Getenv("GOPACKAGE"))
		//fmt.Printf("%s\n", pkg.Types.Scope().String())

		fmt.Printf("TypesInfo: %+v\n", pkg.TypesInfo)
		//fmt.Printf("TypesInfo: %+v\n", pkg.Types.Scope())

		for i, scopeName := range pkg.Types.Scope().Names() {
			fmt.Printf("[%d] %s\n", i, scopeName)
			obj := pkg.Types.Scope().Lookup(scopeName)
			if obj == nil {
				failErr(fmt.Errorf("%s not found in declared types of %s",
					scopeName, pkg))
			}
			//fmt.Printf("\t%v\n", obj.Parent())
			//fmt.Printf("\t%v\n", obj.Pkg().)

			// 4. We check if it is a declared type
			if _, ok := obj.(*types.TypeName); !ok {
				failErr(fmt.Errorf("%v is not a named type", obj))
			}

			// 5. We expect the underlying type to be a struct
			structType, ok := obj.Type().Underlying().(*types.Struct)
			if !ok {
				failErr(fmt.Errorf("type %v is not a struct", obj))
			}

			//fmt.Printf("\t%s\n", structType.String())
			// 6. Now we can iterate through fields and access tags
			for sI := 0; sI < structType.NumFields(); sI++ {
				field := structType.Field(sI)
				tagValue := structType.Tag(sI)
				fmt.Printf("\t%s\t%s\t%s\n", field.Name(), field.Type(), tagValue)
			}
		}
	}

	// 3. Lookup the given source type name in the package declarations
	//obj := pkg.Types.Scope().Lookup(sourceTypeName)
	//if obj == nil {
	//	failErr(fmt.Errorf("%s not found in declared types of %s",
	//		sourceTypeName, pkg))
	//}
	//
	//// 4. We check if it is a declared type
	//if _, ok := obj.(*types.TypeName); !ok {
	//	failErr(fmt.Errorf("%v is not a named type", obj))
	//}
	//// 5. We expect the underlying type to be a struct
	//structType, ok := obj.Type().Underlying().(*types.Struct)
	//if !ok {
	//	failErr(fmt.Errorf("type %v is not a struct", obj))
	//}
	//
	//// 6. Now we can iterate through fields and access tags
	//for i := 0; i < structType.NumFields(); i++ {
	//	field := structType.Field(i)
	//	tagValue := structType.Tag(i)
	//	fmt.Println(field.Name(), tagValue, field.Type())
	//}
}

func loadPackage() *packages.Package {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		failErr(fmt.Errorf("loading packages for inspection: %v", err))
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	return pkgs[0]
}

func splitSourceType(sourceType string) (string, string) {
	idx := strings.LastIndexByte(sourceType, '.')
	if idx == -1 {
		failErr(fmt.Errorf(`expected qualified type as "pkg/path.MyType"`))
	}
	sourceTypePackage := sourceType[0:idx]
	sourceTypeName := sourceType[idx+1:]
	return sourceTypePackage, sourceTypeName
}

func failErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
