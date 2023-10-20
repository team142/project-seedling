# This is the end goal

This is the end goal.

We will have simple ONE struct
```go
type User struct {
	Id        int
	FirstName string
	LastName  string
}
```

This should generate a route and handler for each operation.
The objective is to allow for CRUD operations on the `User` struct ( Create, Read, Update, Delete )
* GET ( Read )
* POST ( Insert/Update )
* PUT ( Insert/Override )
* DELETE ( Delete )