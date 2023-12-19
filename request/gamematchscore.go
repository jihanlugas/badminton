package request

type PageGamematchscore struct {
	Paging
	GameID        string `json:"gameId" form:"gameId" query:"gameId"`
	GamematchID   string `json:"gamematchId" form:"gamematchId" query:"gamematchId"`
	GameName      string `json:"gameName" form:"gameName" query:"gameName"`
	GamematchName string `json:"gamematchName" form:"gamematchName" query:"gamematchName"`
}
