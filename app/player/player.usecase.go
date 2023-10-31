package player

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.PlayerView, error)
	Create(loginUser jwt.UserLogin, req *request.CreatePlayer) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdatePlayer) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PagePlayer) ([]model.PlayerView, int64, error)
}

type usecasePlayer struct {
	repo Repository
}

func (u usecasePlayer) GetById(id string) (model.PlayerView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecasePlayer) Create(loginUser jwt.UserLogin, req *request.CreatePlayer) error {
	var err error
	var data model.Player

	data = model.Player{
		CompanyID: req.CompanyID,
		Name:      req.Name,
		Email:     req.Email,
		NoHp:      req.NoHp,
		Address:   req.Address,
		IsActive:  req.IsActive,
		CreateBy:  loginUser.UserID,
		UpdateBy:  loginUser.UserID,
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

func (u usecasePlayer) Update(loginUser jwt.UserLogin, id string, req *request.UpdatePlayer) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.CompanyID = req.CompanyID
	data.Name = req.Name
	data.Email = req.Email
	data.NoHp = req.NoHp
	data.Address = req.Address
	data.IsActive = req.IsActive
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

func (u usecasePlayer) Delete(loginUser jwt.UserLogin, id string) error {
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

func (u usecasePlayer) Page(req *request.PagePlayer) ([]model.PlayerView, int64, error) {
	var err error
	var data []model.PlayerView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewPlayerUsecase(repo Repository) Usecase {
	return usecasePlayer{
		repo: repo,
	}
}
