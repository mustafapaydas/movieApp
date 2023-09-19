package logic

import (
	"movieProject/common/response"
	"movieProject/config"
)

type IAbstractLogic interface {
	Create(e any) response.DataResponse
	Update()
	GetById()
	Filter()
	Delete(id int)
}

type AbstractLogic struct {
	IAbstractLogic
}

func (a *AbstractLogic) Create(entity any) (any, error) {
	if result := config.GetDB().Create(entity); result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}
