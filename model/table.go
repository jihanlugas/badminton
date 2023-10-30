package model

import "time"

type User struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	Role        string     `gorm:"not null" json:"role"`
	Email       string     `gorm:"not null" json:"email"`
	Username    string     `gorm:"not null" json:"username"`
	NoHp        string     `gorm:"not null" json:"noHp"`
	Fullname    string     `gorm:"not null" json:"fullname"`
	Passwd      string     `gorm:"not null" json:"-"`
	PassVersion int        `gorm:"not null" json:"passVersion"`
	IsActive    bool       `gorm:"not null" json:"isActive"`
	PhotoID     string     `gorm:"not null" json:"photoId"`
	LastLoginDt *time.Time `gorm:"null" json:"lastLoginDt"`
	CreateBy    string     `gorm:"not null" json:"createBy"`
	CreateDt    time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy    string     `gorm:"not null" json:"updateBy"`
	UpdateDt    time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy    string     `gorm:"not null" json:"deleteBy"`
	DeleteDt    *time.Time `gorm:"null" json:"deleteDt"`
}

type Company struct {
	ID          string     `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Balance     int64      `gorm:"not null" json:"balance"`
	CreateBy    string     `gorm:"not null" json:"createBy"`
	CreateDt    time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy    string     `gorm:"not null" json:"updateBy"`
	UpdateDt    time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy    string     `gorm:"not null" json:"deleteBy"`
	DeleteDt    *time.Time `gorm:"null" json:"deleteDt"`
}

type Usercompany struct {
	ID               string     `gorm:"primaryKey" json:"id"`
	UserID           string     `gorm:"not null" json:"userId"`
	CompanyID        string     `gorm:"not null" json:"companyId"`
	IsDefaultCompany bool       `gorm:"not null" json:"isDefaultCompany"`
	IsCreator        bool       `gorm:"not null" json:"isCreator"`
	CreateBy         string     `gorm:"not null" json:"createBy"`
	CreateDt         time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy         string     `gorm:"not null" json:"updateBy"`
	UpdateDt         time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy         string     `gorm:"not null" json:"deleteBy"`
	DeleteDt         *time.Time `gorm:"null" json:"deleteDt"`
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
	ID        string     `gorm:"primaryKey" json:"id"`
	CompanyID string     `gorm:"not null" json:"companyId"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"not null" json:"email"`
	NoHp      string     `gorm:"not null" json:"noHp"`
	Address   string     `gorm:"not null" json:"address"`
	IsActive  bool       `gorm:"not null" json:"isActive"`
	PhotoID   string     `gorm:"not null" json:"photoId"`
	CreateBy  string     `gorm:"not null" json:"createBy"`
	CreateDt  time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy  string     `gorm:"not null" json:"updateBy"`
	UpdateDt  time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy  string     `gorm:"not null" json:"deleteBy"`
	DeleteDt  *time.Time `gorm:"null" json:"deleteDt"`
}

type Gor struct {
	ID              string     `gorm:"primaryKey" json:"id"`
	CompanyID       string     `gorm:"not null" json:"companyId"`
	Name            string     `gorm:"not null" json:"name"`
	Description     string     `gorm:"not null" json:"description"`
	Address         string     `gorm:"not null" json:"address"`
	NormalGamePrice int64      `gorm:"not null" json:"normalGamePrice"`
	RubberGamePrice int64      `gorm:"not null" json:"rubberGamePrice"`
	BallPrice       int64      `gorm:"not null" json:"ballPrice"`
	CreateBy        string     `gorm:"not null" json:"createBy"`
	CreateDt        time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy        string     `gorm:"not null" json:"updateBy"`
	UpdateDt        time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy        string     `gorm:"not null" json:"deleteBy"`
	DeleteDt        *time.Time `gorm:"null" json:"deleteDt"`
}

type Game struct {
	ID              string     `gorm:"primaryKey" json:"id"`
	GorID           string     `gorm:"not null" json:"gorId"`
	Name            string     `gorm:"not null" json:"name"`
	Description     string     `gorm:"not null" json:"description"`
	NormalGamePrice int64      `gorm:"not null" json:"normalGamePrice"`
	RubberGamePrice int64      `gorm:"not null" json:"rubberGamePrice"`
	BallPrice       int64      `gorm:"not null" json:"ballPrice"`
	GameDt          time.Time  `gorm:"not null" json:"gameDt"`
	CreateBy        string     `gorm:"not null" json:"createBy"`
	CreateDt        time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy        string     `gorm:"not null" json:"updateBy"`
	UpdateDt        time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy        string     `gorm:"not null" json:"deleteBy"`
	DeleteDt        *time.Time `gorm:"null" json:"deleteDt"`
}

type Gameplayer struct {
	ID         string     `gorm:"primaryKey" json:"id"`
	GameID     string     `gorm:"not null" json:"gameId"`
	PlayerID   string     `gorm:"not null" json:"playerId"`
	NormalGame int64      `gorm:"not null" json:"normalGame"`
	RubberGame int64      `gorm:"not null" json:"rubberGame"`
	Ball       int64      `gorm:"not null" json:"ball"`
	CreateBy   string     `gorm:"not null" json:"createBy"`
	CreateDt   time.Time  `gorm:"not null" json:"createDt"`
	UpdateBy   string     `gorm:"not null" json:"updateBy"`
	UpdateDt   time.Time  `gorm:"not null" json:"updateDt"`
	DeleteBy   string     `gorm:"not null" json:"deleteBy"`
	DeleteDt   *time.Time `gorm:"null" json:"deleteDt"`
}
