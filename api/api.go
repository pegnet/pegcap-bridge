package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/syndtr/goleveldb/leveldb"
)

type Api struct {
	DB *leveldb.DB
}

func (api *Api) BadRequest(message string) *echo.HTTPError {
	return &echo.HTTPError{Code: http.StatusInternalServerError, Message: message}
}
