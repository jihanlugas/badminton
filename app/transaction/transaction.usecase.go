package transaction

import (
	"errors"
	"github.com/jihanlugas/badminton/app/company"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/constant"
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
	repo        Repository
	companyRepo company.Repository
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
	var company model.Company
	var balance int64

	if loginUser.Role != constant.RoleAdmin {
		if req.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	company, err = u.companyRepo.GetById(conn, req.CompanyID)
	if err != nil {
		return err
	}

	tx := conn.Begin()

	data = model.Transaction{
		CompanyID: req.CompanyID,
		Name:      req.Name,
		IsDebit:   req.IsDebit,
		Price:     req.Price,
		CreateBy:  loginUser.UserID,
	}

	err = u.repo.Create(tx, data)
	if err != nil {
		return err
	}

	if data.IsDebit {
		balance += data.Price
	} else {
		balance -= data.Price
	}

	company.Balance = company.Balance + balance
	company.UpdateBy = loginUser.UserID
	err = u.companyRepo.Update(tx, company)
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

func NewTransactionUsecase(repo Repository, companyRepo company.Repository) Usecase {
	return usecaseTransaction{
		repo:        repo,
		companyRepo: companyRepo,
	}
}
