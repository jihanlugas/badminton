package transaction

import (
	"fmt"
	"github.com/jihanlugas/badminton/model"
	"github.com/jihanlugas/badminton/request"
	"gorm.io/gorm"
)

type Repository interface {
	GetById(conn *gorm.DB, id string) (model.Transaction, error)
	GetViewById(conn *gorm.DB, id string) (model.TransactionView, error)
	Create(conn *gorm.DB, data model.Transaction) error
	//Update(conn *gorm.DB, data model.Transaction) error
	//Delete(conn *gorm.DB, data model.Transaction) error
	Page(conn *gorm.DB, req *request.PageTransaction) ([]model.TransactionView, int64, error)
}

type repository struct {
}

func (r repository) GetById(conn *gorm.DB, id string) (model.Transaction, error) {
	var err error
	var data model.Transaction

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) GetViewById(conn *gorm.DB, id string) (model.TransactionView, error) {
	var err error
	var data model.TransactionView

	err = conn.Where("id = ? ", id).Where("delete_dt IS NULL").First(&data).Error
	return data, err
}

func (r repository) Create(conn *gorm.DB, data model.Transaction) error {
	return conn.Create(&data).Error
}

//func (r repository) Update(conn *gorm.DB, data model.Transaction) error {
//	return conn.Save(&data).Error
//}

//func (r repository) Delete(conn *gorm.DB, data model.Transaction) error {
//	now := time.Now()
//	data.DeleteDt = &now
//
//	return conn.Save(&data).Error
//}

func (r repository) Page(conn *gorm.DB, req *request.PageTransaction) ([]model.TransactionView, int64, error) {
	var err error
	var data []model.TransactionView
	var count int64

	query := conn.Model(&data).
		Where("company_id LIKE ?", "%"+req.CompanyID+"%").
		Where("name LIKE ?", "%"+req.Name+"%").
		Where("delete_dt IS NULL")

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
