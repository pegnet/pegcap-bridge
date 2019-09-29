package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

var Genesis = 206422

func (api *Api) Heights(c echo.Context) error {
	heights := make(map[string]int)
	heights["current"] = 211173
	heights["start"] = 206422
	js, _ := json.Marshal(heights)
	return c.JSONBlob(http.StatusOK, js)
}
