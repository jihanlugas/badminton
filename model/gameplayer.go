package model

import (
	"github.com/jihanlugas/badminton/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Gameplayer) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID == "" {
		m.ID = utils.GetUniqueID()
	}
	m.CreateDt = now
	m.UpdateDt = now
	return
}

func (m *Gameplayer) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.DeleteBy == "" {
		now := time.Now()
		m.UpdateDt = now
	}
	return
}

func (m *Gameplayer) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Save(m).Error
}
