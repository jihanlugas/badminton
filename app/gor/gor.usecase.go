package gor

import (
	"errors"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.GorView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateGor) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateGor) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageGor) ([]model.GorView, int64, error)
}

type usecaseGor struct {
	repo Repository
}

func (u usecaseGor) GetById(id string) (model.GorView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseGor) Create(loginUser jwt.UserLogin, req *request.CreateGor) error {
	var err error
	var data model.Gor

	if loginUser.Role != constant.RoleAdmin {
		if req.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	data = model.Gor{
		CompanyID:       req.CompanyID,
		Name:            req.Name,
		Description:     req.Description,
		Address:         req.Address,
		NormalGamePrice: req.NormalGamePrice,
		RubberGamePrice: req.RubberGamePrice,
		BallPrice:       req.BallPrice,
		CreateBy:        loginUser.UserID,
		UpdateBy:        loginUser.UserID,
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

func (u usecaseGor) Update(loginUser jwt.UserLogin, id string, req *request.UpdateGor) error {
	var err error

	if loginUser.Role != constant.RoleAdmin {
		if req.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.CompanyID = req.CompanyID
	data.Name = req.Name
	data.Description = req.Description
	data.Address = req.Address
	data.NormalGamePrice = req.NormalGamePrice
	data.RubberGamePrice = req.RubberGamePrice
	data.BallPrice = req.BallPrice
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

func (u usecaseGor) Delete(loginUser jwt.UserLogin, id string) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	if loginUser.Role != constant.RoleAdmin {
		if data.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
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

func (u usecaseGor) Page(req *request.PageGor) ([]model.GorView, int64, error) {
	var err error
	var data []model.GorView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewGorUsecase(repo Repository) Usecase {
	return usecaseGor{
		repo: repo,
	}
}
