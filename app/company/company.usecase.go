package company

import (
	"errors"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/app/user"
	"github.com/jihanlugas/badminton/app/usercompany"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/cryption"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/utils"
)

type Usecase interface {
	GetById(id string) (model.CompanyView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateCompany) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateCompany) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageCompany) ([]model.CompanyView, int64, error)
}

type usecaseCompany struct {
	repo            Repository
	userRepo        user.Repository
	usercompanyRepo usercompany.Repository
}

func (u usecaseCompany) GetById(id string) (model.CompanyView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseCompany) Create(loginUser jwt.UserLogin, req *request.CreateCompany) error {
	var err error
	var companyData model.Company
	var userData model.User
	var usercompanyData model.Usercompany

	companyData = model.Company{
		ID:          utils.GetUniqueID(),
		Name:        req.Name,
		Description: req.Description,
		Balance:     req.Balance,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	password, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return errors.New("failed to encrypt")
	}

	userData = model.User{
		ID:          utils.GetUniqueID(),
		Role:        constant.RoleUser,
		Email:       req.Email,
		Username:    req.Username,
		NoHp:        req.NoHp,
		Fullname:    req.Fullname,
		Passwd:      password,
		PassVersion: 1,
		IsActive:    true,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	usercompanyData = model.Usercompany{
		UserID:           userData.ID,
		CompanyID:        companyData.ID,
		IsDefaultCompany: true,
		IsCreator:        true,
		CreateBy:         loginUser.UserID,
		UpdateBy:         loginUser.UserID,
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.Create(tx, companyData)
	if err != nil {
		return err
	}

	err = u.userRepo.Create(tx, userData)
	if err != nil {
		return err
	}

	err = u.usercompanyRepo.Create(tx, usercompanyData)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseCompany) Update(loginUser jwt.UserLogin, id string, req *request.UpdateCompany) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.Name = req.Name
	data.Description = req.Description
	data.Balance = req.Balance
	data.UpdateBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Update(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseCompany) Delete(loginUser jwt.UserLogin, id string) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.DeleteBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Delete(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseCompany) Page(req *request.PageCompany) ([]model.CompanyView, int64, error) {
	var err error
	var data []model.CompanyView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewCompanyUsecase(repo Repository, userRepo user.Repository, usercompanyRepo usercompany.Repository) Usecase {
	return usecaseCompany{
		repo:            repo,
		userRepo:        userRepo,
		usercompanyRepo: usercompanyRepo,
	}
}
