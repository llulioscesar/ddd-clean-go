package repository

import (
	"dddcleango/internal/app/adapter/postgresql"
	"dddcleango/internal/app/adapter/postgresql/model"
	"dddcleango/internal/app/domain"
)

// Parameter is the repository of domain.Parameter
type Parameter struct{}

// Get gets parameter
func (r Parameter) Get() domain.Parameter {
	db := postgresql.Connection()
	var param model.Parameter
	result := db.First(&param, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return domain.Parameter{
		Funds: param.Funds,
		Btc:   param.Btc,
	}
}
