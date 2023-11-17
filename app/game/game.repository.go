package game

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
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

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GameView, error) {
	var err error
	var data model.GameView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Game) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Game) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Game) error {
	return conn.Delete(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageGame) ([]model.GameView, int64, error) {
	var err error
	var data []model.GameView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(gor_id) LIKE LOWER(?)", "%"+req.GorID+"%").
		Where("LOWER(name) LIKE LOWER(?)", "%"+req.Name+"%").
		Where("LOWER(description) LIKE LOWER(?)", "%"+req.Description+"%")

	err = query.Count(&count).Error
	if err != nil {
		return data, count, err
	}

	if req.SortField != "" {
		query = query.Order(fmt.Sprintf("%s %s", req.SortField, req.SortOrder))
	} else {
		query = query.Order(fmt.Sprintf("%s %s", "name", "asc"))
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
