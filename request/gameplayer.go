package request

type CreateGameplayer struct {
	GameID   string `json:"gameId" validate:"required"`
	PlayerID string `json:"playerId" validate:"required"`
}

type CreateBulkGameplayer struct {
	GameID       string   `json:"gameId" validate:"required"`
	ListPlayerID []string `json:"listPlayerId" validate:"required"`
}

type UpdateGameplayer struct {
	GameID     string `json:"gameId" validate:"required"`
	PlayerID   string `json:"playerId" validate:"required"`
	NormalGame int64  `json:"normalGame" validate:""`
	RubberGame int64  `json:"rubberGame" validate:""`
	Ball       int64  `json:"ball" validate:""`
	IsPay      bool   `json:"isPay" validate:""`
}

type PageGameplayer struct {
	Paging
	GameID     string `json:"gameId" form:"gameId" query:"gameId"`
	PlayerID   string `json:"playerId" form:"playerId" query:"playerId"`
	GameName   string `json:"gameName" form:"gameName" query:"gameName"`
	PlayerName string `json:"playerName" form:"playerName" query:"playerName"`
	Gender     string `json:"gender" form:"gender" query:"gender"`
}

type PageRankGameplayer struct {
	Paging
	Gender string `json:"gender" form:"gender" query:"gender"`
}
