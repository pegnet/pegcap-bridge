package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ParamsRichList struct {
	Asset string `json:"asset,omitempty"`
	Count int    `json:"count"`
}
type ResultRichList struct {
	Address string `json:"address"`
	Amount  uint64 `json:"amount"`
	Equiv   uint64 `json:"pusd"`
}
type ResultGlobalRichList struct {
	Address string `json:"address"`
	Equiv   uint64 `json:"pusd"`
}

func (a *Api) _getRichList(asset string) ([]ResultRichList, error) {
	var res []ResultRichList
	err := a.Cli.Request("get-rich-list", ParamsRichList{Asset: asset, Count: 100}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (a *Api) _getGlobalRichList() ([]ResultGlobalRichList, error) {
	var res []ResultGlobalRichList
	err := a.Cli.Request("get-global-rich-list", ParamsRichList{Count: 100}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (api *Api) GlobalRichList(c echo.Context) error {
	h, err := api._getGlobalRichList()
	if err != nil {
		return api.BadRequest("unable to contact endpoint: " + err.Error())
	}
	js, _ := json.Marshal(h)
	return c.JSONBlob(http.StatusOK, js)
}

func (api *Api) RichList(c echo.Context) error {
	ass, ok := api.VerifyAsset(c)
	if !ok {
		return api.BadRequest("invalid asset specified")
	}

	h, err := api._getRichList(ass)
	if err != nil {
		return api.BadRequest("unable to contact endpoint: " + err.Error())
	}
	js, _ := json.Marshal(h)
	return c.JSONBlob(http.StatusOK, js)
}
