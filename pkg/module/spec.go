package module

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"reflect"
	"text/template"
)

var (
	ErrorNoTemplate     = errors.New("no template")
	ErrorNoTemplateFile = errors.New("no template file")
	ErrorNoStructName   = errors.New("no struct name")
)

type StructSpec struct {
	APIName string // snake_case
	Name    string // PascalCase
	VarName string // camelCase
	Package string // This is the Structs package

	PackageImportPath    string
	MiddlewareImportPath string
	PresenterImportPath  string
	HandlerImportPath    string
}

type TypeSpec struct {
	Backtick string

	TypeSpec *ast.StructType // internal TypeSpec
	Fields   []Field         // These are all the fields
	Ignored  bool
	Version  string
	APIPath  string

	PrimaryKeys     []Field // These are all the Primary Keys
	PrimaryKeyCount int

	Struct StructSpec

	Module           string // This is the application module.
	Package          string // This is the package
	IntermediaryPath string // This is the path between the Module and the Package Path

	Auth bool
}

type Field struct {
	Field      *ast.Field
	Rules      []Rule
	Name       string // PascalCase
	APIName    string // snake_case
	VarName    string // camelCase
	Type       string // We may want to change this to a token.Token
	PrimaryKey bool
	Ignore     bool              // Do we ignore this field
	Tag        reflect.StructTag // We want to work with the tags
}

// Rule
// I am not sure how to structure this yet
// The idea is to hold any rules
type Rule struct {
}

// GetAllFiles will generate and return all the required `File` for the TypeSpec
// This will generate the handler and the router,
// it may also generate specific middleware, auth, validation.
// This will not write the files
func (ts *TypeSpec) GetAllFiles(config *Config) ([]*File, error) {
	files := make([]*File, 0)
	ts.IntermediaryPath = config.Directories.IntermediaryPath
	ts.Module = config.Directories.Module
	ts.Version = config.Version
	ts.APIPath = config.APIPath
	//fmt.Printf("%+v", config)
	// Generate middleware
	if config.CreateMiddleware {
		ts.Package = "middleware"
		//TODO: We need to USE the config.FileNameConfig #FileNameConfig
		generate, err := ts.Generate(config.ValidationMiddlewareTemplate,
			config.getFullPath(
				config.OutputDir,
				config.Directories.MiddlewareDirectory,
			),
			ts.Struct.APIName+".middleware.validate.go",
		)
		if err != nil {
			return nil, err
		}
		files = append(files, generate)
	}

	// Generate routers
	if config.CreateRouter {
		ts.Package = "router"
		fmt.Println("")
		//TODO: We need to USE the config.FileNameConfig #FileNameConfig
		generate, err := ts.Generate(config.RouterTemplate,
			config.getFullPath(
				config.OutputDir,
				config.Directories.RouterDirectory,
			),
			ts.Struct.APIName+".router.go")
		if err != nil {
			return nil, err
		}
		files = append(files, generate)
	}

	// Generate handlers
	if config.CreateHandler {
		ts.Package = "handler"

		//TODO: We need to USE the config.FileNameConfig #FileNameConfig
		generate, err := ts.Generate(config.HandlerTemplate,
			config.getFullPath(
				config.OutputDir,
				config.Directories.HandlerDirectory,
			),
			ts.Struct.APIName+".handler.go")
		if err != nil {
			return nil, err
		}
		files = append(files, generate)
	}

	return files, nil
}

// Generate will generate the File for the TypeSpec using the template passed.
// This will generate both the handler and the router
// It may also generate specific middleware, auth, validation
func (ts *TypeSpec) Generate(templateText, path, name string) (*File, error) {
	if templateText == "" {
		return nil, ErrorNoTemplate
	}
	if ts.Struct.APIName == "" {
		return nil, ErrorNoStructName
	}
	parse, err := template.New(ts.Struct.APIName + "-template").Parse(templateText)
	if err != nil {
		fmt.Println("Error processing" + path + "  " + name)
		return nil, err
	}

	return ts.generate(parse, path, name)
}

// GenerateUsingFile will generate the File for the TypeSpec using the template passed.
// This will generate both the handler and the router
// It may also generate specific middleware, auth, validation
func (ts *TypeSpec) GenerateUsingFile(file, path, name string) (*File, error) {
	if file == "" {
		return nil, ErrorNoTemplateFile
	}

	parse, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}

	return ts.generate(parse, path, name)
}

// generate will generate a file using the template passed
func (ts *TypeSpec) generate(parse *template.Template, path, name string) (*File, error) {
	var err error
	// We do not want to template.Must as this will panic if the template is invalid
	//parse = template.Must(parse, err)

	// Create a buffer to hold the rendered output
	var buf bytes.Buffer

	// Execute the template and write the output to the buffer
	err = parse.Execute(&buf, *ts)
	if err != nil {
		return nil, err
	}

	return &File{
		Path:    path,
		Content: buf.String(),
		Name:    name,
	}, nil
}
