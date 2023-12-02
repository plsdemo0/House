package main

import (
	"House_project/Filling_the_house/Family"
	"House_project/Filling_the_house/Things"
	"fmt"
)

// Дисклеймер
//Потом переместить создание и ввывод объектов в CreateFile

func main() {
	cat := Family.Pet{
		NamePet:   "Соня",
		AgePet:    15,
		PetBreed:  "Британская порода",
		HairColor: "Серый с белыми пятнами",
	}
	fmt.Println(cat)

	brother := Family.Brother{
		NameBrother:   "Егор",
		AgeBrother:    12,
		HeightBrother: 140,
		HairColor:     "Каштановые",
	}
	fmt.Println(brother)

	mother := Family.Mother{
		NameMother:   "Вера",
		AgeMother:    47,
		HeightMother: 170,
		HairColor:    "Каштановые",
	}
	fmt.Println(mother)

	father := Family.Father{
		NameFather:   "Михаил",
		AgeFather:    47,
		HeightFather: 175,
		HairColor:    "Седые",
	}
	fmt.Println(father)

	bed := Things.Bed{
		Length:   3,
		Width:    2,
		Material: "Дерево",
		Color:    "Кремовый",
		Room:     "Спальня",
	}
	fmt.Println(bed)

	couch_in_kids_room := Things.Couch{
		Length:   2,
		Width:    1.5,
		Material: "Текстиль, сталь",
		Color:    "Зеленый",
		Unfolds:  true,
		Room:     "Детская",
	}
	fmt.Println(couch_in_kids_room)

	couch_in_the_hall := Things.Couch{
		Length:   5,
		Width:    4,
		Material: "Текстиль, сталь, дерево",
		Color:    "белый",
		Unfolds:  true,
		Room:     "Зал",
	}
	fmt.Println(couch_in_the_hall)

	kitchen_table := Things.Table{
		Length:   3,
		Width:    2,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Кухня",
	}
	fmt.Println(kitchen_table)

	mothers_desk := Things.Table{
		Length:   3,
		Width:    1,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Зал",
	}
	fmt.Println(mothers_desk)

	my_desk := Things.Table{
		Length:   2,
		Width:    1,
		Material: "Дерево, сталь",
		Color:    "Тёмно синий",
		Room:     "Детская",
	}
	fmt.Println(my_desk)

	brothers_desk := Things.Table{
		Length:   2,
		Width:    1,
		Material: "ДСП",
		Color:    "Зеленый",
		Room:     "Детская",
	}
	fmt.Println(brothers_desk)

	hall_closet := Things.Closet{
		Length:   3,
		Width:    3,
		Material: "Дерево",
		Color:    "Белый",
		Room:     "Зал",
	}
	fmt.Println(hall_closet)

	elitebook := Things.Laptop{
		Brand:      "Lenovo",
		ScreenType: "IPS",
	}
	fmt.Println(elitebook)

	macbook := Things.Laptop{
		Brand:      "Apple",
		ScreenType: "Mini-Led",
	}
	fmt.Println(macbook)

	kitchen_tv := Things.TV{
		Brand:      "Sony",
		ScreenType: "IPS",
		ScreenSize: 47,
	}
	fmt.Println(kitchen_tv)

	hall_tv := Things.TV{
		Brand:      "Sony",
		ScreenType: "IPS",
		ScreenSize: 57,
	}
	fmt.Println(hall_tv)
}
