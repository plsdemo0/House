package Maker

import (
	"House_project/Filling_the_house/Family"
	"House_project/Filling_the_house/Things"
	"fmt"
)

func PrintConsole_FamilyPet(x []interface{}) {
	for i := 0; i < len(x); i++ {
		switch v := x[i].(type) {

		//семья
		case Family.Pet:
			fmt.Println(fmt.Sprintf("\tПитомец\n"+
				"Тип питомца:\t%s\n"+
				"Имя питомца:\t%s\n"+
				"Возраст питомца:\t%d\n"+
				"Порода питомца:\t%s\n"+
				"Цвет шерсти питомца:\t%s\n", v.PetType, v.Name, v.Age, v.PetBreed, v.HairColor))
		case Family.Brother:
			fmt.Println(fmt.Sprintf("\tБрат\n"+
				"Имя брата:\t%s\n"+
				"Возраст брата:\t%d\n"+
				"Рост брата:\t%.2f\n"+
				"Цвет волос:\t%s\n", v.Name, v.Age, v.Height, v.HairColor))
		case Family.Mother:
			fmt.Println(fmt.Sprintf("\tМама\n"+
				"Имя мамы:\t%s\n"+
				"Возраст мамы:\t%d\n"+
				"Рост мамы:\t%.2f\n"+
				"Цвет волос:\t%s\n", v.Name, v.Age, v.Height, v.HairColor))
		case Family.Father:
			fmt.Println(fmt.Sprintf("\tОтец\n"+
				"Имя отец:\t%s\n"+
				"Возраст отца:\t%d\n"+
				"Рост отца:\t%.2f\n"+
				"Цвет волос:\t%s\n", v.Name, v.Age, v.Height, v.HairColor))

			//мебель
		case Things.Bed:
			fmt.Println(fmt.Sprintf("\tКровать\n"+
				"Длина: \t%.2f\n"+
				"Ширина: \t%.2f\n"+
				"Материал: \t:%s\n"+
				"Цвет: \t%s\n"+
				"Комната: \t%s\n", v.Length, v.Width, v.Material, v.Color, v.Room))
		case Things.Couch:
			fmt.Println(fmt.Sprintf("\tДиван\n"+
				"Длина: \t%.2f\n"+
				"Ширина: \t%.2f\n"+
				"Материал: \t%s\n"+
				"Цвет: \t%s\n"+
				"Раскладной: \t%t\n"+
				"Комната: \t%s\n", v.Length, v.Width, v.Material, v.Color, v.Unfolds, v.Room))
		case Things.Table:
			fmt.Println(fmt.Sprintf("\tСтол\n"+
				"Длина:\t%.2f\n"+
				"Ширина:\t%.2f\n"+
				"Материал:\t%s\n"+
				"Цвет:\t%s\n"+
				"Комната:\t%s\n", v.Length, v.Width, v.Material, v.Color, v.Room))
		case Things.Closet:
			fmt.Println(fmt.Sprintf("\tШкаф\n"+
				"Длина:\t%.2f\n"+
				"Ширина:\t%.2f\n"+
				"Материал:\t%s\n"+
				"Цвет:\t%s\n"+
				"Комната:\t%s\n", v.Length, v.Width, v.Material, v.Color, v.Room))

			//Девайсы
		case Things.Laptop:
			fmt.Println(fmt.Sprintf("\tНоутбук\n"+
				"Производитель:\t%s\n"+
				"Тип экрана:\t%s\n", v.Brand, v.ScreenType))
		case Things.TV:
			fmt.Println(fmt.Sprintf("\tТелевизор\n"+
				"Производитель:\t%s\n"+
				"Тип экрана:\t%s\n"+
				"Размер экрана\t%d\n", v.Brand, v.ScreenType, v.ScreenSize))
			//Если ошибка
		default:
			fmt.Println("Неизвестная вещь")
		}
	}
}
