package usecase

type Request struct {
	UID      string `json:"uid"`
	Nickname string `json:"nickname"`
	IconURL  string `json:"icon_url"`
}
