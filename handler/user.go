package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/hackathon_23_spring_14_server/model"
)

type (
	User struct {
		ID     string `json:"id"`
		UserID string `json:"uid"`
		Score  int    `json:"score"`
	}
)

func GetUsers(c echo.Context) error {
	rawUsers, err := model.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rawUsers)
}
func GetMe(c echo.Context) error {
	rawUsers, err := model.GetMe()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rawUsers)
}

func GetUserSpecific(c echo.Context) error {
	uid := c.Param("uid")
	rawUsers, err := model.GetUserSpecific(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rawUsers)
}

func GetRanking(c echo.Context) error {
	rawUsers, err := model.GetRanking()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, rawUsers)
}
