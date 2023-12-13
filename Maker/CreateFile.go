package Maker

import (
	"House_project/Filling_the_house/Family"
	"House_project/Filling_the_house/Things"
)

func Create() {

	var list_of_things []interface{}

	cat := Family.Pet{
		PetType:   "Кошка",
		Name:      "Соня",
		Age:       15,
		PetBreed:  "Британская порода",
		HairColor: "Серый с белыми пятнами",
	}
	list_of_things = append(list_of_things, cat)

	brother := Family.Brother{
		Name:      "Егор",
		Age:       12,
		Height:    140.123,
		HairColor: "Каштановые",
	}
	list_of_things = append(list_of_things, brother)

	mother := Family.Mother{
		Name:      "Вера",
		Age:       47,
		Height:    170.21,
		HairColor: "Каштановые",
	}
	list_of_things = append(list_of_things, mother)

	father := Family.Father{
		Name:      "Михаил",
		Age:       47,
		Height:    175.91,
		HairColor: "Седые",
	}
	list_of_things = append(list_of_things, father)

	bed := Things.Bed{
		Length:   3,
		Width:    2,
		Material: "Дерево",
		Color:    "Кремовый",
		Room:     "Спальня",
	}
	list_of_things = append(list_of_things, bed)

	couch_in_kids_room := Things.Couch{
		Length:   2,
		Width:    1.5,
		Material: "Текстиль, сталь",
		Color:    "Зеленый",
		Unfolds:  true,
		Room:     "Детская",
	}
	list_of_things = append(list_of_things, couch_in_kids_room)

	couch_in_the_hall := Things.Couch{
		Length:   5,
		Width:    4,
		Material: "Текстиль, сталь, дерево",
		Color:    "белый",
		Unfolds:  true,
		Room:     "Зал",
	}
	list_of_things = append(list_of_things, couch_in_the_hall)

	kitchen_table := Things.Table{
		Length:   3,
		Width:    2,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Кухня",
	}
	list_of_things = append(list_of_things, kitchen_table)

	mothers_desk := Things.Table{
		Length:   3,
		Width:    1,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Зал",
	}
	list_of_things = append(list_of_things, mothers_desk)

	my_desk := Things.Table{
		Length:   2,
		Width:    1,
		Material: "Дерево, сталь",
		Color:    "Тёмно синий",
		Room:     "Детская",
	}
	list_of_things = append(list_of_things, my_desk)

	brothers_desk := Things.Table{
		Length:   2,
		Width:    1,
		Material: "ДСП",
		Color:    "Зеленый",
		Room:     "Детская",
	}
	list_of_things = append(list_of_things, brothers_desk)

	hall_closet := Things.Closet{
		Length:   3,
		Width:    3,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Зал",
	}
	list_of_things = append(list_of_things, hall_closet)

	elitebook := Things.Laptop{
		Brand:      "Lenovo",
		ScreenType: "IPS",
	}
	list_of_things = append(list_of_things, elitebook)

	macbook := Things.Laptop{
		Brand:      "Apple",
		ScreenType: "Mini-Led",
	}
	list_of_things = append(list_of_things, macbook)

	kitchen_tv := Things.TV{
		Brand:      "Sony",
		ScreenType: "IPS",
		ScreenSize: 47,
	}
	list_of_things = append(list_of_things, kitchen_tv)

	hall_tv := Things.TV{
		Brand:      "Sony",
		ScreenType: "IPS",
		ScreenSize: 57,
	}
	list_of_things = append(list_of_things, hall_tv)

	PrintConsole_FamilyPet(list_of_things)
}
