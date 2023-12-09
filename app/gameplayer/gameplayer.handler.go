package gameplayer

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

func GameplayerHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// GetById
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer/{id} [get]
func (h Handler) GetById(c echo.Context) error {
	var err error

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	data, err := h.usecase.GetById(id)
	if err != nil {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	res := response.Gameplayer(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Create
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateGameplayer true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateGameplayer)
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

// CreateBulk
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateBulkGameplayer true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer/bulk [post]
func (h Handler) CreateBulk(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateBulkGameplayer)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.CreateBulk(loginUser, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Update
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param req body request.UpdateGameplayer true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer/{id} [put]
func (h Handler) Update(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	req := new(request.UpdateGameplayer)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	if err = c.Validate(req); err != nil {
		return response.Error(http.StatusBadRequest, "error validation", response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Update(loginUser, id, req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Delete
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer/{id} [delete]
func (h Handler) Delete(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	id := c.Param("id")
	if id == "" {
		return response.Error(http.StatusBadRequest, "data not found", response.Payload{}).SendJSON(c)
	}

	err = h.usecase.Delete(loginUser, id)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{}).SendJSON(c)
}

// Page
// @Tags Gameplayer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageGameplayer false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /gameplayer/page [get]
func (h Handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageGameplayer)
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
