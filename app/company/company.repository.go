package company

import (
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Company, error)
	GetByName(conn *gorm.DB, name string) (model.Company, error)
	GetViewById(conn *gorm.DB, id string) (model.CompanyView, error)
	GetViewByName(conn *gorm.DB, name string) (model.CompanyView, error)
	Create(conn *gorm.DB, data model.Company) error
	Update(conn *gorm.DB, data model.Company) error
	Delete(conn *gorm.DB, data model.Company) error
	Page(conn *gorm.DB, req *request.PageCompany) ([]model.CompanyView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Company, error) {
	var err error
	var data model.Company

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetByName(conn *gorm.DB, name string) (model.Company, error) {
	var err error
	var data model.Company

	err = conn.Where("name = ? ", name).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.CompanyView, error) {
	var err error
	var data model.CompanyView

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewByName(conn *gorm.DB, name string) (model.CompanyView, error) {
	var err error
	var data model.CompanyView

	err = conn.Where("name = ? ", name).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Company) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Company) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Company) error {
	now := time.Now()
	data.DeleteDt = &now

	return conn.Save(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageCompany) ([]model.CompanyView, int64, error) {
	var err error
	var data []model.CompanyView
	var count int64

	err = conn.Model(&data).
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("description LIKE ?", "%"+req.Description+"%").
		Where("delete_dt IS NULL").
		Count(&count).Error
	if err != nil {
		return data, count, err
	}

	err = conn.
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
