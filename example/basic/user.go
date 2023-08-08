//go:generate go run ../../cmd/api/main.go -i user.go -version v1
package basic

// User is a basic user of the system
// @DiscoverFunction
// @BasePath /api
// @Version
// #GET AUTH
// #POST AUTH
// #DELETE AUTH
type User struct {
	//Hello there
	//#PK
	Id int `json:"id,omitempty" fun-gen:"get" api-gen:"get,post,delete" attributes:"pk"` //#PK
	//@API
	FirstName string `json:"first_name,omitempty"` // FirstName fn
	//@API
	LastName string `json:"last_name,omitempty"` // LastName ln
}

// UserIgnored should be ignored during the generation
// @IGNORE
type UserIgnored struct {
}
