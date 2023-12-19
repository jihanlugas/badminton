package request

type PageGamematchteamplayer struct {
	Paging
	GameID            string `json:"gameId" form:"gameId" query:"gameId"`
	GamematchID       string `json:"gamematchId" form:"gamematchId" query:"gamematchId"`
	GamematchteamID   string `json:"gamematchteamId" form:"gamematchteamId" query:"gamematchteamId"`
	PlayerID          string `json:"playerId" form:"playerId" query:"playerId"`
	GameName          string `json:"gameName" form:"gameName" query:"gameName"`
	GamematchName     string `json:"gamematchName" form:"gamematchName" query:"gamematchName"`
	GamematchteamName string `json:"gamematchteamName" form:"gamematchteamName" query:"gamematchteamName"`
	PlayerName        string `json:"playerName" form:"playerName" query:"playerName"`
}
