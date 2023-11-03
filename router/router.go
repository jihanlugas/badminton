package router

import (
	"encoding/json"
	"fmt"
	"github.com/jihanlugas/badminton/app/app"
	"github.com/jihanlugas/badminton/app/authentication"
	"github.com/jihanlugas/badminton/app/company"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/app/user"
	"github.com/jihanlugas/badminton/app/usercompany"
	"github.com/jihanlugas/badminton/config"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/db"
	_ "github.com/jihanlugas/badminton/docs"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/response"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Init() *echo.Echo {
	router := websiteRouter()

	authenticationRepo := authentication.NewRepository()
	userRepo := user.NewRepository()
	companyRepo := company.NewRepository()
	usercompanyRepo := usercompany.NewRepository()
	//itemRepo := item.NewItemRepository()

	authenticationUsecase := authentication.NewAuthenticationUsecase(authenticationRepo, userRepo, companyRepo, usercompanyRepo)
	userUsecase := user.NewUserUsecase(userRepo)
	companyUsecase := company.NewCompanyUsecase(companyRepo, userRepo, usercompanyRepo)

	authenticationHandler := authentication.NewAuthenticationHandler(authenticationUsecase)
	userHandler := user.UserHandler(userUsecase)
	companyHandler := company.CompanyHandler(companyUsecase)

	//router.Use(logMiddleware)
	//router.Use(loggerMiddleware)

	router.GET("/swg/*", echoSwagger.WrapHandler)
	router.GET("/", app.Ping)

	router.POST("/sign-in", authenticationHandler.SignIn)
	//router.GET("/sign-out", authenticationHandler.SignOut)
	//router.POST("/sign-up", authenticationHandler.SignUp)
	router.GET("/refresh-token", authenticationHandler.RefreshToken, checkTokenAdminMiddleware)

	userRouter := router.Group("/user")
	userRouter.GET("/:id", userHandler.GetById)
	userRouter.POST("", userHandler.Create, checkTokenAdminMiddleware)
	userRouter.PUT("/:id", userHandler.Update, checkTokenAdminMiddleware)
	userRouter.DELETE("/:id", userHandler.Delete, checkTokenAdminMiddleware)
	userRouter.GET("/page", userHandler.Page, checkTokenAdminMiddleware)
	userRouter.POST("/change-password", userHandler.ChangePassword, checkTokenMiddleware)

	companyRouter := router.Group("/company")
	companyRouter.GET("/:id", companyHandler.GetById)
	companyRouter.POST("", companyHandler.Create, checkTokenAdminMiddleware)
	companyRouter.PUT("/:id", companyHandler.Update, checkTokenAdminMiddleware)
	companyRouter.DELETE("/:id", companyHandler.Delete, checkTokenAdminMiddleware)
	companyRouter.GET("/page", companyHandler.Page, checkTokenAdminMiddleware)

	return router

}

func httpErrorHandler(err error, c echo.Context) {
	var errorResponse *response.Response
	code := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		// Handle pada saat URL yang di request tidak ada. atau ada kesalahan server.
		code = e.Code
		errorResponse = &response.Response{
			Status:  false,
			Message: fmt.Sprintf("%v", e.Message),
			Payload: map[string]interface{}{},
			Code:    code,
		}
	case *response.Response:
		errorResponse = e
	default:
		// Handle error dari panic
		code = http.StatusInternalServerError
		if config.Debug {
			errorResponse = &response.Response{
				Status:  false,
				Message: err.Error(),
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		} else {
			errorResponse = &response.Response{
				Status:  false,
				Message: "Internal server error",
				Payload: map[string]interface{}{},
				Code:    http.StatusInternalServerError,
			}
		}
	}

	js, err := json.Marshal(errorResponse)
	if err == nil {
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, js)
	} else {
		b := []byte("{error: true, message: \"unresolved error\"}")
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
	}
}

func checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		userLogin, err := jwt.ExtractClaims(c.Request().Header.Get(config.HeaderAuthName))
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, err.Error(), response.Payload{}).SendJSON(c)
		}

		conn, closeConn := db.GetConnection()
		defer closeConn()

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
		}

		if user.PassVersion != userLogin.PassVersion {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired~", response.Payload{}).SendJSON(c)
		}

		c.Set(constant.TokenUserContext, userLogin)
		return next(c)
	}
}

func checkTokenAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		userLogin, err := jwt.ExtractClaims(c.Request().Header.Get(config.HeaderAuthName))
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, err.Error(), response.Payload{}).SendJSON(c)
		}

		conn, closeConn := db.GetConnection()
		defer closeConn()

		if userLogin.Role != constant.RoleAdmin {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired.", response.Payload{}).SendJSON(c)
		}

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
		}

		if user.PassVersion != userLogin.PassVersion {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired~", response.Payload{}).SendJSON(c)
		}

		c.Set(constant.TokenUserContext, userLogin)
		return next(c)
	}
}

//func logMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		body, _ := io.ReadAll(c.Request().Body)
//		c.Set(constant.Request, string(body))
//		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
//
//		return next(c)
//	}
//}

//func loggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		var err error
//		body, _ := io.ReadAll(c.Request().Body)
//		c.Set(constant.Request, string(body))
//		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
//
//		// Call next handler
//		if err := next(c); err != nil {
//			c.Error(err)
//		}
//
//		res := ""
//		if c.Get(constant.Response) != nil {
//			res = string(c.Get(constant.Response).([]byte))
//		}
//
//		loginUserString := ""
//		loginUser, err := user.GetUserLoginInfo(c)
//		if err == nil {
//			loginUserByte, _ := json.Marshal(loginUser)
//			loginUserString = string(loginUserByte)
//		}
//
//		logData := model.Log{
//			ClientIP:  c.Request().RemoteAddr,
//			Method:    c.Request().Method,
//			Path:      c.Request().URL.String(),
//			Code:      c.Response().Status,
//			Loginuser: loginUserString,
//			Request:   string(body),
//			Response:  res,
//		}
//
//		go log.AddLog(logData)
//
//		return nil
//	}
//}
