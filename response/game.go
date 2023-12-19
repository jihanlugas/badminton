package response

import "github.com/jihanlugas/badminton/model"

type Game model.GameView

type GameDetail struct {
	Game                 model.GameView                  `json:"game"`
	Gameplayers          []model.GameplayerView          `json:"gameplayers"`
	Gamematches          []model.GamematchView           `json:"gamematches"`
	Gamematchscores      []model.GamematchscoreView      `json:"gamematchscores"`
	Gamematchteams       []model.GamematchteamView       `json:"gamematchteams"`
	Gamematchteamplayers []model.GamematchteamplayerView `json:"gamematchteamplayers"`
}
