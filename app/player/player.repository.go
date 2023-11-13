package player

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Player, error)
	GetViewById(conn *gorm.DB, id string) (model.PlayerView, error)
	Create(conn *gorm.DB, data model.Player) error
	Update(conn *gorm.DB, data model.Player) error
	Delete(conn *gorm.DB, data model.Player) error
	Page(conn *gorm.DB, req *request.PagePlayer) ([]model.PlayerView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Player, error) {
	var err error
	var data model.Player

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.PlayerView, error) {
	var err error
	var data model.PlayerView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Player) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Player) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Player) error {
	return conn.Delete(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PagePlayer) ([]model.PlayerView, int64, error) {
	var err error
	var data []model.PlayerView
	var count int64

	query := conn.Model(&data).
		Where("company_id LIKE ?", "%"+req.CompanyID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("email LIKE ?", "%"+req.Email+"%").
		Where("no_hp LIKE ?", "%"+req.NoHp+"%").
		Where("address LIKE ?", "%"+req.Address+"%")

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
