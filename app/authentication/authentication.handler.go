package authentication

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthenticationHandler struct {
	usecase AuthenticationUsecase
}

func NewAuthenticationHandler(usecase AuthenticationUsecase) AuthenticationHandler {
	return AuthenticationHandler{
		usecase: usecase,
	}
}

// SignIn
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-in [post]
func (h AuthenticationHandler) SignIn(c echo.Context) error {
	var err error

	req := new(request.Signin)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	err = c.Validate(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.ValidationError(err)).SendJSON(c)
	}

	token, userLogin, err := h.usecase.SignIn(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{
		"token":     token,
		"userLogin": userLogin,
	}).SendJSON(c)
}

// SignOut Sign out user
// @Tags Authentication
// @Accept json
// @Produce json
// // @Param req body request.Signin true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-out [get]
func (h AuthenticationHandler) SignOut(c echo.Context) error {
	return nil
}

//// SignUp
//// @Tags Authentication
//// @Accept json
//// @Produce json
//// // @Param req body request.Signin true "json req body"
//// @Success      200  {object}	response.Response
//// @Failure      500  {object}  response.Response
//// @Router /sign-up [post]
//func (h AuthenticationHandler) SignUp(c echo.Context) error {
//	return nil
//}

// RefreshToken
// @Tags Authentication
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /refresh-token [get]
func (h AuthenticationHandler) RefreshToken(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	token, err := h.usecase.RefreshToken(loginUser)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	return response.Success(http.StatusOK, "success", response.Payload{
		"token": token,
	}).SendJSON(c)
}

// Init
// @Tags Authentication
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /refresh-token [get]
func (h AuthenticationHandler) Init(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	user, company, err := h.usecase.Init(loginUser)
	if err != nil {
		return response.Error(http.StatusBadRequest, err.Error(), response.Payload{}).SendJSON(c)
	}

	resUser := response.User(*user)
	resCompany := response.Company(*company)

	res := response.Init{
		User:    &resUser,
		Company: &resCompany,
	}

	return response.Success(http.StatusOK, "success", res).SendJSON(c)
}
