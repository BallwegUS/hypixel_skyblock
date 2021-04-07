package api

type PlayerReturn struct {
	Success bool
	Player  PlayerData
}

type PlayerData struct {
	Name        string `json:"playername"`
	DisplayName string `json:"displayname"`
}
