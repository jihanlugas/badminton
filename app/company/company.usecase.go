package company

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.CompanyView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateCompany) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateCompany) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageCompany) ([]model.CompanyView, int64, error)
}

type usecaseCompany struct {
	repo Repository
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
	var data model.Company

	data = model.Company{
		Name:        req.Name,
		Description: req.Description,
		Balance:     req.Balance,
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

func NewCompanyUsecase(repo Repository) Usecase {
	return usecaseCompany{
		repo: repo,
	}
}
