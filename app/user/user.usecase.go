package user

import (
	"errors"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/cryption"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/utils"
)

type Usecase interface {
	GetById(id string) (model.UserView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateUser) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateUser) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageUser) ([]model.UserView, int64, error)
	ChangePassword(loginUser jwt.UserLogin, req *request.ChangePassword) error
}

type usecaseUser struct {
	repo Repository
}

func (u usecaseUser) Create(loginUser jwt.UserLogin, req *request.CreateUser) error {
	var err error
	var data model.User

	password, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return errors.New("failed to encrypt")
	}

	data = model.User{
		Role:        constant.RoleUser,
		Email:       req.Email,
		Username:    req.Username,
		NoHp:        utils.FormatPhoneTo62(req.NoHp),
		Fullname:    req.Fullname,
		Passwd:      password,
		PassVersion: 1,
		IsActive:    true,
		PhotoID:     "",
		LastLoginDt: nil,
		CreateBy:    loginUser.UserID,
		UpdateBy:    loginUser.UserID,
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.Create(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseUser) GetById(id string) (model.UserView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseUser) Update(loginUser jwt.UserLogin, id string, req *request.UpdateUser) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.Fullname = req.Fullname
	data.Email = req.Email
	data.Username = req.Username
	data.NoHp = utils.FormatPhoneTo62(req.NoHp)
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

func (u usecaseUser) Delete(loginUser jwt.UserLogin, id string) error {
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

func (u usecaseUser) Page(req *request.PageUser) ([]model.UserView, int64, error) {
	var err error
	var data []model.UserView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (u usecaseUser) ChangePassword(loginUser jwt.UserLogin, req *request.ChangePassword) error {
	var err error
	var data model.User

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err = u.repo.GetById(conn, loginUser.UserID)
	if err != nil {
		return err
	}

	password, err := cryption.EncryptAES64(req.Passwd)
	if err != nil {
		return errors.New("failed to encrypt")
	}

	if password != data.Passwd {
		return errors.New("password not match")
	}

	data.PassVersion++
	data.Passwd = password
	data.UpdateBy = loginUser.UserID

	tx := conn.Begin()

	err = u.repo.Update(tx, data)
	if err != nil {
		return err
	}

	return err
}

func NewUserUsecase(repo Repository) Usecase {
	return usecaseUser{
		repo: repo,
	}
}
