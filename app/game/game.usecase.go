package game

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.GameView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateGame) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateGame) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageGame) ([]model.GameView, int64, error)
}

type usecaseGame struct {
	repo Repository
}

func (u usecaseGame) GetById(id string) (model.GameView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseGame) Create(loginUser jwt.UserLogin, req *request.CreateGame) error {
	var err error
	var data model.Game

	data = model.Game{
		GorID:           req.GorID,
		Name:            req.Name,
		Description:     req.Description,
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

func (u usecaseGame) Update(loginUser jwt.UserLogin, id string, req *request.UpdateGame) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.GorID = req.GorID
	data.Name = req.Name
	data.Description = req.Description
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

func (u usecaseGame) Delete(loginUser jwt.UserLogin, id string) error {
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

func (u usecaseGame) Page(req *request.PageGame) ([]model.GameView, int64, error) {
	var err error
	var data []model.GameView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewGameUsecase(repo Repository) Usecase {
	return usecaseGame{
		repo: repo,
	}
}