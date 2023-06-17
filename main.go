package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
	"github.com/traP-jp/hackathon_23_spring_14_server/handler"
	"github.com/traP-jp/hackathon_23_spring_14_server/model"
)

func main() {
	if err := model.Setup(); err != nil {
		panic(err)
	}
	e := echo.New()
	e.Use(mid.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	api := e.Group("/api")
	{

		api.File("/swagger.yaml", "./document/swagger.yaml")
		api.Static("/", "./document/swagger-ui/dist")
		api.Any("", func(c echo.Context) error {
			return c.Redirect(http.StatusFound, c.Path()+"/")
		})

		apiUser := api.Group("/user")
		{
			apiUser.GET("", handler.GetUsers)
		}
		apiItem := api.Group("/item")
		{
			apiItem.GET("", handler.GetItems)
			apiItem.POST("", handler.AddItems)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
