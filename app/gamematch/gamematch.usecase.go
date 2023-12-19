package gamematch

import (
	"errors"
	"github.com/jihanlugas/badminton/app/gamematchscore"
	"github.com/jihanlugas/badminton/app/gamematchteam"
	"github.com/jihanlugas/badminton/app/gamematchteamplayer"
	"github.com/jihanlugas/badminton/app/gameplayer"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/utils"
)

type Usecase interface {
	Create(loginUser jwt.UserLogin, req *request.CreateGamematch) error
	Page(req *request.PageGamematch) ([]model.GamematchView, int64, error)
}

type usecaseGamematch struct {
	repo                    Repository
	gameplayerRepo          gameplayer.Repository
	gamematchscoreRepo      gamematchscore.Repository
	gamematchteamRepo       gamematchteam.Repository
	gamematchteamplayerRepo gamematchteamplayer.Repository
}

func (u usecaseGamematch) Create(loginUser jwt.UserLogin, req *request.CreateGamematch) error {
	var err error
	var gamematch model.Gamematch
	var gameplayer model.Gameplayer
	var gamematchscore model.Gamematchscore
	var gamematchteam model.Gamematchteam
	var gamematchteamplayer model.Gamematchteamplayer
	var players []string

	if loginUser.Role != constant.RoleAdmin {
		if req.CompanyID != loginUser.CompanyID {
			return errors.New("permission denied")
		}
	}

	conn, closeConn := db.GetConnection()
	defer closeConn()

	tx := conn.Begin()

	leftTeamID := utils.GetUniqueID()
	rightTeamID := utils.GetUniqueID()

	gamematch = model.Gamematch{
		ID:             utils.GetUniqueID(),
		CompanyID:      req.CompanyID,
		GameID:         req.GameID,
		Name:           req.MatchName,
		LeftTeamID:     leftTeamID,
		RightTeamID:    rightTeamID,
		LeftTeamPoint:  req.LeftTeamPoint,
		RightTeamPoint: req.RightTeamPoint,
		IsRubber:       req.IsRubber,
		CreateBy:       loginUser.UserID,
	}

	err = u.repo.Create(tx, gamematch)
	if err != nil {
		return err
	}

	for index, data := range req.GameMatchScores {
		gamematchscore = model.Gamematchscore{
			GameID:         req.GameID,
			GamematchID:    gamematch.ID,
			Set:            int64(index + 1),
			LeftTeamScore:  data.LeftScore,
			RightTeamScore: data.RightScore,
			CreateBy:       loginUser.UserID,
		}

		err = u.gamematchscoreRepo.Create(tx, gamematchscore)
		if err != nil {
			return err
		}
	}

	for index, data := range req.GameMatchTeams {
		id := ""
		if index == 0 {
			id = leftTeamID
		} else {
			id = rightTeamID
		}
		gamematchteam = model.Gamematchteam{
			ID:          id,
			GameID:      req.GameID,
			GamematchID: gamematch.ID,
			Name:        data.Name,
			CreateBy:    loginUser.UserID,
		}
		err = u.gamematchteamRepo.Create(tx, gamematchteam)
		if err != nil {
			return err
		}

		for _, player := range data.GameMatchTeamPlayers {
			players = append(players, player.PlayerID)
			gamematchteamplayer = model.Gamematchteamplayer{
				GameID:          req.GameID,
				GamematchID:     gamematch.ID,
				GamematchteamID: gamematchteam.ID,
				PlayerID:        player.PlayerID,
				CreateBy:        loginUser.UserID,
			}
			err = u.gamematchteamplayerRepo.Create(tx, gamematchteamplayer)
			if err != nil {
				return err
			}
		}
	}

	for _, playerId := range players {
		gameplayer, err = u.gameplayerRepo.GetByGameIdPlayerId(conn, req.GameID, playerId)
		if err != nil {
			return err
		}
		if req.IsRubber {
			gameplayer.RubberGame = gameplayer.RubberGame + 1
		} else {
			gameplayer.NormalGame = gameplayer.NormalGame + 1
		}
		gameplayer.Ball = gameplayer.Ball + req.Ball
		gameplayer.UpdateBy = loginUser.UserID
		err = u.gameplayerRepo.Update(conn, gameplayer)
		if err != nil {
			return err
		}
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func (u usecaseGamematch) Page(req *request.PageGamematch) ([]model.GamematchView, int64, error) {
	var err error
	var data []model.GamematchView
	var count int64

	conn, closeConn := db.GetConnection()
	defer closeConn()

	data, count, err = u.repo.Page(conn, req)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewGamematchUsecase(repo Repository, gameplayerRepo gameplayer.Repository, gamematchscoreRepo gamematchscore.Repository, gamematchteamRepo gamematchteam.Repository, gamematchteamplayerRepo gamematchteamplayer.Repository) Usecase {
	return usecaseGamematch{
		repo:                    repo,
		gameplayerRepo:          gameplayerRepo,
		gamematchscoreRepo:      gamematchscoreRepo,
		gamematchteamRepo:       gamematchteamRepo,
		gamematchteamplayerRepo: gamematchteamplayerRepo,
	}
}
