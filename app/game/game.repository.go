package game

import (
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Game, error)
	GetViewById(conn *gorm.DB, id string) (model.GameView, error)
	Create(conn *gorm.DB, data model.Game) error
	Update(conn *gorm.DB, data model.Game) error
	Delete(conn *gorm.DB, data model.Game) error
	Page(conn *gorm.DB, req *request.PageGame) ([]model.GameView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Game, error) {
	var err error
	var data model.Game

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GameView, error) {
	var err error
	var data model.GameView

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Game) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Game) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Game) error {
	now := time.Now()
	data.DeleteDt = &now

	return conn.Save(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageGame) ([]model.GameView, int64, error) {
	var err error
	var data []model.GameView
	var count int64

	err = conn.Model(&data).
		Where("gor_id LIKE ?", "%"+req.GorID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Where("delete_dt IS NULL").
		Count(&count).Error
	if err != nil {
		return data, count, err
	}

	err = conn.
		Where("gor_id LIKE ?", "%"+req.GorID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
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
