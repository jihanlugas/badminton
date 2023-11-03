package gor

import (
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Gor, error)
	GetViewById(conn *gorm.DB, id string) (model.GorView, error)
	Create(conn *gorm.DB, data model.Gor) error
	Update(conn *gorm.DB, data model.Gor) error
	Delete(conn *gorm.DB, data model.Gor) error
	Page(conn *gorm.DB, req *request.PageGor) ([]model.GorView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Gor, error) {
	var err error
	var data model.Gor

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.GorView, error) {
	var err error
	var data model.GorView

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Gor) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Gor) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Gor) error {
	now := time.Now()
	data.DeleteDt = &now

	return conn.Save(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageGor) ([]model.GorView, int64, error) {
	var err error
	var data []model.GorView
	var count int64

	err = conn.Model(&data).
		Where("company_id LIKE ?", "%"+req.CompanyID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Where("address LIKE ?", "%"+req.Address+"%").
		Where("delete_dt IS NULL").
		Count(&count).Error
	if err != nil {
		return data, count, err
	}

	err = conn.
		Where("company_id LIKE ?", "%"+req.CompanyID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Where("address LIKE ?", "%"+req.Address+"%").
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