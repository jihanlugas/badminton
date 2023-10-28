package user

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

func UserHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// GetById
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [get]
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

	res := response.User(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Create
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateUser true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateUser)
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

// Update
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param req body request.UpdateUser true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [put]
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

	req := new(request.UpdateUser)
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
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/{id} [delete]
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
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageUser false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/page [get]
func (h Handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageUser)
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
