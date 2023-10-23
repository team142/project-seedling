package module

import (
	"fmt"
	"github.com/team142/project-seedling/pkg/module/templates"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//type Config interface {
//}

type DirectoryName int

type File struct {
	Path    string
	Content string
	Name    string
}

type DirectoryConfig struct {
	DirectoryHasVersion bool

	RouterDirectory     string
	HandlerDirectory    string
	MiddlewareDirectory string
	PresenterDirectory  string

	Module           string
	IntermediaryPath string
}

//TODO: We need to implement this #FileNameConfig
//type FileNameConfig struct {
//	RouterNameFormat     string
//	HandlerNameFormat    string
//	MiddlewareNameFormat string
//}

type Config struct {
	Version     string   // This is the version we want to use for the generator, this will version the API
	Structs     []string // This is the structs we want to process
	OutputDir   string   // This is the output directory
	Directories DirectoryConfig
	//FileNameConfig               FileNameConfig
	OutputDirPermissions         os.FileMode       // This is the permissions we will use for the files we generate
	APIPath                      string            // This is the API Path. The prefix
	Separator                    rune              // This is the file separator
	Auth                         bool              // Are we going to do auth
	FileName                     string            // This is the filename we want to process
	DiscoverFunction             DiscoveryFunction // This is the function which will be called to get the information from the structs
	TemplateFolder               string            // If this is set, we will create
	TemplateSingleton            string            //
	TemplateExtension            string            //
	CreateFromTemplate           bool
	Template                     []Template // These are any custom templates
	CreateRouter                 bool       // Do we want to create the router. This should only be set to false in specific cases
	CreateHandler                bool       // Do we want to create the handler. This should only be set to false in specific cases
	CreateMiddleware             bool       // Do we want to create the middleware ( struct validation and auth ). This should only be set to false in specific cases
	RouterTemplate               string
	HandlerTemplate              string
	ValidationMiddlewareTemplate string
	AuthMiddlewareTemplate       string
	PresenterTemplate            string

	OverrideFiles bool // If we write the files to disk, and a file already exists, do we want to override the current files contents
	WriteToDisk   bool // Will we write the files to disk
}

type Template struct {
	Name    string // This is the name of the file
	OutPath string // This is the output path, relative to where we are working from
	Content string // This is the template
	Package string
}

// FilepathSeparator will return the file path separator, it will generally return the filepath.Separator
func (c *Config) FilepathSeparator() string {
	if c.Separator == 0 {
		c.Separator = filepath.Separator
	}
	return string(c.Separator)
}

// validateFinalCharacter will return the file path separator, it will generally return the filepath.Separator
// Examples
// validateFinalCharacter(".","/") == "./"
// validateFinalCharacter("","/") == "/"
func validateFinalCharacter(base, sep string) string {
	if base == "" {
		return sep
	}
	if base[len(base)-1:] != sep {
		return base + sep
	}
	return base
}

// getPathWithSeparator will return the file path
func (c *Config) getPathWithSeparator(base, dir, sep string) string {
	if base == "" {
		base = "." + sep
	}
	// We make sure the final character is the separator
	base = validateFinalCharacter(base, sep)
	if c.Version == "" {
		return fmt.Sprintf("%s%s", base, dir)
	} else {
		return fmt.Sprintf("%s%s%s%s", base, dir, sep, strings.ToLower(c.Version))
	}
}

// getFullPath will return the full file path
func (c *Config) getFullPath(base, dir string) string {
	return c.getPathWithSeparator(base, dir, c.FilepathSeparator()) + c.FilepathSeparator()
}
func (c *Config) GetAPIPath(base string) string {
	sep := "/"
	if c.APIPath == "" {
		c.APIPath = sep
	}

	if c.Version == "" {
		return getAPIPath(c.APIPath, base, sep)
	} else {
		apiPath := c.APIPath
		if !strings.Contains(strings.ToLower(c.APIPath), strings.ToLower(c.Version)) {
			apiPath = getAPIPath(apiPath, strings.ToLower(c.Version), sep)
		}
		return getAPIPath(apiPath, base, sep)
	}
}

func getAPIPath(start, end, sep string) string {
	//fmt.Println("getAPIPath", start, end, sep)
	if start == "" {
		return sep
	}
	if start == sep {
		return start + end
	}
	if start[:1] != sep {
		start = sep + start
	}
	if end == "" {
		return start
	}
	return start + sep + end
}

// getFullFilePath will return the full file path
func (c *Config) getFullFilePath(base, dir, fileName string) string {
	return fmt.Sprintf("%s%s", c.getFullPath(base, dir), fileName)
}

// SetDefaults sets the config values
func (c *Config) SetDefaults() {
	//We set up the default directory information
	if c.OutputDir == "" {
		c.OutputDir = "." + c.FilepathSeparator()
	}
	if c.OutputDirPermissions == 0 {
		c.OutputDirPermissions = 0755
	}
	if c.Directories.RouterDirectory == "" {
		c.Directories.RouterDirectory = "router"
	}
	if c.Directories.HandlerDirectory == "" {
		c.Directories.HandlerDirectory = "handler"
	}
	if c.Directories.MiddlewareDirectory == "" {
		c.Directories.MiddlewareDirectory = "middleware"
	}
	if c.Directories.PresenterDirectory == "" {
		c.Directories.PresenterDirectory = "presenter"
	}

	// We set up the default templates
	if c.RouterTemplate == "" {
		c.RouterTemplate = templates.RouterTemplate
	}
	if c.HandlerTemplate == "" {
		c.HandlerTemplate = templates.HandlerTemplate
	}
	if c.AuthMiddlewareTemplate == "" {
		c.AuthMiddlewareTemplate = templates.AuthMiddlewareTemplate
	}
	if c.ValidationMiddlewareTemplate == "" {
		c.ValidationMiddlewareTemplate = templates.StructValidationMiddlewareTemplate
	}
	if c.PresenterTemplate == "" {
		c.PresenterTemplate = templates.PresenterTemplate
	}

	if c.Directories.IntermediaryPath == "" {
		// We need to calculate the Path
	}

	if c.Directories.Module == "" {
		//c.GetModuleName()
		// We need to calculate the Module
	}

	c.SetModule()
}

// SetupFolders will create all the folders required for the packages
func (c *Config) SetupFolders() error {
	if c.CreateFromTemplate {
		for _, template := range c.Template {
			err := createDir(template.OutPath, c.OutputDirPermissions)
			if err != nil {
				return err
			}
		}
		return nil
	}

	var err error
	fmt.Println("Directory:", c.OutputDir)
	err = createDir(c.OutputDir, c.OutputDirPermissions)
	if err != nil {
		return err
	}

	if c.CreateMiddleware {
		err = createDir(c.getFullPath(c.OutputDir, c.Directories.MiddlewareDirectory), c.OutputDirPermissions)
		if err != nil {
			return err
		}
	}

	if c.CreateRouter {
		err = createDir(c.getFullPath(c.OutputDir, c.Directories.RouterDirectory), c.OutputDirPermissions)
		if err != nil {
			return err
		}
	}

	if c.CreateHandler {
		err = createDir(c.getFullPath(c.OutputDir, c.Directories.HandlerDirectory), c.OutputDirPermissions)
		if err != nil {
			return err
		}
	}

	if c.CreateMiddleware || c.CreateRouter || c.CreateHandler {
		err = createDir(c.getFullPath(c.OutputDir, c.Directories.PresenterDirectory), c.OutputDirPermissions)
		if err != nil {
			return err
		}
	}

	return nil
}

func createDir(dir string, outputDirPermissions os.FileMode) error {
	err := os.MkdirAll(dir, outputDirPermissions)
	if err != nil {
		fmt.Println("Error processing Directory:", dir, err)
		return err
	}
	return nil
}

// Process will process the request
func (c *Config) Process() error {
	// We make sure the config has been set up
	c.SetDefaults()

	//AstReader(c)

	// we need to process the files
	err, spec := GoDocReader(c)
	if err != nil {
		return err
	}

	if c.CreateFromTemplate {
		// We need to read all the templates
		err = c.GetAllTemplates()
		if err != nil {
			return err
		}
	}

	err = c.CreateAll(spec)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) GetDefaultTypeSpec() TypeSpec {
	return TypeSpec{
		Backtick:         "`",
		Version:          c.Version,
		APIPath:          c.APIPath,
		Struct:           StructSpec{},
		Module:           c.Directories.Module,
		IntermediaryPath: c.Directories.IntermediaryPath,
		Auth:             c.Auth,
	}
}

// CreateAll will create everything needed for the templates and generate the templates
func (c *Config) CreateAll(defs []TypeSpec) error {
	// We make sure all the folders are created
	err := c.SetupFolders()
	if err != nil {
		return err
	}

	if !c.CreateFromTemplate {
		ts := c.GetDefaultTypeSpec()
		ts.Struct.APIName = "presenter"
		ts.Package = "presenter"
		file, errGen := ts.Generate(
			c.PresenterTemplate,
			c.getFullPath(
				c.OutputDir,
				c.Directories.PresenterDirectory,
			),
			"presenter.go",
		)
		if errGen != nil {
			fmt.Println(errGen)
			return errGen
		}

		if c.WriteToDisk {
			create := false
			if c.WriteToDisk {
				create = true
			}

			if create {
				err := os.WriteFile(
					file.Path+file.Name,
					[]byte(file.Content),
					c.OutputDirPermissions,
				)
				if err != nil {
					fmt.Println("Processing File Error:", file.Path, file.Name, err)
					return err
				}
			}
		} else {
			// We do not want to write the file. So we are just going to print the content
			fmt.Println(file.Content)
		}
		// We only want to create the auth if Auth is set
		if c.Auth {
			ts.Struct.APIName = "middleware"
			ts.Package = "middleware"
			file, err = ts.Generate(c.AuthMiddlewareTemplate,
				c.getFullPath(
					c.OutputDir,
					c.Directories.MiddlewareDirectory,
				),
				"middleware.auth.go",
			)
			if err != nil {
				return err
			}
			if c.WriteToDisk {
				err := os.WriteFile(
					file.Path+file.Name,
					[]byte(file.Content),
					c.OutputDirPermissions,
				)
				if err != nil {
					fmt.Println("Processing File Error:", file.Path, file.Name, err)
					return err
				}
			} else {
				// We do not want to write the file. So we are just going to print the content
				fmt.Println(file.Content)
			}
		}

	} else {
		for _, template := range c.Template {
			if strings.Contains(strings.ToLower(template.Name), c.TemplateSingleton) {
				err2 := c.CreateFileFromTemplate(template)
				if err2 != nil {
					return err2
				}
			}
		}
	}

	// We are going to process every TypeSpec
	for _, def := range defs {
		if def.Ignored {
			break
		}
		fmt.Println("Processing Struct:", def.Struct.Name)

		all := make([]*File, 0)
		all, err = def.GetAllFiles(c)
		if err != nil {
			fmt.Println("GetAllFiles error: ", err)
			return err
		}

		for _, file := range all {
			// We need to write the file
			fmt.Println("Processing File:", file.Path, file.Name)
			if c.WriteToDisk {
				err := os.WriteFile(
					file.Path+file.Name,
					[]byte(file.Content),
					c.OutputDirPermissions,
				)
				if err != nil {
					fmt.Println("Processing File Error:", file.Path, file.Name, err)
					return err
				}
			} else {
				// We do not want to write the file. So we are just going to print the content
				fmt.Println(file.Content)
			}
		}
	}
	return nil
}

func (c *Config) CreateFileFromTemplate(template Template) error {
	ts := c.GetDefaultTypeSpec()
	ts.Struct.APIName = template.Package
	ts.Package = template.Package
	file, err := ts.Generate(template.Content,
		template.OutPath,
		template.Name,
	)
	if err != nil {
		return err
	}
	if c.WriteToDisk {
		err := os.WriteFile(
			template.OutPath+template.Name,
			[]byte(file.Content),
			c.OutputDirPermissions,
		)
		if err != nil {
			fmt.Println("Processing File Error:", file.Path, file.Name, err)
			return err
		}
	} else {
		// We do not want to write the file. So we are just going to print the content
		fmt.Println(file.Content)
	}
	return nil
}

func (file *File) Save(separator string, saveToDisk bool, filePerms os.FileMode) error {
	if saveToDisk {
		err := os.WriteFile(fmt.Sprintf("%s%s%s", file.Path, separator, file.Name), []byte(file.Content), filePerms)
		return err
	}
	fmt.Println(file.Content)
	return nil
}

func (c *Config) SetModule() {
	if c.Directories.Module == "" {
		getwd, err := os.Getwd()
		if err != nil {
			return
		}

		module, importPathDiff, err := getProjectInformation(getwd, "")
		if err != nil {
			return
		}
		c.Directories.IntermediaryPath = importPathDiff
		c.Directories.Module = module
	}
}

// GetAllTemplates This will recursively read all files and folders in `Config.TemplateFolder`
// adding all the files to `Config.Template`
// The file content will be the `Template.Content`
// file name will be `Template.NameFormat` with `%s.` as a prefix
// the file path will be the OutPath with `Config.TemplateFolder` being stripped from the path and replaced with `.`
func (c *Config) GetAllTemplates() error {
	packageRegex := regexp.MustCompile(`package (\w+)`)
	if c.Template == nil {
		c.Template = make([]Template, 0)
	}
	templateExtension := c.TemplateExtension
	err := filepath.Walk(c.TemplateFolder,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.Contains(strings.ToLower(info.Name()), templateExtension) {
				content, readErr := os.ReadFile(path)
				if readErr != nil {
					return readErr
				}

				matches := packageRegex.FindStringSubmatch(string(content))
				path = strings.ReplaceAll(path, info.Name(), "")
				template := Template{
					Content: "//Package {{.Package}} code generated by team142\n" + string(content),
					Name:    strings.ReplaceAll(info.Name(), templateExtension, ""),
					OutPath: c.getFullPath(c.OutputDir, strings.ReplaceAll(path, c.TemplateFolder, "")),
				}
				if len(matches) > 0 {

					template.Package = strings.Split(matches[0], "")[1]
				}

				c.Template = append(c.Template, template)
			}

			return nil
		})

	return err
}
