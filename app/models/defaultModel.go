package models

type SuccessResponse struct {
	Detail string `json:"detail"`
}

type FailedResponse struct {
	Detail string `json:"detail"`
}
