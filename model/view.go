package model

import (
	"gorm.io/gorm"
	"time"
)

type UserView struct {
	ID          string         `json:"id"`
	Role        string         `json:"role"`
	Email       string         `json:"email"`
	Username    string         `json:"username"`
	NoHp        string         `json:"noHp"`
	Fullname    string         `json:"fullname"`
	Passwd      string         `json:"-"`
	PassVersion int            `json:"passVersion"`
	IsActive    bool           `json:"isActive"`
	LastLoginDt *time.Time     `json:"lastLoginDt"`
	PhotoID     string         `json:"photoId"`
	PhotoUrl    string         `json:"photoUrl"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteBy    string         `json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`
	DeleteName  string         `json:"deleteName"`
}

func (UserView) TableName() string {
	return VIEW_USER
}

type CompanyView struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Balance     int64          `json:"balance"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteBy    string         `json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`
	DeleteName  string         `json:"deleteName"`
	TotalGor    int64          `json:"totalGor"`
	TotalPlayer int64          `json:"totalPlayer"`
}

func (CompanyView) TableName() string {
	return VIEW_COMPANY
}

type UsercompanyView struct {
	ID               string         `json:"id"`
	UserID           string         `json:"userId"`
	CompanyID        string         `json:"companyId"`
	IsDefaultCompany bool           `json:"IsDefaultCompany"`
	IsCreator        bool           `json:"IsCreator"`
	CreateBy         string         `json:"createBy"`
	CreateDt         time.Time      `json:"createDt"`
	UpdateBy         string         `json:"updateBy"`
	UpdateDt         time.Time      `json:"updateDt"`
	DeleteBy         string         `json:"deleteBy"`
	DeleteDt         gorm.DeletedAt `json:"deleteDt"`
	UserName         string         `json:"userName"`
	CompanyName      string         `json:"companyName"`
	CreateName       string         `json:"createName"`
	UpdateName       string         `json:"updateName"`
	DeleteName       string         `json:"deleteName"`
}

func (UsercompanyView) TableName() string {
	return VIEW_USERCOMPANY
}

type TransactionView struct {
	ID          string    `json:"id"`
	CompanyID   string    `json:"companyId"`
	IsDebit     bool      `json:"isDebit"`
	CreateBy    string    `json:"createBy"`
	CreateDt    time.Time `json:"createDt"`
	CompanyName string    `json:"companyName"`
	CreateName  string    `json:"createName"`
}

func (TransactionView) TableName() string {
	return VIEW_TRANSACTION
}

type GorView struct {
	ID              string         `json:"id"`
	CompanyID       string         `json:"companyId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Address         string         `json:"address"`
	NormalGamePrice int64          `json:"normalGamePrice"`
	RubberGamePrice int64          `json:"rubberGamePrice"`
	BallPrice       int64          `json:"ballPrice"`
	CreateBy        string         `json:"createBy"`
	CreateDt        time.Time      `json:"createDt"`
	UpdateBy        string         `json:"updateBy"`
	UpdateDt        time.Time      `json:"updateDt"`
	DeleteBy        string         `json:"deleteBy"`
	DeleteDt        gorm.DeletedAt `json:"deleteDt"`
	CompanyName     string         `json:"companyName"`
	CreateName      string         `json:"createName"`
	UpdateName      string         `json:"updateName"`
	DeleteName      string         `json:"deleteName"`
}

func (GorView) TableName() string {
	return VIEW_GOR
}

type GameView struct {
	ID              string         `json:"id"`
	GorID           string         `json:"gorId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	NormalGamePrice int64          `json:"normalGamePrice"`
	RubberGamePrice int64          `json:"rubberGamePrice"`
	BallPrice       int64          `json:"ballPrice"`
	GameDt          time.Time      `json:"gameDt"`
	CreateBy        string         `json:"createBy"`
	CreateDt        time.Time      `json:"createDt"`
	UpdateBy        string         `json:"updateBy"`
	UpdateDt        time.Time      `json:"updateDt"`
	DeleteBy        string         `json:"deleteBy"`
	DeleteDt        gorm.DeletedAt `json:"deleteDt"`
	GorName         string         `json:"gorName"`
	CreateName      string         `json:"createName"`
	UpdateName      string         `json:"updateName"`
	DeleteName      string         `json:"deleteName"`
}

func (GameView) TableName() string {
	return VIEW_GAME
}

type PlayerView struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"companyId"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	NoHp        string         `json:"noHp"`
	Address     string         `json:"address"`
	Gender      string         `json:"gender"`
	IsActive    bool           `json:"isActive"`
	PhotoID     string         `json:"photoId"`
	CreateBy    string         `json:"createBy"`
	CreateDt    time.Time      `json:"createDt"`
	UpdateBy    string         `json:"updateBy"`
	UpdateDt    time.Time      `json:"updateDt"`
	DeleteBy    string         `json:"deleteBy"`
	DeleteDt    gorm.DeletedAt `json:"deleteDt"`
	CompanyName string         `json:"companyName"`
	CreateName  string         `json:"createName"`
	UpdateName  string         `json:"updateName"`
	DeleteName  string         `json:"deleteName"`
}

func (PlayerView) TableName() string {
	return VIEW_PLAYER
}

type GameplayerView struct {
	ID         string         `json:"id"`
	GameID     string         `json:"gameId"`
	PlayerID   string         `json:"playerId"`
	NormalGame int64          `json:"normalGame"`
	RubberGame int64          `json:"rubberGame"`
	Ball       int64          `json:"ball"`
	CreateBy   string         `json:"createBy"`
	CreateDt   time.Time      `json:"createDt"`
	UpdateBy   string         `json:"updateBy"`
	UpdateDt   time.Time      `json:"updateDt"`
	DeleteBy   string         `json:"deleteBy"`
	DeleteDt   gorm.DeletedAt `json:"deleteDt"`
	GameName   string         `json:"gameName"`
	PlayerName string         `json:"playerName"`
	CreateName string         `json:"createName"`
	UpdateName string         `json:"updateName"`
	DeleteName string         `json:"deleteName"`
}

func (GameplayerView) TableName() string {
	return VIEW_GAMEPLAYER
}
