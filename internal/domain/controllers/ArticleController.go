package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func StoreArticle(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("title")
	fmt.Println("title:", title)
	fmt.Println("description:", description)
	return c.String(http.StatusOK, "title:"+title+", desc:"+description)
}
func GetArticle(c echo.Context) error {
	slug := c.Param("slug")
	return c.String(http.StatusOK, slug)
}
