package transaction

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

func TransactionHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// GetById
// @Tags Transaction
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /transaction/{id} [get]
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

	res := response.Transaction(data)

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}

// Create
// @Tags Transaction
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreateTransaction true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /transaction [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	req := new(request.CreateTransaction)
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

// Page
// @Tags Transaction
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req query request.PageTransaction false "query string"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /transaction/page [get]
func (h Handler) Page(c echo.Context) error {
	var err error

	req := new(request.PageTransaction)
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
