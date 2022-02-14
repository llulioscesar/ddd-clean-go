package usecase

import (
	"dddcleango/internal/app/domain"
	"dddcleango/internal/app/domain/repository"
	"dddcleango/internal/app/domain/valueobject"
	"github.com/google/uuid"
)

// AddNewCardAndEatCheese updates payment card and jerry's weight
func AddNewCardAndEatCheese(o repository.IOrder) domain.Order {
	order := o.Get()
	newCardBrand := valueobject.VISA
	if order.Payment.Card.Brand == valueobject.VISA {
		newCardBrand = valueobject.AMEX
	}
	newCard := valueobject.Card{
		ID:    uuid.New().String(),
		Brand: newCardBrand,
	}
	order.Person.Weight++
	order.Payment.Card = newCard
	o.Update(order)
	return order
}
