package router

import (
	"code-gen/example/end/fiber/handler/v1"
	"code-gen/example/end/fiber/middleware"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(api fiber.Router, db *sql.DB) {
	api.Get("/user/:id", middleware.AuthMiddleware(db, []string{"read_user"}), handler.GetUserById(db))
	api.Get("/user", middleware.AuthMiddleware(db, []string{"read_user"}), handler.GetMultipleUsers(db))
	api.Put("/user", middleware.AuthMiddleware(db, []string{"create_user"}), middleware.VerifyUserBody(), handler.SaveUser(db, true))
	api.Post("/user", middleware.AuthMiddleware(db, []string{"create_user"}), middleware.VerifyUserBody(), handler.SaveUser(db, false))
	api.Delete("/user", middleware.AuthMiddleware(db, []string{"delete_user"}), middleware.VerifyUserBody(), handler.DeleteUser(db))
	api.Delete("/user/:id", middleware.AuthMiddleware(db, []string{"delete_user"}), middleware.VerifyUserBody(), handler.DeleteUserById(db))
	api.Trace("/user", middleware.AuthMiddleware(db, nil), middleware.VerifyUserBody(), handler.TraceUser())
}