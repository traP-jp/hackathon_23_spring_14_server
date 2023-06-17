package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/hackathon_23_spring_14_server/model"
)

func GetItems(c echo.Context) error {
	includeSuspended := c.QueryParam("include-suspended")
	var include bool
	if includeSuspended == "true" {
		include = true
	} else {
		include = false
	}
	var items []model.Items
	items = model.GetItems(include)

	return c.JSON(http.StatusOK, items)
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
	item := model.Items{
		ID:          id,
		Description: description,
		Point:       point,
		Report:      0,
	}
	fmt.Println(item)
	model.AddItems(item)

	return c.JSON(http.StatusOK, item)
}
