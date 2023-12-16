package models

import (
	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	NamaStatus string
}

type StatusRequest struct {
	NamaStatus string `json:"nama_status" form:"nama_status"`
}

type StatusResponse struct {
	Id       uint `json:"id"`
	NamaStatus string `json:"nama_status"`
}

func RequestToModelStatus(dt StatusRequest) Status {
	return Status{
		NamaStatus: dt.NamaStatus,
	}
}

func ModelToResponseStatus(dt Status) StatusResponse {
	return StatusResponse{
		Id:           dt.ID,
		NamaStatus: dt.NamaStatus,
	}
}

func ListModelToResponseStatus(dt []Status) []StatusResponse {
	var responses = []StatusResponse{}
	for _, v := range dt {
		responses = append(responses, ModelToResponseStatus(v))
	}
	return responses
}
