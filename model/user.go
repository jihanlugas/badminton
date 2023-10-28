package model

import (
	"github.com/jihanlugas/badminton/utils"
	"gorm.io/gorm"
	"time"
)

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if u.ID != "" {
		u.ID = utils.GetUniqueID()
	}

	u.CreateDt = now
	u.UpdateDt = now
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.DeleteDt == nil {
		now := time.Now()
		u.UpdateDt = now
	}
	return
}
