package authentication

import (
	"errors"
	"github.com/jihanlugas/badminton/app/company"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/app/user"
	"github.com/jihanlugas/badminton/app/usercompany"
	"github.com/jihanlugas/badminton/config"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/cryption"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/utils"
	"time"
)

type AuthenticationUsecase interface {
	SignIn(req *request.Signin) (string, jwt.UserLogin, error)
	RefreshToken(loginUser jwt.UserLogin) (string, error)
	Init(loginUser jwt.UserLogin) (*model.UserView, *model.CompanyView, error)
}

type usecaseAuthentication struct {
	repo            Repository
	userRepo        user.Repository
	companyRepo     company.Repository
	usercompanyRepo usercompany.Repository
}

func (u usecaseAuthentication) SignIn(req *request.Signin) (string, jwt.UserLogin, error) {
	var err error
	var data model.User
	var companydata model.Company
	var usercompanydata model.Usercompany
	var userLogin jwt.UserLogin

	conn, closeConn := db.GetConnection()
	defer closeConn()

	if utils.IsValidEmail(req.Username) {
		data, err = u.userRepo.GetByEmail(conn, req.Username)
	} else {
		data, err = u.userRepo.GetByUsername(conn, req.Username)
	}

	if err != nil {
		return "", userLogin, err
	}

	err = cryption.CheckAES64(req.Passwd, data.Passwd)
	if err != nil {
		return "", userLogin, errors.New("invalid username or password")
	}

	if !data.IsActive {
		return "", userLogin, errors.New("user not active")
	}

	if data.Role != constant.RoleAdmin {
		usercompanydata, err = u.usercompanyRepo.GetCompanyDefaultByUserId(conn, data.ID)
		if err != nil {
			return "", userLogin, errors.New("usercompany not found : " + err.Error())
		}

		companydata, err = u.companyRepo.GetById(conn, usercompanydata.CompanyID)
		if err != nil {
			return "", userLogin, errors.New("company not found : " + err.Error())
		}
	}

	now := time.Now()
	tx := conn.Begin()

	data.LastLoginDt = &now
	data.UpdateBy = data.ID
	err = u.userRepo.Update(tx, data)
	if err != nil {
		return "", userLogin, err
	}

	err = tx.Commit().Error
	if err != nil {
		return "", userLogin, err
	}

	expiredAt := time.Now().Add(time.Hour * time.Duration(config.AuthTokenExpiredHour))

	userLogin.UserID = data.ID
	userLogin.Role = data.Role
	userLogin.PassVersion = data.PassVersion
	userLogin.CompanyID = companydata.ID
	userLogin.UsercompanyID = usercompanydata.ID
	token, err := jwt.CreateToken(userLogin, expiredAt)
	if err != nil {
		return "", userLogin, err
	}

	return token, userLogin, err
}

func (u usecaseAuthentication) RefreshToken(loginUser jwt.UserLogin) (string, error) {
	var err error

	expiredAt := time.Now().Add(time.Hour * time.Duration(config.AuthTokenExpiredHour))

	token, err := jwt.CreateToken(loginUser, expiredAt)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u usecaseAuthentication) Init(loginUser jwt.UserLogin) (*model.UserView, *model.CompanyView, error) {
	var err error
	var dataUser model.UserView
	var dataCompany model.CompanyView

	conn, closeConn := db.GetConnection()
	defer closeConn()

	dataUser, err = u.userRepo.GetViewById(conn, loginUser.UserID)
	if err != nil {
		return &dataUser, &dataCompany, err
	}

	if loginUser.Role != constant.RoleAdmin {
		dataCompany, err = u.companyRepo.GetViewById(conn, loginUser.CompanyID)
		if err != nil {
			return &dataUser, &dataCompany, err
		}
	}

	return &dataUser, &dataCompany, err
}

func NewAuthenticationUsecase(repo Repository, userRepo user.Repository, companyRepo company.Repository, usercompanyRepo usercompany.Repository) AuthenticationUsecase {
	return usecaseAuthentication{
		repo:            repo,
		userRepo:        userRepo,
		companyRepo:     companyRepo,
		usercompanyRepo: usercompanyRepo,
	}
}
