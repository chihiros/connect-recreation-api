package usecase

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrorCode    string `json:"code"`
	ErrorMessage string `json:"message"`
}
