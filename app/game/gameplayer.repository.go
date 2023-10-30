package gor

import (
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Gameplayer, error)
	GetViewById(conn *gorm.DB, id string) (model.GameplayerView, error)
	Create(conn *gorm.DB, data model.Gameplayer) error
	Update(conn *gorm.DB, data model.Gameplayer) error
	Delete(conn *gorm.DB, data model.Gameplayer) error
	Page(conn *gorm.DB, req *request.PageGameplayer) ([]model.GameplayerView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Gameplayer, error) {
	var err error
	var data model.Gameplayer

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GameplayerView, error) {
	var err error
	var data model.GameplayerView

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Gameplayer) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Gameplayer) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Gameplayer) error {
	now := time.Now()
	data.DeleteDt = &now

	return conn.Save(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageGameplayer) ([]model.GameplayerView, int64, error) {
	var err error
	var data []model.GameplayerView
	var count int64

	err = conn.Model(&data).
		Where("gor_id LIKE ?", "%"+req.GorID+"%").
		Where("player_id LIKE ?", "%"+req.PlayerID+"%").
		Where("delete_dt IS NULL").
		Count(&count).Error
	if err != nil {
		return data, count, err
	}

	err = conn.
		Where("gor_id LIKE ?", "%"+req.GorID+"%").
		Where("player_id LIKE ?", "%"+req.PlayerID+"%").
		Where("delete_dt IS NULL").
		Offset((req.GetPage() - 1) * req.GetLimit()).
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
