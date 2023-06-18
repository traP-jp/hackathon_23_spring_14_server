package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/hackathon_23_spring_14_server/model"
)

func AddTimeCards(c echo.Context) error {

	fmt.Println("AddTimeCards")
	rawuuid := c.QueryParam("uuid")
	rawtuid := c.QueryParam("tuid")
	fmt.Println(rawuuid)
	fmt.Println(rawtuid)
	userid := c.Get("userid").(string)
	model.AddTimeCards(rawuuid, rawtuid, userid)
	return nil
}