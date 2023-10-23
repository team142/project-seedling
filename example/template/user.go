//go:generate go run ../../cmd/template/v1/main.go -i user.go -t template
package basic

import "database/sql"

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

func (u *User) GetMultiple(db *sql.DB, params map[string]string) (error, []User, int, int, int64) {
	//TODO implement me
	panic("implement me")
}

func (u *User) Delete(db *sql.DB) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) Save(db *sql.DB, override bool) (error, bool) {
	//TODO implement me
	panic("implement me")
}

func (u *User) Validate() error {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetDeleteStatement() (error, string) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetSelectStatement(params map[string]string) (error, string) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUpdateStatement() (error, string) {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetCreateStatement() (error, string) {
	//TODO implement me
	panic("implement me")
}

// UserIgnored should be ignored during the generation
// #IGNORE
type UserIgnored struct {
}
