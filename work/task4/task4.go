package main

import (
	"fmt"
)

func main() {
	var places [5]string

	for {
		var token string
		_, err := fmt.Scan(&token)
		if err != nil {
			break
		}

		// Команда "очередь"
		if token == "очередь" {
			for i, name := range places {
				if name == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, name)
				}
			}
			continue
		}

		// Команда "количество"
		if token == "количество" {
			free := 0
			occupied := 0
			for _, name := range places {
				if name == "" {
					free++
				} else {
					occupied++
				}
			}
			fmt.Printf("Осталось свободных мест: %d\n", free)
			fmt.Printf("Всего человек в очереди: %d\n", occupied)
			continue
		}

		// Команда "конец"
		if token == "конец" {
			for i, name := range places {
				if name == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, name)
				}
			}
			return
		}

		// Если не команда — значит имя, читаем номер
		name := token
		var numStr string
		_, err = fmt.Scan(&numStr)
		if err != nil {
			break
		}

		var num int
		_, err = fmt.Sscanf(numStr, "%d", &num)
		if err != nil || num < 1 || num > 5 {
			fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", numStr)
			continue
		}

		// Проверка переполнения
		full := true
		for _, p := range places {
			if p == "" {
				full = false
				break
			}
		}
		if full {
			fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
			continue
		}

		// Проверка, не занято ли место
		if places[num-1] != "" {
			fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
			continue
		}

		places[num-1] = name
	}
}