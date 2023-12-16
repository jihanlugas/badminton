package gamematch

import (
	"errors"
	"github.com/jihanlugas/badminton/app/gamematchscore"
	"github.com/jihanlugas/badminton/app/gamematchteam"
	"github.com/jihanlugas/badminton/app/gamematchteamplayer"
	"github.com/jihanlugas/badminton/app/jwt"
	"github.com/jihanlugas/badminton/constant"
	"github.com/jihanlugas/badminton/db"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"github.com/jihanlugas/badminton/utils"
)

type Usecase interface {
	Create(loginUser jwt.UserLogin, req *request.CreateGamematch) error
}

type usecaseGamematch struct {
	repo                    Repository
	gamematchscoreRepo      gamematchscore.Repository
	gamematchteamRepo       gamematchteam.Repository
	gamematchteamplayerRepo gamematchteamplayer.Repository
}

func (u usecaseGamematch) Create(loginUser jwt.UserLogin, req *request.CreateGamematch) error {
	var err error
	var gamematch model.Gamematch
	var gamematchscore model.Gamematchscore
	var gamematchteam model.Gamematchteam
	var gamematchteamplayer model.Gamematchteamplayer

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

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return err
}

func NewGamematchUsecase(repo Repository, gamematchscoreRepo gamematchscore.Repository, gamematchteamRepo gamematchteam.Repository, gamematchteamplayerRepo gamematchteamplayer.Repository) Usecase {
	return usecaseGamematch{
		repo:                    repo,
		gamematchscoreRepo:      gamematchscoreRepo,
		gamematchteamRepo:       gamematchteamRepo,
		gamematchteamplayerRepo: gamematchteamplayerRepo,
	}
}
