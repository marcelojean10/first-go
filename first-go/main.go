package main

import (
	"github.com/marcelojean10/first-go/internal/entity"
)


func main() {
	order, erro := entity.NewOrder("1", 10, 1)
	if erro != nil {
		println(erro.Error())
	} else {
		println(order.ID)
	}

}