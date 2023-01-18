package usecase

type Request struct {
	UID          string `json:"uid"`
	Username     string `json:"username"`
	Mail         string `json:"mail"`
	PrefectureID int    `json:"prefecture_id"`
}
