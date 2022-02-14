package usecase

import (
	"dddcleango/internal/app/domain"
	"dddcleango/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
