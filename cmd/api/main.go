package main

import (
	"code-gen/pkg/generator"
	"code-gen/pkg/module"
	"flag"
)

var (
	input     = flag.String("i", "", "input file")
	structs   = flag.String("s", "", "specify structs to generate (comma seperated), the default is every struct in the file")
	verbose   = flag.String("v", "", "verbose")
	api       = flag.String("api", "fiber", "what is the api framework you want to use")
	version   = flag.String("version", "", "version pass a version for the generated files, this will put the files into a version folder. It will also be used in the API version")
	outputDir = flag.String("o", ".", "output directory default is \".\". This is used to control the generated files. Pass \"\" if you dont want files to be generated")
)

func main() {
	flag.Parse()

	fileName := "user.go"

	conf := &module.Config{
		Version:          "V1",
		Structs:          nil,
		Auth:             true,
		FileName:         fileName,
		DiscoverFunction: generator.FiberGenerator,
		CreateHandler:    true,
		CreateRouter:     true,
		CreateMiddleware: true,
		WriteToDisk:      true,
		OverrideFiles:    true,
	}

	err := conf.Process()
	if err != nil {
		return
	}

	//err = extractStructInfo("user.go")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Print the struct information

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
	//cfg := &packages.Config{
	//	Mode:  packages.NeedTypes | packages.NeedImports,
	//	Tests: false,
	//	ParseFile: func(fset *token.FileSet, parseFilename string, _ []byte) (*ast.File, error) {
	//		var src interface{}
	//		mode := parser.ParseComments // | parser.AllErrors
	//		file, err := parser.ParseFile(fset, parseFilename, src, mode)
	//		if file == nil {
	//			return nil, err
	//		}
	//		for _, decl := range file.Decls {
	//			if fd, ok := decl.(*ast.FuncDecl); ok {
	//				fd.Body = nil
	//			}
	//		}
	//		return file, nil
	//	},
	//}
	//
	//pkgs, err := packages.Load(cfg, ".")
	//if err != nil {
	//	failErr(fmt.Errorf("loading packages for inspection: %v", err))
	//}
	//if packages.PrintErrors(pkgs) > 0 {
	//	os.Exit(1)
	//}
	//for pkgIndex, pkg := range pkgs {
	//	if *verbose == "v" {
	//		fmt.Printf("Pakage index: %d\n", pkgIndex)
	//		fmt.Printf("ID: %s\n", pkg.ID)
	//		fmt.Printf("Name: %s\n", pkg.Name)
	//		fmt.Printf("GOPACKAGE: %s\n", os.Getenv("GOPACKAGE"))
	//		fmt.Printf("TypesInfo: %+v\n", pkg.TypesInfo)
	//	}
	//
	//	for i, scopeName := range pkg.Types.Scope().Names() {
	//		fmt.Printf("[%d] %s\n", i, scopeName)
	//		obj := pkg.Types.Scope().Lookup(scopeName)
	//		if obj == nil {
	//			failErr(fmt.Errorf("%s not found in declared types of %s",
	//				scopeName, pkg))
	//		}
	//
	//		// We check if it is a declared type
	//		if _, ok := obj.(*types.TypeName); !ok {
	//			failErr(fmt.Errorf("%v is not a named type", obj))
	//		}
	//
	//		// We expect the underlying type to be a struct
	//		structType, ok := obj.Type().Underlying().(*types.Struct)
	//		if !ok {
	//			failErr(fmt.Errorf("type %v is not a struct", obj))
	//		}
	//
	//		for sI := 0; sI < structType.NumFields(); sI++ {
	//			field := structType.Field(sI)
	//			tagValue := structType.Tag(sI)
	//			if *verbose == "v" {
	//				fmt.Printf("\t%s\t%s\t%s\n", field.Name(), field.Type(), tagValue)
	//			}
	//		}
	//	}
	//}
}
