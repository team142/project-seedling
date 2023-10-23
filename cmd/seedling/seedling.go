package main

import (
	"flag"
	"fmt"
	"github.com/team142/project-seedling/pkg/generator/v1"
	"github.com/team142/project-seedling/pkg/module"
	"io/ioutil"
	"strings"
)

var (
	//structs           = flag.String("s", "", "specify structs to generate (comma seperated), the default is every struct in the file")
	//verbose           = flag.String("v", "", "verbose")
	input             = flag.String("i", "", "input file")
	override          = flag.Bool("o", true, "do you want to override existing files")
	templateFolder    = flag.String("t", "", "this is the template folder name")
	templateExtension = flag.String("extension", ".template", "this is the extension if templates are used")
	templateSingleton = flag.String("singleton", "singular", "if this word is in the name of a file, it will not be used while generating structs")
	//api               = flag.String("api", "fiber", "what is the api framework you want to use")
	//version           = flag.String("version", "", "version pass a version for the generated files, this will put the files into a version folder. It will also be used in the API version")
	outputDir = flag.String("o", ".", "output directory default is \".\". This is used to control the generated files. Pass \"\" if you dont want files to be generated")
)

func main() {
	flag.Parse()

	//var fileName string
	if input != nil && *input != "" {
		err := executeForFile(*input)
		if err != nil {
			return
		}
	} else {
		files, err := ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".go") {
				err := executeForFile(file.Name())
				if err != nil {
					return
				}
			}
		}
	}

}

func executeForFile(fileName string) error {
	conf := &module.Config{
		Version:          "V1",
		FileName:         fileName,
		DiscoverFunction: generator.TemplateGenerator,
		WriteToDisk:      true,
		OverrideFiles:    true,
	}

	if override != nil {
		conf.OverrideFiles = *override
	}

	if templateFolder != nil && *templateFolder != "" {
		conf.CreateFromTemplate = true
		conf.TemplateFolder = *templateFolder
		conf.TemplateSingleton = "singular"
		conf.TemplateExtension = ".template"
	}

	if outputDir != nil && *outputDir != "" {
		conf.OutputDir = *outputDir
	} else {
		conf.OutputDir = "."
	}

	if templateSingleton != nil && *templateSingleton != "" {
		conf.TemplateSingleton = *templateSingleton
	}

	if templateExtension != nil && *templateExtension != "" {
		conf.TemplateExtension = *templateExtension
	}

	err := conf.Process()
	if err != nil {
		return err
	}
	return nil
}
