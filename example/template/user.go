//go:generate go run ../../cmd/api/main.go -i user.go -t template
package basic

// User is a basic user of the system
// @DiscoverFunction
// @BasePath /api
// @Version
type User struct {
	//Hello there
	//#PK
	Id        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"` // FirstName fn
	LastName  string `json:"last_name,omitempty"`  // LastName ln
	CreatedAt string `json:"-"`
}

// UserIgnored should be ignored during the generation
// #IGNORE
type UserIgnored struct {
}
