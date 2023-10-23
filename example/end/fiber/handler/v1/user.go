package handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	end "github.com/team142/project-seedling/example/end/fiber"
	"github.com/team142/project-seedling/example/end/fiber/presenter/v1"
	"strconv"
	"strings"
)

// UserInterface is the core interface required for the API to be useful
type UserInterface interface {
	// GetMultipleUsers will get `end.User` by the parameters passed in the request
	// Returning error, []Users, limit, nextId, total
	GetMultipleUsers(db *sql.DB, params map[string]string) (error, []end.User, int, int, int64)

	// DeleteUser will delete `end.User` by the data passed in the body
	DeleteUser(db *sql.DB, id int) error

	// Save will update or create `end.User`
	// if override is passed, the update will update every field.
	// if the return bool is true, it means the User was created
	// if the return bool is false, it means the User was updated/overridden
	Save(db *sql.DB, override bool) (error, bool)

	// Validate will validate if the content is valid, returning nil if valid
	Validate() error

	// GetPrimaryKey will return the Primary key of the object.
	// If the database is a composite key, the key will be a concatenation of the values
	// If the primary key is not set a nil value will be returned
	GetPrimaryKey() *int
}

// GetMultipleUsers  get multiple `end.User` filter by the query params
// @Summary          get multiple `end.User` filter by the query params
// @Description      Get multiple users based on the query params
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          200  {object} presenter.DataArraySuccess{data=[]end.User}
// @Success          204  {object} presenter.DataArraySuccess{data=[]end.User}
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          500  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user [get]
// @Param            limit 			query      int     	false  "Limit" default(100)
// @Param            id   			query      int  	false  "Id"
// @Param            first_name		query      string  	false  "FirstName"
// @Param            last_name		query      string	false  "LastName"
func GetMultipleUsers(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user UserInterface
		u := &end.User{}
		user = u

		// STANDARD
		var err error
		var limit int
		var total int64

		// STRUCT SPECIFIC
		var users = make([]end.User, 0, 0)
		var nextId int

		// We are going to get multiple users and the required fields for the response
		if err, users, limit, nextId, total = user.GetMultipleUsers(
			db,
			c.AllParams(),
		); err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(
				presenter.DataError{
					Error: err,
				},
			)
		} else {
			first := end.User{}
			if len(users) > 0 {
				first = users[0]
			} else {
				c.Status(fiber.StatusNoContent)
				return c.JSON(
					presenter.DataArraySuccess{
						Data:        users,
						Limit:       limit,
						ReturnValue: len(users),
						Total:       total,
					},
				)
			}
			c.Status(fiber.StatusOK)
			return c.JSON(
				presenter.DataArraySuccess{
					Data:        users,
					Limit:       limit,
					ReturnValue: len(users),
					From:        first.Id,
					Next:        nextId,
					Total:       total,
				},
			)
		}
	}
}

// GetUserById get a single `end.User` by the `end.User.Id`
// @Summary          This will get a single `end.User` filtering using the `User.Id`
// @Description      Get a single User
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          200  {object} presenter.Data{data=end.User}
// @Success          204  {object} presenter.DataError{data=string}
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          500  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user/{id} [get]
// @Param         	 id		path      int  true  "Id"
func GetUserById(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var user UserInterface
		u := &end.User{}
		user = u

		var users []end.User
		if err, users, _, _, _ = user.GetMultipleUsers(
			db,
			map[string]string{"Id": c.Params("Id", "")},
		); err != nil {
			if strings.Contains(err.Error(), "not found") {
				c.Status(fiber.StatusNoContent)
				return c.JSON(
					presenter.DataError{
						Message: "not found",
						Data:    u,
						Error:   err,
					},
				)
			}

			c.Status(fiber.StatusInternalServerError)
			return c.JSON(
				presenter.DataError{
					Message: "error retrieving requested User",
					Data:    u,
					Error:   err,
				},
			)
		} else {
			if len(users) > 0 {
				c.Status(fiber.StatusAccepted)
				return c.JSON(
					presenter.DataSuccess{
						Data: users[0],
					},
				)
			} else {
				c.Status(fiber.StatusNoContent)
				return c.JSON(
					presenter.DataSuccess{},
				)
			}
		}
	}
}

// TraceUser trace a user request
// @Summary          Validate is the user sent in the body is accepted by the server
// @Description      Trace for User
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          202  {object} presenter.DataSuccess{data=end.User}
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          406  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user [trace]
func TraceUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		user := getUser(c)

		if err = user.Validate(); err != nil {
			c.Status(fiber.StatusNotAcceptable)
			return c.JSON(
				presenter.DataError{
					Message: "invalid data",
					Data:    user,
					Error:   err,
				},
			)
		}

		c.Status(fiber.StatusAccepted)
		return c.JSON(
			presenter.DataSuccess{
				Data: user,
			},
		)
	}
}

// SaveUser trace a user request
// @Summary          Save a user
// @Description      This will create a new User or replace/update a representation of the User with the request payload.
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          202  {object} presenter.DataSuccess{data=end.User}
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          406  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user [put, post]
func SaveUser(db *sql.DB, override bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		user := getUser(c)

		var created bool
		err, created = user.Save(db, override)
		if err != nil {
			c.Status(fiber.StatusNotAcceptable)
			return c.JSON(
				presenter.DataError{
					Message: "invalid data",
					Data:    user,
					Error:   err,
				},
			)
		}

		if created {
			c.Status(fiber.StatusCreated)
		} else {
			c.Status(fiber.StatusOK)
		}

		return c.JSON(
			presenter.DataSuccess{
				Data: user,
			},
		)
	}
}

// DeleteUser trace a user request
// @Summary          PUT a user
// @Description      This will delete User or replace a representation of the User with the request payload.
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          202  ""
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          406  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user [delete]
func DeleteUser(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		user := getUser(c)

		pk := user.GetPrimaryKey()
		if pk == nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(
				presenter.DataError{
					Message: "bad primary key",
					Error:   err,
				},
			)
		}

		err = user.DeleteUser(db, *pk)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(
				presenter.DataError{
					Message: "internal database error",
					Data:    user,
					Error:   err,
				},
			)
		}

		c.Status(fiber.StatusOK)
		return c.Send(nil)
	}
}

// DeleteUserById deletes a user
// @Summary          DELETE a user
// @Description      This will delete User or replace a representation of the User with the request payload.
// @Tags             User
// @Accept           json,text/xml
// @Produce          json
// @Success          202  ""
// @Failure          400  {object} presenter.DataError{data=string}
// @Failure          406  {object} presenter.DataError{data=end.User}
// @Router           /api/v1/user [delete]
// @Param         	 id		path      int  true  "Id"
func DeleteUserById(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		user := getUser(c)

		// This will only do this if the PK is an INT
		pk, errAtoi := strconv.Atoi(c.Params("Id", ""))
		if errAtoi != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(
				presenter.DataError{
					Message: "bad id",
					Error:   err,
				},
			)
		}

		err = user.DeleteUser(db, pk)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(
				presenter.DataError{
					Message: "internal database error",
					Data:    user,
					Error:   err,
				},
			)
		}

		c.Status(fiber.StatusOK)
		return c.Send(nil)
	}
}

func getUser(c *fiber.Ctx) UserInterface {
	userLocal := c.Locals("User")
	if userLocal == nil {
		//return errors.New("user local not set")
		return nil
	}
	return userLocal.(*end.User)
}
