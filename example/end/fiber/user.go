package end

// User is a basic user of the system
// @DiscoverFunction
// @BasePath /api
// @Version
// #GET AUTH ROLE
// #POST AUTH ROLE
// #DELETE AUTH ROLE
type User struct { // @NOOOO
	Id        int        `json:"id,omitempty"`         // Id is the PK
	FirstName string     `json:"first_name,omitempty"` // FirstName fn
	LastName  string     `json:"last_name,omitempty"`  // LastName ln
	Languages []Language `json:"languages,omitempty"`
}

// Language is a basic user of the system
// @BasePath /user/:id
type Language struct {
	Id       int    `json:"id,omitempty"`
	Language string `json:"language,omitempty"`
	Country  string `json:"country,omitempty"`
}
