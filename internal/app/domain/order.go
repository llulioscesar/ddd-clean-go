package domain

import "dddcleango/internal/app/domain/valueobject"

type Order struct {
	ID      string
	Payment valueobject.Payment
	Person  Person
}
