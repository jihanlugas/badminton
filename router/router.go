package router

import (
	"encoding/json"
	"fmt"
	"github.com/jihanlugas/badminton/app/app"
	"github.com/jihanlugas/badminton/app/authentication"
	"github.com/jihanlugas/badminton/app/company"
	"github.com/jihanlugas/badminton/app/game"
	"github.com/jihanlugas/badminton/app/gamematch"
	"github.com/jihanlugas/badminton/app/gamematchscore"
	"github.com/jihanlugas/badminton/app/gamematchteam"
	"github.com/jihanlugas/badminton/app/gamematchteamplayer"
	"github.com/jihanlugas/badminton/app/gameplayer"
	"github.com/jihanlugas/badminton/app/gor"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/app/player"
	"github.com/jihanlugas/badminton/app/transaction"
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
	transactionRepo := transaction.NewRepository()
	gorRepo := gor.NewRepository()
	gameRepo := game.NewRepository()
	playerRepo := player.NewRepository()
	gameplayerRepo := gameplayer.NewRepository()
	gamematchRepo := gamematch.NewRepository()
	gamematchscoreRepo := gamematchscore.NewRepository()
	gamematchteamRepo := gamematchteam.NewRepository()
	gamematchteamplayerRepo := gamematchteamplayer.NewRepository()

	authenticationUsecase := authentication.NewAuthenticationUsecase(authenticationRepo, userRepo, companyRepo, usercompanyRepo)
	userUsecase := user.NewUserUsecase(userRepo)
	companyUsecase := company.NewCompanyUsecase(companyRepo, userRepo, usercompanyRepo)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepo, companyRepo)
	gorUsecase := gor.NewGorUsecase(gorRepo)
	gameUsecase := game.NewGameUsecase(gameRepo, companyRepo, gamematchRepo, gameplayerRepo, gamematchscoreRepo, gamematchteamRepo, gamematchteamplayerRepo, transactionRepo)
	playerUsecase := player.NewPlayerUsecase(playerRepo)
	gameplayerUsecase := gameplayer.NewGameplayerUsecase(gameplayerRepo)
	gamematchUsecase := gamematch.NewGamematchUsecase(gamematchRepo, gameplayerRepo, gamematchscoreRepo, gamematchteamRepo, gamematchteamplayerRepo)

	authenticationHandler := authentication.NewAuthenticationHandler(authenticationUsecase)
	userHandler := user.UserHandler(userUsecase)
	companyHandler := company.CompanyHandler(companyUsecase)
	transactionHandler := transaction.TransactionHandler(transactionUsecase)
	gorHandler := gor.GorHandler(gorUsecase)
	gameHandler := game.GameHandler(gameUsecase)
	playerHandler := player.PlayerHandler(playerUsecase)
	gameplayerHandler := gameplayer.GameplayerHandler(gameplayerUsecase)
	gamematchHandler := gamematch.GamematchHandler(gamematchUsecase)

	//router.Use(logMiddleware)
	//router.Use(loggerMiddleware)

	router.GET("/swg/*", echoSwagger.WrapHandler)
	router.GET("/", app.Ping)

	router.POST("/sign-in", authenticationHandler.SignIn)
	router.GET("/sign-out", authenticationHandler.SignOut)
	router.POST("/identity/:provider", authenticationHandler.Identity)
	router.GET("/refresh-token", authenticationHandler.RefreshToken, checkTokenMiddleware)
	router.GET("/init", authenticationHandler.Init, checkTokenMiddleware)

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

	transactionRouter := router.Group("/transaction")
	transactionRouter.GET("/:id", transactionHandler.GetById)
	transactionRouter.POST("", transactionHandler.Create, checkTokenMiddleware)
	transactionRouter.GET("/page", transactionHandler.Page, checkTokenMiddleware)

	gorRouter := router.Group("/gor")
	gorRouter.GET("/:id", gorHandler.GetById)
	gorRouter.POST("", gorHandler.Create, checkTokenMiddleware)
	gorRouter.PUT("/:id", gorHandler.Update, checkTokenMiddleware)
	gorRouter.DELETE("/:id", gorHandler.Delete, checkTokenMiddleware)
	gorRouter.GET("/page", gorHandler.Page, checkTokenMiddleware)

	gameRouter := router.Group("/game")
	gameRouter.GET("/:id", gameHandler.GetById)
	gameRouter.GET("/:id/detail", gameHandler.GetByIdDetail)
	gameRouter.POST("/:id/finish", gameHandler.FinishGame, checkTokenMiddleware)
	gameRouter.POST("", gameHandler.Create, checkTokenMiddleware)
	gameRouter.PUT("/:id", gameHandler.Update, checkTokenMiddleware)
	gameRouter.DELETE("/:id", gameHandler.Delete, checkTokenMiddleware)
	gameRouter.GET("/page", gameHandler.Page, checkTokenMiddleware)

	playerRouter := router.Group("/player")
	playerRouter.GET("/:id", playerHandler.GetById)
	playerRouter.POST("", playerHandler.Create, checkTokenMiddleware)
	playerRouter.PUT("/:id", playerHandler.Update, checkTokenMiddleware)
	playerRouter.DELETE("/:id", playerHandler.Delete, checkTokenMiddleware)
	playerRouter.GET("/page", playerHandler.Page, checkTokenMiddleware)

	gameplayerRouter := router.Group("/gameplayer")
	gameplayerRouter.GET("/:id", gameplayerHandler.GetById)
	gameplayerRouter.POST("", gameplayerHandler.Create, checkTokenMiddleware)
	gameplayerRouter.POST("/bulk", gameplayerHandler.CreateBulk, checkTokenMiddleware)
	gameplayerRouter.PUT("/:id", gameplayerHandler.Update, checkTokenMiddleware)
	gameplayerRouter.DELETE("/:id", gameplayerHandler.Delete, checkTokenMiddleware)
	gameplayerRouter.GET("/page", gameplayerHandler.Page, checkTokenMiddleware)
	gameplayerRouter.GET("/page-rank", gameplayerHandler.PageRank, checkTokenMiddleware)

	gamematchRouter := router.Group("/gamematch")
	gamematchRouter.GET("/page", gamematchHandler.Page, checkTokenMiddleware)
	gamematchRouter.POST("/match-point", gamematchHandler.CreateMatchpoint, checkTokenMiddleware)
	gamematchRouter.POST("/match", gamematchHandler.CreateMatch, checkTokenMiddleware)

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

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, "Token Expired!", response.Payload{}).SendJSON(c)
		}

		if user.Role != constant.RoleAdmin {
			return response.ErrorForce(http.StatusUnauthorized, "permission denied.", response.Payload{}).SendJSON(c)
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
