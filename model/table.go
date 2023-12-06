package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Role        string         `gorm:"not null" json:"role"`
	Email       string         `gorm:"not null" json:"email"`
	Username    string         `gorm:"not null" json:"username"`
	NoHp        string         `gorm:"not null" json:"noHp"`
	Fullname    string         `gorm:"not null" json:"fullname"`
	Passwd      string         `gorm:"not null" json:"-"`
	PassVersion int            `gorm:"not null" json:"passVersion"`
	IsActive    bool           `gorm:"not null" json:"isActive"`
	PhotoID     string         `gorm:"not null" json:"photoId"`
	LastLoginDt *time.Time     `gorm:"null" json:"lastLoginDt"`
	CreateBy    string         `gorm:"not null" json:"createBy"`
	CreateDt    time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy    string         `gorm:"not null" json:"updateBy"`
	UpdateDt    time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy    string         `gorm:"not null" json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Company struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Balance     int64          `gorm:"not null" json:"balance"`
	CreateBy    string         `gorm:"not null" json:"createBy"`
	CreateDt    time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy    string         `gorm:"not null" json:"updateBy"`
	UpdateDt    time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy    string         `gorm:"not null" json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Usercompany struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	UserID           string         `gorm:"not null" json:"userId"`
	CompanyID        string         `gorm:"not null" json:"companyId"`
	IsDefaultCompany bool           `gorm:"not null" json:"isDefaultCompany"`
	IsCreator        bool           `gorm:"not null" json:"isCreator"`
	CreateBy         string         `gorm:"not null" json:"createBy"`
	CreateDt         time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy         string         `gorm:"not null" json:"updateBy"`
	UpdateDt         time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy         string         `gorm:"not null" json:"deleteBy"`
	DeleteDt         gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Transaction struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	CompanyID string    `gorm:"not null" json:"companyId"`
	Name      string    `gorm:"not null" json:"name"`
	IsDebit   bool      `gorm:"not null" json:"isDebit"`
	Price     int64     `gorm:"not null" json:"price"`
	CreateBy  string    `gorm:"not null" json:"createBy"`
	CreateDt  time.Time `gorm:"not null" json:"createDt"`
}

type Player struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	CompanyID string         `gorm:"not null" json:"companyId"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"not null" json:"email"`
	NoHp      string         `gorm:"not null" json:"noHp"`
	Address   string         `gorm:"not null" json:"address"`
	Gender    string         `gorm:"not null" json:"gender"`
	IsActive  bool           `gorm:"not null" json:"isActive"`
	PhotoID   string         `gorm:"not null" json:"photoId"`
	CreateBy  string         `gorm:"not null" json:"createBy"`
	CreateDt  time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy  string         `gorm:"not null" json:"updateBy"`
	UpdateDt  time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy  string         `gorm:"not null" json:"deleteBy"`
	DeleteDt  gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Gor struct {
	ID              string         `gorm:"primaryKey" json:"id"`
	CompanyID       string         `gorm:"not null" json:"companyId"`
	Name            string         `gorm:"not null" json:"name"`
	Description     string         `gorm:"not null" json:"description"`
	Address         string         `gorm:"not null" json:"address"`
	NormalGamePrice int64          `gorm:"not null" json:"normalGamePrice"`
	RubberGamePrice int64          `gorm:"not null" json:"rubberGamePrice"`
	BallPrice       int64          `gorm:"not null" json:"ballPrice"`
	CreateBy        string         `gorm:"not null" json:"createBy"`
	CreateDt        time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy        string         `gorm:"not null" json:"updateBy"`
	UpdateDt        time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy        string         `gorm:"not null" json:"deleteBy"`
	DeleteDt        gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Game struct {
	ID              string         `gorm:"primaryKey" json:"id"`
	CompanyID       string         `gorm:"not null" json:"companyId"`
	GorID           string         `gorm:"not null" json:"gorId"`
	Name            string         `gorm:"not null" json:"name"`
	Description     string         `gorm:"not null" json:"description"`
	NormalGamePrice int64          `gorm:"not null" json:"normalGamePrice"`
	RubberGamePrice int64          `gorm:"not null" json:"rubberGamePrice"`
	BallPrice       int64          `gorm:"not null" json:"ballPrice"`
	GameDt          time.Time      `gorm:"not null" json:"gameDt"`
	IsFinish        bool           `gorm:"not null" json:"isFinish"`
	CreateBy        string         `gorm:"not null" json:"createBy"`
	CreateDt        time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy        string         `gorm:"not null" json:"updateBy"`
	UpdateDt        time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy        string         `gorm:"not null" json:"deleteBy"`
	DeleteDt        gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Gameplayer struct {
	ID         string         `gorm:"primaryKey" json:"id"`
	GameID     string         `gorm:"not null" json:"gameId"`
	PlayerID   string         `gorm:"not null" json:"playerId"`
	NormalGame int64          `gorm:"not null" json:"normalGame"`
	RubberGame int64          `gorm:"not null" json:"rubberGame"`
	Ball       int64          `gorm:"not null" json:"ball"`
	CreateBy   string         `gorm:"not null" json:"createBy"`
	CreateDt   time.Time      `gorm:"not null" json:"createDt"`
	UpdateBy   string         `gorm:"not null" json:"updateBy"`
	UpdateDt   time.Time      `gorm:"not null" json:"updateDt"`
	DeleteBy   string         `gorm:"not null" json:"deleteBy"`
	DeleteDt   gorm.DeletedAt `gorm:"null" json:"deleteDt"`
}

type Gamematch struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	CompanyID      string    `gorm:"not null" json:"companyId"`
	GameID         string    `gorm:"not null" json:"gameId"`
	Name           string    `gorm:"not null" json:"name"`
	LeftTeamID     string    `gorm:"not null" json:"leftTeamId"`
	RightTeamID    string    `gorm:"not null" json:"rightTeamId"`
	LeftTeamPoint  int64     `gorm:"not null" json:"leftTeamPoint"`
	RightTeamPoint int64     `gorm:"not null" json:"rightTeamPoint"`
	IsRubber       bool      `gorm:"not null" json:"isRubber"`
	CreateBy       string    `gorm:"not null" json:"createBy"`
	CreateDt       time.Time `gorm:"not null" json:"createDt"`
}

type Gamematchscore struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	GameID         string    `gorm:"not null" json:"gameId"`
	GamematchID    string    `gorm:"not null" json:"gamematchId"`
	Set            int64     `gorm:"not null" json:"set"`
	LeftTeamScore  int64     `gorm:"not null" json:"leftTeamScore"`
	RightTeamScore int64     `gorm:"not null" json:"rightTeamScore"`
	CreateBy       string    `gorm:"not null" json:"createBy"`
	CreateDt       time.Time `gorm:"not null" json:"createDt"`
}

type Gamematchteam struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	GameID      string    `gorm:"not null" json:"gameId"`
	GamematchID string    `gorm:"not null" json:"gamematchId"`
	Name        string    `gorm:"not null" json:"name"`
	CreateBy    string    `gorm:"not null" json:"createBy"`
	CreateDt    time.Time `gorm:"not null" json:"createDt"`
}

type Gamematchteamplayer struct {
	ID              string    `gorm:"primaryKey" json:"id"`
	GameID          string    `gorm:"not null" json:"gameId"`
	GamematchID     string    `gorm:"not null" json:"gamematchId"`
	GamematchteamID string    `gorm:"not null" json:"gamematchteamId"`
	PlayerID        string    `gorm:"not null" json:"playerId"`
	CreateBy        string    `gorm:"not null" json:"createBy"`
	CreateDt        time.Time `gorm:"not null" json:"createDt"`
}
