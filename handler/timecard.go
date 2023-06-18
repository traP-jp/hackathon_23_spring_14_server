package handler

import (
	"fmt"
	"net/http"

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
	card, err := model.AddTimeCards(rawuuid, rawtuid, userid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, card)
}
