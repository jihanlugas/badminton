package gameplayer

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Gameplayer, error)
	GetByGameIdPlayerId(conn *gorm.DB, gameId, playerId string) (model.Gameplayer, error)
	GetViewById(conn *gorm.DB, id string) (model.GameplayerView, error)
	Create(conn *gorm.DB, data model.Gameplayer) error
	ListCreate(conn *gorm.DB, data []model.Gameplayer) error
	Update(conn *gorm.DB, data model.Gameplayer) error
	Delete(conn *gorm.DB, data model.Gameplayer) error
	Page(conn *gorm.DB, req *request.PageGameplayer) ([]model.GameplayerView, int64, error)
	PageRank(conn *gorm.DB, req *request.PageRankGameplayer) ([]model.GameplayerRangking, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Gameplayer, error) {
	var err error
	var data model.Gameplayer

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetByGameIdPlayerId(conn *gorm.DB, gameId, playerId string) (model.Gameplayer, error) {
	var err error
	var data model.Gameplayer

	err = conn.Where("game_id = ? AND player_id = ?", gameId, playerId).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GameplayerView, error) {
	var err error
	var data model.GameplayerView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Gameplayer) error {
	return conn.Create(&data).Error
}

func (r repository) ListCreate(conn *gorm.DB, data []model.Gameplayer) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Gameplayer) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Gameplayer) error {
	return conn.Delete(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageGameplayer) ([]model.GameplayerView, int64, error) {
	var err error
	var data []model.GameplayerView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(player_name) LIKE LOWER(?)", "%"+req.PlayerName+"%").
		Where("LOWER(game_name) LIKE LOWER(?)", "%"+req.GameName+"%").
		Where("game_id LIKE ?", "%"+req.GameID+"%").
		Where("player_id LIKE ?", "%"+req.PlayerID+"%")

	if req.Gender != "" {
		query = query.Where("gender = ?", req.Gender)
	}

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "create_dt", "asc"))
	}
	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (r repository) PageRank(conn *gorm.DB, req *request.PageRankGameplayer) ([]model.GameplayerRangking, int64, error) {
	var err error
	var data []model.GameplayerRangking
	var count int64

	query := conn.Model(&data).
		Select("player_id, player_name, gender,sum(normal_game) as normal_game,sum(rubber_game) as rubber_game,sum(normal_game + rubber_game) as game,sum(ball) as ball,sum(point) as point, RANK () OVER (ORDER BY sum(point) DESC) rank ").
		Where("is_finish = ?", true)

	if req.Gender != "" {
		query = query.Where("gender = ?", req.Gender)
	}

	query = query.Group("player_id, player_name, gender")

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	query = query.Order(fmt.Sprintf("%s %s", "rank", "asc"))

	err = query.Offset((req.GetPage() - 1) * req.GetLimit()).
		Limit(req.GetLimit()).
		Find(&data).Error
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func NewRepository() Repository {
	return repository{}
}
