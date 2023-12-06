package request

import "time"

type CreateGame struct {
	GorID           string    `json:"gorId" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:""`
	NormalGamePrice int64     `json:"normalGamePrice" validate:""`
	RubberGamePrice int64     `json:"rubberGamePrice" validate:""`
	BallPrice       int64     `json:"ballPrice" validate:""`
	GameDt          time.Time `json:"gameDt" validate:""`
	IsFinish        bool      `json:"isFinish" validate:""`
}

type UpdateGame struct {
	GorID           string    `json:"gorId" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:""`
	NormalGamePrice int64     `json:"normalGamePrice" validate:""`
	RubberGamePrice int64     `json:"rubberGamePrice" validate:""`
	BallPrice       int64     `json:"ballPrice" validate:""`
	GameDt          time.Time `json:"gameDt" validate:""`
	IsFinish        bool      `json:"isFinish" validate:""`
}

type PageGame struct {
	Paging
	GorID       string `json:"gorId" json:"gorId" json:"gorId"`
	Name        string `json:"name" json:"name" json:"name"`
	Description string `json:"description" json:"description" json:"description"`
}
