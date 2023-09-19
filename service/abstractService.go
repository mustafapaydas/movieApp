package service

import (
	"movieProject/common/response"
	"movieProject/logic"
	"net/http"
)

type IAbstractService interface {
	Create(entity any) response.DataResponse
	Update()
	GetById()
	Filter()
	Delete(id int)
}

type AbstractService struct {
	IAbstractService
}

var _logic = logic.AbstractLogic{}

func (abstract *AbstractService) Create(entity any) response.DataResponse {
	result, err := _logic.Create(entity)
	if err != nil {
		res := response.DataResponse{Error: err, StatusCode: http.StatusCreated}
		return res
	}
	res := response.DataResponse{Data: result, StatusCode: http.StatusInternalServerError}
	return res
}
