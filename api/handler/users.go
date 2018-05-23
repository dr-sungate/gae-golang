package handler

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"gaeapp/gae-golang/api/dao"
	log "gaeapp/gae-golang/api/logger"
)

type Users struct {
}



func (us Users) CreateUser(c echo.Context) error {
	c.Request().ParseForm()
	userdao := dao.UserDAONew(c.Request())
	userdao.SetData(c.FormValue("name"), c.FormValue("email"), true)
	if err := userdao.Create(); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, userdao.User)
}

func (us Users) GetUsers(c echo.Context) error {
	log.Warn(c.QueryParams())
	userdao := dao.UserDAONew(c.Request())
	if err := userdao.GetAll(c.QueryParams()); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, userdao.Users)
}

func (us Users) GetUser(c echo.Context) error {
	userdao := dao.UserDAONew(c.Request())
	idint, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := userdao.Get(idint); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, userdao.User)
}

func (us Users) DeleteUser(c echo.Context) error {
	userdao := dao.UserDAONew(c.Request())
	idint, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := userdao.Delete(idint); err != nil {
		log.Error(err)
		return err
	}
	return c.JSON(http.StatusOK, userdao.User)
}

