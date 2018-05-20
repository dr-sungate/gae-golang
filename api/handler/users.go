package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

type Users struct {
}

type (
	user struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users map[string]user
)

func init() {
	users = map[string]user{
		"1": user{
			ID:   "1",
			Name: "Wreck-It Ralph",
		},
	}
}

func (us Users) CreateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = *u
	return c.JSON(http.StatusCreated, u)
}

func (us Users) GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func (us Users) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users[c.Param("id")])
}