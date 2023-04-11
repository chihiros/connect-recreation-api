package usecase

type Request struct {
	UUID     string `json:"uuid"`
	Nickname string `json:"nickname"`
	IconURL  string `json:"icon_url"`
}
