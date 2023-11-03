package transaction

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.TransactionView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateTransaction) error
	Page(req *request.PageTransaction) ([]model.TransactionView, int64, error)
}

type usecaseTransaction struct {
	repo Repository
}

func (u usecaseTransaction) GetById(id string) (model.TransactionView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseTransaction) Create(loginUser jwt.UserLogin, req *request.CreateTransaction) error {
	var err error
	var data model.Transaction

	data = model.Transaction{
		CompanyID: req.CompanyID,
		Name:      req.Name,
		IsDebit:   req.IsDebit,
		Price:     req.Price,
		CreateBy:  loginUser.UserID,
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

func (u usecaseTransaction) Page(req *request.PageTransaction) ([]model.TransactionView, int64, error) {
	var err error
	var data []model.TransactionView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewTransactionUsecase(repo Repository) Usecase {
	return usecaseTransaction{
		repo: repo,
	}
}
