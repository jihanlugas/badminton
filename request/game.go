package request

import "time"

type CreateGame struct {
	CompanyID       string    `json:"companyId" validate:"required"`
	GorID           string    `json:"gorId" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:""`
	NormalGamePrice int64     `json:"normalGamePrice" validate:""`
	RubberGamePrice int64     `json:"rubberGamePrice" validate:""`
	BallPrice       int64     `json:"ballPrice" validate:""`
	GameDt          time.Time `json:"gameDt" validate:""`
}

type UpdateGame struct {
	CompanyID       string    `json:"companyId" validate:"required"`
	GorID           string    `json:"gorId" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:""`
	NormalGamePrice int64     `json:"normalGamePrice" validate:""`
	RubberGamePrice int64     `json:"rubberGamePrice" validate:""`
	BallPrice       int64     `json:"ballPrice" validate:""`
	GameDt          time.Time `json:"gameDt" validate:""`
	IsFinish        bool      `json:"isFinish" validate:""`
	ExpectedDebit   int64     `json:"expectedDebit"`
	Debit           int64     `json:"debit"`
}

type PageGame struct {
	Paging
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId"  validate:"required"`
	GorID       string `json:"gorId" form:"gorId" query:"gorId"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
}

type FinishGame struct {
	GameID       string              `json:"gameId" validate:"required"`
	Transactions []CreateTransaction `json:"transactions"`
}
