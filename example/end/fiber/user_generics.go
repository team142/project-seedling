package end

import (
	"database/sql"
	"errors"
	"strconv"
)

var (
	// Please note if you change this mapping, you will need to change the swagger documentation
	// ParamToField has the mapping for api parameters to fields
	ParamToField = map[string]string{
		"id": "Id",
	}
)

// GetMultipleUsers will get `end.User` by the parameters passed in the request
// Returning error, []Users, limit, nextId, total
func (u *User) GetMultipleUsers(db *sql.DB, params map[string]string) (error, []User, int, int, int64) {
	//TODO implement me
	limit := 100
	if k, ok := params["limit"]; ok {
		var err error
		limit, err = strconv.Atoi(k)
		if err != nil {
			return err, nil, 0, 0, 0
		}
	}
	return errors.New("implement me"), nil, 0, 0, 0
}

func (u *User) UpdateUser(db *sql.DB, id int) error {
	//TODO implement me
	return errors.New("implement me")
}

func (u *User) DeleteUser(db *sql.DB, id int) error {
	//TODO implement me
	return errors.New("implement me")
}

func (u *User) Save(db *sql.DB, override bool) (error, bool) {
	//TODO implement me
	return errors.New("implement me"), false
}

func (u *User) Validate() error {
	//TODO implement me
	return errors.New("implement me")
}

func (u *User) GetPrimaryKey() *int {
	if u.Id == 0 {
		return nil
	}
	return &u.Id
}
