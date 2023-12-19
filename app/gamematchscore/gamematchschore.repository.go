package gamematchscore

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Gamematchscore, error)
	GetViewById(conn *gorm.DB, id string) (model.GamematchscoreView, error)
	Create(conn *gorm.DB, data model.Gamematchscore) error
	Page(conn *gorm.DB, req *request.PageGamematchscore) ([]model.GamematchscoreView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Gamematchscore, error) {
	var err error
	var data model.Gamematchscore

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GamematchscoreView, error) {
	var err error
	var data model.GamematchscoreView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Gamematchscore) error {
	return conn.Create(&data).Error
}

//func (r repository) Update(conn *gorm.DB, data model.Gamematchscore) error {
//	return conn.Save(&data).Error
//}

//func (r repository) Delete(conn *gorm.DB, data model.Gamematchscore) error {
//	now := time.Now()
//	data.DeleteDt = &now
//
//	return conn.Save(&data).Error
//}

func (r repository) Page(conn *gorm.DB, req *request.PageGamematchscore) ([]model.GamematchscoreView, int64, error) {
	var err error
	var data []model.GamematchscoreView
	var count int64

	query := conn.Model(&data).
		Where("game_id LIKE ?", "%"+req.GameID+"%").
		Where("gamematch_id LIKE ?", "%"+req.GamematchID+"%").
		Where("LOWER(game_name) LIKE LOWER(?)", "%"+req.GameName+"%").
		Where("LOWER(gamematch_name) LIKE LOWER(?)", "%"+req.GamematchName+"%")

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

func NewRepository() Repository {
	return repository{}
}
