package routes

import (
	c "implement_middleware/constant"
	"implement_middleware/controller/books"
	users "implement_middleware/controller/user"
	middlewares "implement_middleware/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() *echo.Echo {
	e := echo.New()

	// use middleware log
	middlewares.LogMiddleware(e)

	// login user
	e.POST("/login", users.LoginUsersController)

	// users routes not use middleware
	e.POST("/register", users.RegisterController)
	e.GET("/books", books.GetBooksController)
	e.GET("/books/:id", books.GetBookController)

	// use middleware jwt
	m := e.Group("")
	m.Use(middleware.JWT([]byte(c.SECRET_JWT)))
	m.GET("/users", users.GetUsersController)
	m.GET("/users/:id", users.GetUserController)
	m.DELETE("/users/:id", users.DeleteUserController)
	m.PUT("/users/:id", users.UpdateUserController)
	m.POST("/books", books.CreateBookController)
	m.DELETE("/books/:id", books.DeleteBookController)
	m.PUT("/books/:id", books.UpdateBookController)

	return e
}
