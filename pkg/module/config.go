package module

import (
	"code-gen/pkg/generator"
	"code-gen/pkg/parser"
)

//type Config interface {
//}

type Config struct {
	InputFile string
	Verbose   string
	API       string
	Version   string
	OutputDir string
	Parser    string
	Generator string
	Structs   []string

	// ParserFunc is the function which is used to parse the go file
	ParserFunc parser.ParserFunction

	// GeneratorFunc is the function which will generate all the files and fields
	GeneratorFunc generator.GeneratorFunction
}

func (c *Config) Generate() {

}

func (c *Config) Set() {

}
