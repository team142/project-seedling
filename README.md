# Golang API Code Gen - Project Seedling

This is an opinionated generator, to help build out and stub API for projects.

This generator is used to get a project up to speed as soon as possible with as little effort as possible.

The generation is opinionated, but trying to keep as little lock in as possible.
You can use the generator to produce a single executable or the functions for a microservice architecture.

# Quick Start
```golang
//go:generate go run github.com/team142/project-seedling -i user.go -version v1 -auth
package basic

// User is a basic user structure for the system
// @BasePath /api
// @Version
// #GET AUTH
// #POST AUTH
// #DELETE AUTH
type User struct {
	//@API
	//#PK
	Id int `json:"id,omitempty"`
	//@API
	FirstName string `json:"first_name,omitempty"`
	//@API
	LastName string `json:"last_name,omitempty"`
	//#Ignore
	CreatedAt string `json:"-"`
}
```

`go generate ./...`

# Goals

## The full `planned` ecosystem:
1. [ ] DB to YAML
2. [ ] YAML to struct
   1. YAML to DB
   2. YAML to DB changes
3. [X] Generate go files from templates 
   1. [ ] Example: Struct Core Functions
   2. [X] Example: Struct API Functions
   3. [ ] Example: API Client - This will generate a **SDK** to use the API's
      1. GO


##  Ecosystem Goals


##  Project Goals

1. Everything must be driven from a struct ( including the struct tags )
   1. I don't believe it should be part of the CI process
   2. Although it can be
2. Automagically create the functions for a web framework:
   1. Fiber -> https://github.com/gofiber/recipes/tree/master/clean-architecture
      1. Handlers
      2. Routers
      3. Presenters  
      4. Common Middleware
3. Auto Create the docs required for https://github.com/gofiber/recipes/tree/master/swagger
4. No Lock In, BUT opinionated ( we cannot support every use case )
   1. Authentication
   2. Database
   3. Cache
   4. Web Framework

The objective is to allow for CRUD operations on the struct ( Create, Read, Update, Delete )
* GET ( Read, Singular and Multiple )
* POST ( Insert/Update )
* PUT ( Insert/Override )
* DELETE ( Delete )

# USAGE
There are a couple simple ways for one to use the project

```
//go:generate github.com/team142/project-seedling -i user.go
```

Using if you provide your own templates
```
//go:generate github.com/team142/project-seedling -template template
```

A More complex 

```
//go:generate github.com/team142/project-seedling -i user.go -version v1 -api fiber -s User,UserRole -o ../../
```

# Example
Example are in the `example` folder