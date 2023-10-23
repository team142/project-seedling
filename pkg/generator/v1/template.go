package generator

import (
	"fmt"
	"github.com/team142/project-seedling/pkg/module/v1"
)

func TemplateGenerator(conf *module.Config, spec []module.TypeSpec) (error, []module.File) {
	files := make([]module.File, 0)

	for _, typeSpec := range spec {
		fmt.Println(typeSpec.Struct)
		for _, field := range typeSpec.Fields {
			fmt.Printf("%+v\n", field)
		}
	}

	return nil, files
}
