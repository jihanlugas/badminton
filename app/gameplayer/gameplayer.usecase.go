package gameplayer

import (
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
)

type Usecase interface {
	GetById(id string) (model.GameplayerView, error)
	Create(loginUser jwt.UserLogin, req *request.CreateGameplayer) error
	CreateBulk(loginUser jwt.UserLogin, req *request.CreateBulkGameplayer) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateGameplayer) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageGameplayer) ([]model.GameplayerView, int64, error)
	PageRank(req *request.PageRankGameplayer) ([]model.GameplayerRangking, int64, error)
}

type usecaseGameplayer struct {
	repo Repository
}

func (u usecaseGameplayer) GetById(id string) (model.GameplayerView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)

	return data, err
}

func (u usecaseGameplayer) Create(loginUser jwt.UserLogin, req *request.CreateGameplayer) error {
	var err error
	var data model.Gameplayer

	data = model.Gameplayer{
		GameID:     req.GameID,
		PlayerID:   req.PlayerID,
		NormalGame: 0,
		RubberGame: 0,
		Ball:       0,
		IsPay:      false,
		CreateBy:   loginUser.UserID,
		UpdateBy:   loginUser.UserID,
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

func (u usecaseGameplayer) CreateBulk(loginUser jwt.UserLogin, req *request.CreateBulkGameplayer) error {
	var err error
	var data []model.Gameplayer

	for _, playerID := range req.ListPlayerID {
		data = append(data, model.Gameplayer{
			GameID:     req.GameID,
			PlayerID:   playerID,
			NormalGame: 0,
			RubberGame: 0,
			Ball:       0,
			IsPay:      false,
			CreateBy:   loginUser.UserID,
			UpdateBy:   loginUser.UserID,
		})
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	err = u.repo.ListCreate(tx, data)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseGameplayer) Update(loginUser jwt.UserLogin, id string, req *request.UpdateGameplayer) error {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	data.GameID = req.GameID
	data.PlayerID = req.PlayerID
	data.NormalGame = req.NormalGame
	data.RubberGame = req.RubberGame
	data.Ball = req.Ball
	data.IsPay = req.IsPay
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

func (u usecaseGameplayer) Delete(loginUser jwt.UserLogin, id string) error {
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

func (u usecaseGameplayer) Page(req *request.PageGameplayer) ([]model.GameplayerView, int64, error) {
	var err error
	var data []model.GameplayerView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (u usecaseGameplayer) PageRank(req *request.PageRankGameplayer) ([]model.GameplayerRangking, int64, error) {
	var err error
	var data []model.GameplayerRangking
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.PageRank(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewGameplayerUsecase(repo Repository) Usecase {
	return usecaseGameplayer{
		repo: repo,
	}
}
