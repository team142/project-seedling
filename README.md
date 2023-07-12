# Golang API Code Gen

The main goals for this project:
1. Everything must be driven from a struct ( including the struct tags )
2. Automagically create the functions for a web framework:
   1: Fiber -> https://github.com/gofiber/recipes/tree/master/clean-architecture
      1. Handlers
      2. Routers
3. Auto Create the docs required for https://github.com/gofiber/recipes/tree/master/swagger
4. No Lock In
   1. Authentication
   2. Database
   3. Cache
   4. Web Framework

