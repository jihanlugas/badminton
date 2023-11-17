package model

import (
	"github.com/jihanlugas/badminton/utils"
	"gorm.io/gorm"
	"time"
)

func (m *Player) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.ID == "" {
		m.ID = utils.GetUniqueID()
	}
	m.CreateDt = now
	m.UpdateDt = now
	return
}

func (m *Player) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	m.UpdateDt = now
	return
}
