package handler

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/hackathon_23_spring_14_server/model"
)

func GetItems(c echo.Context) error {
	includeSuspended := c.QueryParam("include-suspended")

	rawitems, err := model.GetItems()
	if includeSuspended == "false" {
		rawitems, err = model.GetActiveItems()
	}
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rawitems)
}

func AddItems(c echo.Context) error {
	id := c.QueryParam("title")
	description := c.QueryParam("description")
	points := c.QueryParam("score")
	var point int
	switch points {
	case "great":
		point = 2
	case "good":
		point = 1
	case "bad":
		point = -1
	case "terrible":
		point = -2
	}
	if item, err := model.EnsureExisistenceID(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if len(item) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "item already exists")
	} else {
		fmt.Println(item)
	}
	item := model.PublicItem{
		UUID:        uuid.Nil,
		ID:          id,
		Description: description,
		Point:       point,
		Report:      0,
	}
	returnItem, err := model.AddItems(item)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, returnItem)
}
