package entities

const TableName = "player"

type Player struct {
	PlayerId   int64  `json:"player_id"`
	PlayerName string `json:"player_name"`
	Score      int32 `json:"score"`
}

type PlayerId struct {
	PlayerId int64 `json:"player_id"`
}

type PlayerInfo struct {
	PlayerName string `json:"player_name"`
	Score      int32 `json:"score"`
}
