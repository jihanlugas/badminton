package game

import (
	"errors"
	"github.com/jihanlugas/badminton/app/company"
	"github.com/jihanlugas/badminton/app/gamematch"
	"github.com/jihanlugas/badminton/app/gamematchscore"
	"github.com/jihanlugas/badminton/app/gamematchteam"
	"github.com/jihanlugas/badminton/app/gamematchteamplayer"
	"github.com/jihanlugas/badminton/app/gameplayer"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/app/transaction"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/response"
)

type Usecase interface {
	GetById(id string) (model.GameView, error)
	GetByIdDetail(id string) (response.GameDetail, error)
	Create(loginUser jwt.UserLogin, req *request.CreateGame) error
	Update(loginUser jwt.UserLogin, id string, req *request.UpdateGame) error
	Delete(loginUser jwt.UserLogin, id string) error
	Page(req *request.PageGame) ([]model.GameView, int64, error)
	FinishGame(id string, loginUser jwt.UserLogin, req *request.FinishGame) error
}

type usecaseGame struct {
	repo                    Repository
	companyRepo             company.Repository
	gamematchRepo           gamematch.Repository
	gameplayerRepo          gameplayer.Repository
	gamematchscoreRepo      gamematchscore.Repository
	gamematchteamRepo       gamematchteam.Repository
	gamematchteamplayerRepo gamematchteamplayer.Repository
	transactionRepo         transaction.Repository
}

func (u usecaseGame) GetById(id string) (model.GameView, error) {
	var err error

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)
	if err != nil {
		return data, err
	}

	return data, err
}

func (u usecaseGame) GetByIdDetail(id string) (response.GameDetail, error) {
	var err error
	var res response.GameDetail
	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, err := u.repo.GetViewById(conn, id)
	if err != nil {
		return res, err
	}

	gamematches, _, err := u.gamematchRepo.Page(conn, &request.PageGamematch{
		GameID:    id,
		CompanyID: data.CompanyID,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return res, err
	}

	gameplayers, _, err := u.gameplayerRepo.Page(conn, &request.PageGameplayer{
		GameID: id,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return res, err
	}

	gamematchscores, _, err := u.gamematchscoreRepo.Page(conn, &request.PageGamematchscore{
		GameID: id,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return res, err
	}

	gamematchteams, _, err := u.gamematchteamRepo.Page(conn, &request.PageGamematchteam{
		GameID: id,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return res, err
	}

	gamematchteamplayers, _, err := u.gamematchteamplayerRepo.Page(conn, &request.PageGamematchteamplayer{
		GameID: id,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return res, err
	}

	res.Game = data
	res.Gamematches = gamematches
	res.Gameplayers = gameplayers
	res.Gamematchscores = gamematchscores
	res.Gamematchteams = gamematchteams
	res.Gamematchteamplayers = gamematchteamplayers

	return res, err
}

func (u usecaseGame) Create(loginUser jwt.UserLogin, req *request.CreateGame) error {
	var err error
	var data model.Game

	if loginUser.Role != constant.RoleAdmin {
		if req.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	data = model.Game{
		CompanyID:       req.CompanyID,
		GorID:           req.GorID,
		Name:            req.Name,
		Description:     req.Description,
		NormalGamePrice: req.NormalGamePrice,
		RubberGamePrice: req.RubberGamePrice,
		BallPrice:       req.BallPrice,
		GameDt:          req.GameDt,
		IsFinish:        false,
		ExpectedDebit:   0,
		Debit:           0,
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

	data.GorID = req.GorID
	data.Name = req.Name
	data.Description = req.Description
	data.NormalGamePrice = req.NormalGamePrice
	data.RubberGamePrice = req.RubberGamePrice
	data.BallPrice = req.BallPrice
	data.GameDt = req.GameDt
	data.IsFinish = req.IsFinish
	data.ExpectedDebit = req.ExpectedDebit
	data.Debit = req.Debit
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

func (u usecaseGame) FinishGame(id string, loginUser jwt.UserLogin, req *request.FinishGame) error {
	var err error
	var game model.Game
	var company model.Company
	var gameplayers []model.GameplayerView
	var transactions []model.Transaction
	var expectedDebit int64
	var balance int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	game, err = u.repo.GetById(conn, id)
	if err != nil {
		return err
	}

	if game.IsFinish {
		return errors.New("game already finished")
	}

	if loginUser.Role != constant.RoleAdmin {
		if game.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	company, err = u.companyRepo.GetById(conn, game.CompanyID)
	if err != nil {
		return err
	}

	gameplayers, _, err = u.gameplayerRepo.Page(conn, &request.PageGameplayer{
		GameID: id,
		Paging: request.Paging{
			Limit: 1000,
			Page:  1,
		},
	})
	if err != nil {
		return err
	}

	for _, data := range gameplayers {
		expectedDebit += (data.NormalGame * game.NormalGamePrice) + (data.RubberGame * game.RubberGamePrice) + (data.Ball * game.BallPrice)
	}

	for _, data := range req.Transactions {
		if data.CompanyID != game.CompanyID {
			return errors.New("permission denied")
		}
		if data.IsDebit {
			balance += data.Price
		} else {
			balance -= data.Price
		}
		newTransaction := model.Transaction{
			CompanyID: data.CompanyID,
			Name:      data.Name,
			IsDebit:   data.IsDebit,
			Price:     data.Price,
			CreateBy:  loginUser.UserID,
		}
		transactions = append(transactions, newTransaction)
	}

	tx := conn.Begin()

	err = u.transactionRepo.ListCreate(tx, transactions)
	if err != nil {
		return err
	}

	game.IsFinish = true
	game.ExpectedDebit = expectedDebit
	game.UpdateBy = loginUser.UserID
	err = u.repo.Update(tx, game)
	if err != nil {
		return err
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

func NewGameUsecase(repo Repository, companyRepo company.Repository, gamematchRepo gamematch.Repository, gameplayerRepo gameplayer.Repository, gamematchscoreRepo gamematchscore.Repository, gamematchteamRepo gamematchteam.Repository, gamematchteamplayerRepo gamematchteamplayer.Repository, transactionRepo transaction.Repository) Usecase {
	return usecaseGame{
		repo:                    repo,
		companyRepo:             companyRepo,
		gamematchRepo:           gamematchRepo,
		gameplayerRepo:          gameplayerRepo,
		gamematchscoreRepo:      gamematchscoreRepo,
		gamematchteamRepo:       gamematchteamRepo,
		gamematchteamplayerRepo: gamematchteamplayerRepo,
		transactionRepo:         transactionRepo,
	}
}
