package main

import (
	"code-gen/pkg/generator/v1"
	"code-gen/pkg/module/v1"
	"flag"
)

var (
	input             = flag.String("i", "", "input file")
	structs           = flag.String("s", "", "specify structs to generate (comma seperated), the default is every struct in the file")
	verbose           = flag.String("v", "", "verbose")
	templateFolder    = flag.String("t", "", "this is the template folder name")
	templateExtension = flag.String("extension", ".template", "this is the extension if templates are used")
	templateSingleton = flag.String("singleton", "singular", "if this word is in the name of a file, it will not be used while generating structs")
	api               = flag.String("api", "fiber", "what is the api framework you want to use")
	version           = flag.String("version", "", "version pass a version for the generated files, this will put the files into a version folder. It will also be used in the API version")
	outputDir         = flag.String("o", ".", "output directory default is \".\". This is used to control the generated files. Pass \"\" if you dont want files to be generated")
)

func main() {
	flag.Parse()

	fileName := "user.go"

	conf := &module.Config{
		Version:          "V1",
		Structs:          nil,
		Auth:             true,
		FileName:         fileName,
		DiscoverFunction: generator.TemplateGenerator,
		WriteToDisk:      true,
		OverrideFiles:    true,
	}

	if templateFolder != nil && *templateFolder != "" {
		conf.CreateFromTemplate = true
		conf.TemplateFolder = *templateFolder
		conf.TemplateSingleton = "singular"
		conf.TemplateExtension = ".template"
	}

	if templateSingleton != nil && *templateSingleton != "" {
		conf.TemplateSingleton = *templateSingleton
	}

	if templateExtension != nil && *templateExtension != "" {
		conf.TemplateExtension = *templateExtension
	}

	err := conf.Process()
	if err != nil {
		return
	}

}
