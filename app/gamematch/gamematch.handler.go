package gamematch

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	usecase Usecase
}

func GamematchHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// Page
// @Tags Game
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageGamematch false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gamematch/page [get]
func (h Handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageGamematch)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	data, count, err := h.usecase.Page(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.PayloadPagination(req, data, count)).SendJSON(c)
}

// Create
// @Tags Gamematch
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateGamematch true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gamematch [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateGamematch)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Create(loginUser, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}
