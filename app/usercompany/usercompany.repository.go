package usercompany

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Usercompany, error)
	GetCreatorByCompanyId(conn *gorm.DB, companyID string) (model.Usercompany, error)
	GetCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.Usercompany, error)
	GetViewById(conn *gorm.DB, id string) (model.UsercompanyView, error)
	GetViewCreatorByCompanyId(conn *gorm.DB, companyID string) (model.UsercompanyView, error)
	GetViewCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.UsercompanyView, error)
	Create(conn *gorm.DB, data model.Usercompany) error
	Update(conn *gorm.DB, data model.Usercompany) error
	Delete(conn *gorm.DB, data model.Usercompany) error
	Page(conn *gorm.DB, req *request.PageUsercompany) ([]model.UsercompanyView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Usercompany, error) {
	var err error
	var data model.Usercompany

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetCreatorByCompanyId(conn *gorm.DB, companyID string) (model.Usercompany, error) {
	var err error
	var data model.Usercompany

	err = conn.
		Where("company_id = ? ", companyID).
		Where("is_creator = ? ", true).
		First(&data).Error
	return data, err
}

func (r repository) GetCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.Usercompany, error) {
	var err error
	var data model.Usercompany

	err = conn.
		Where("user_id = ? ", userID).
		Where("is_default_company = ? ", true).
		First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.UsercompanyView, error) {
	var err error
	var data model.UsercompanyView

	err = conn.Where("id = ? ", id).First(&data).Error
	return data, err
}

func (r repository) GetViewCreatorByCompanyId(conn *gorm.DB, companyID string) (model.UsercompanyView, error) {
	var err error
	var data model.UsercompanyView

	err = conn.
		Where("company_id = ? ", companyID).
		Where("is_creator = ? ", true).
		First(&data).Error
	return data, err
}

func (r repository) GetViewCompanyDefaultByUserId(conn *gorm.DB, userID string) (model.UsercompanyView, error) {
	var err error
	var data model.UsercompanyView

	err = conn.
		Where("user_id = ? ", userID).
		Where("is_default_company = ? ", true).
		First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Usercompany) error {
	return conn.Create(&data).Error
}

func (r repository) Update(conn *gorm.DB, data model.Usercompany) error {
	return conn.Save(&data).Error
}

func (r repository) Delete(conn *gorm.DB, data model.Usercompany) error {
	return conn.Delete(&data).Error
}

func (r repository) Page(conn *gorm.DB, req *request.PageUsercompany) ([]model.UsercompanyView, int64, error) {
	var err error
	var data []model.UsercompanyView
	var count int64

	query := conn.Model(&data).
		Where("LOWER(company_id) LIKE LOWER(?)", "%"+req.CompanyID+"%").
		Where("LOWER(user_id) LIKE LOWER(?)", "%"+req.UserID+"%")

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
