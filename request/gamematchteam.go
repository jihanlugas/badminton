package request

type PageGamematchteam struct {
	Paging
	GameID        string `json:"gameId" form:"gameId" query:"gameId"`
	GamematchID   string `json:"gamematchId" form:"gamematchId" query:"gamematchId"`
	Name          string `json:"name" form:"name" query:"name"`
	GameName      string `json:"gameName" form:"gameName" query:"gameName"`
	GamematchName string `json:"gamematchName" form:"gamematchName" query:"gamematchName"`
}
