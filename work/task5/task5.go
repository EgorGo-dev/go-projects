package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 1. Ввод даты
	var input string
	fmt.Scan(&input)
	date, _ := time.Parse("02.01.2006", input) // макет правильный
	newDate := date.AddDate(0, 0, 15)

	// 2. Ввод ФИО
	var name, surname, patronymic string
	fmt.Scan(&name, &surname, &patronymic)

	// 3. Ввод трёх сумм
	var sum1, sum2, sum3 float64
	fmt.Scan(&sum1, &sum2, &sum3)

	// 4. Вычисления
	total := sum1 + sum2 + sum3
	rub := int(total)
	kop := int(math.Round((total - float64(rub)) * 100))

	// 5. Вывод
	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n", surname, name, patronymic)
	fmt.Printf("Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n", newDate.Format("02.01.2006"))
	fmt.Printf("Общая сумма выплат составит %d руб. %d коп.\n", rub, kop)
	fmt.Println("\nС уважением,\nГл. бух. Иванов А.Е.")
}