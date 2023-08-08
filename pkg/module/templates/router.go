package templates

const RouterTemplate = `// code generated by team142
package {{.Package}}

import (
	"{{.Struct.HandlerImportPath}}"
	"{{.Struct.MiddlewareImportPath}}"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

const (
	RoleRead{{.Struct.Name}}   = "read_{{.Struct.APIName}}"
	RoleCreate{{.Struct.Name}} = "create_{{.Struct.APIName}}"
	RoleDelete{{.Struct.Name}} = "delete_{{.Struct.APIName}}"
)

func {{.Struct.Name}}Routes(api fiber.Router, db *sql.DB) { 
	api.Options("/{{.Struct.APIName}}",
		middleware.AuthMiddleware(db, nil),
		handler.{{.Struct.Name}}Options(),
	){{if eq .PrimaryKeyCount 1}}
	api.Get("/{{.Struct.APIName}}/:id",  
		middleware.AuthMiddleware(db, []string{ RoleRead{{.Struct.Name}} }),  
		handler.Get{{.Struct.Name}}ById(db),
	){{end}}
	api.Get("/{{.Struct.APIName}}",  
		middleware.AuthMiddleware(db, []string{ RoleRead{{.Struct.Name}} }),  
		handler.GetMultiple{{.Struct.Name}}s(db),
	)
	api.Put("/{{.Struct.APIName}}",  
		middleware.AuthMiddleware(db, []string{ RoleCreate{{.Struct.Name}} }),  
		middleware.Verify{{.Struct.Name}}Body(),  
		handler.Save{{.Struct.Name}}(db, true),
	)
	api.Post("/{{.Struct.APIName}}",  
		middleware.AuthMiddleware(db, []string{ RoleCreate{{.Struct.Name}} }),  
		middleware.Verify{{.Struct.Name}}Body(),  
		handler.Save{{.Struct.Name}}(db, false),
	)
	api.Delete("/{{.Struct.APIName}}",  
		middleware.AuthMiddleware(db, []string{ RoleDelete{{.Struct.Name}} }),  
		middleware.Verify{{.Struct.Name}}Body(),  
		handler.Delete{{.Struct.Name}}(db),
	)
	api.Delete("/{{.Struct.APIName}}/multi",  
		middleware.AuthMiddleware(db, []string{ RoleDelete{{.Struct.Name}} }),  
		middleware.Verify{{.Struct.Name}}Body(),  
		handler.Delete{{.Struct.Name}}(db),
	){{if eq .PrimaryKeyCount 1}}
	api.Delete("/{{.Struct.APIName}}/:id",  
		middleware.AuthMiddleware(db, []string{ RoleDelete{{.Struct.Name}} }),  
		middleware.Verify{{.Struct.Name}}Body(), handler.Delete{{.Struct.Name}}ById(db).
	){{end}}
	api.Trace("/{{.Struct.APIName}}",  
		middleware.AuthMiddleware(db, nil),  
		middleware.Verify{{.Struct.Name}}Body(),  
		handler.Trace{{.Struct.Name}}(),
	)
}
`
