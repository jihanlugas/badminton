package request

type CreateMatchpointGamematch struct {
	CompanyID      string `json:"companyId" validate:"required"`
	GameID         string `json:"gameId" validate:"required"`
	MatchName      string `json:"matchName" validate:"required"`
	LeftTeamPoint  int64  `json:"leftPoint" validate:""`
	RightTeamPoint int64  `json:"rightPoint" validate:""`
	IsRubber       bool   `json:"isRubber" validate:""`
	GameMatchTeams []struct {
		Name                 string `json:"name" validate:"required"`
		GameMatchTeamPlayers []struct {
			PlayerID string `json:"playerId" validate:"required"`
		} `json:"gameMatchTeamPlayers"`
	} `json:"gameMatchTeams"`
	GameMatchScores []struct {
		LeftScore  int64 `json:"leftScore" validate:""`
		RightScore int64 `json:"rightScore" validate:""`
	} `json:"gameMatchScores"`
	Ball int64 `json:"ball" validate:""`
}

type CreateMatchGamematch struct {
	CompanyID      string `json:"companyId" validate:"required"`
	GameID         string `json:"gameId" validate:"required"`
	MatchName      string `json:"matchName" validate:"required"`
	IsRubber       bool   `json:"isRubber" validate:""`
	GameMatchTeams []struct {
		Name                 string `json:"name" validate:"required"`
		GameMatchTeamPlayers []struct {
			PlayerID string `json:"playerId" validate:"required"`
		} `json:"gameMatchTeamPlayers"`
	} `json:"gameMatchTeams"`
	Ball int64 `json:"ball" validate:""`
}

type UpdateGamematch struct {
}

type PageGamematch struct {
	Paging
	CompanyID string `json:"companyId" form:"companyId" query:"companyId"  validate:"required"`
	GameID    string `json:"gameId" form:"gameId" query:"gameId"`
	Name      string `json:"name" form:"name" query:"name"`
}
