package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

//http://localhost:8000/gophers/json?name=arnold&type=fluffy
func GetGophers(c echo.Context) error {
	gopherName := c.QueryParam("name")
	gopherType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your gopher name is : %s\nand gopher type is : %s\n", gopherName, gopherType))
	} else if dataType == "json" {
		gopher := vo.gopher{
			Name: gopherName,
			Type: gopherType,
		}
		return c.JSON(http.StatusOK, gopher)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as String or JSON",
		})
	}

}

func AddGopher(c echo.Context) error {
	gopher := vo.gopher{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&gopher)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	log.Printf("this is your gopher %#v", gopher)
	return c.String(http.StatusOK, "We got your gopher!!!")
}
