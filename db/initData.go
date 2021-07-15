package db

import "golangorm/entity"

func InitData() {
	person := entity.Person{
		ID:   "1",
		Name: "John Doe",
		Age:  22,
	}

	GetDB().Save(&person)
}
