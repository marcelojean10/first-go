package entity

import "errors"

type Order struct {
	ID 		   string
	Price 	   float64
	Tax 	   float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order {
		ID: id,
		Price: price,
		Tax: tax,
	}

	erro := order.Validate()
	if erro != nil {
		return nil, erro
	}

	return order, nil
}

func (o *Order) Validate () error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0  {
		return errors.New("price must be greater than zero")
	}

	if o.Tax <= 0  {
		return errors.New("invalid tax")
	}
	return nil
}

// * go trabalha com ponteiro
func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}