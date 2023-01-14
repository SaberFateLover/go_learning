package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	f := func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	}
	e.GET("/", f)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)

	type User struct {
		Name  string `json:"name" xml:"name" form:"name" query:"name"`
		Email string `json:"email" xml:"email" form:"email" query:"email"`
	}

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":8080"))

}

func saveUser(c echo.Context) error {
	return nil
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)

}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
