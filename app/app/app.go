package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func ErrorInternal(c echo.Context, err error) {
	//log.System.Error().Err(err).Str("Host", c.Request().Host).Str("Path", c.Path()).Send()
	panic(err)
}

// Ping godoc
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {string} interface{} "ようこそ、美しい世界へ"
// @Router       / [get]
func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "ようこそ、美しい世界へ")
}
