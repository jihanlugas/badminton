package request

type CreateGameplayer struct {
	GorID      string `json:"gorId" validate:"required"`
	PlayerID   string `json:"playerId" validate:"required"`
	NormalGame int64  `json:"normalGame" validate:""`
	RubberGame int64  `json:"rubberGame" validate:""`
	Ball       int64  `json:"ballPrice" validate:""`
}

type UpdateGameplayer struct {
	GorID      string `json:"gorId" validate:"required"`
	PlayerID   string `json:"playerId" validate:"required"`
	NormalGame int64  `json:"normalGame" validate:""`
	RubberGame int64  `json:"rubberGame" validate:""`
	Ball       int64  `json:"ballPrice" validate:""`
}

type PageGameplayer struct {
	Paging
	GorID    string `json:"gorId" json:"gorId" json:"gorId"`
	PlayerID string `json:"playerId" json:"playerId" json:"playerId"`
}
