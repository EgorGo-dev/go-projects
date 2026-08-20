package main

import (
	"fmt"
	"strings"
)

// User — структура пользователя
type User struct {
	ID    int
	Name  string
	Email string
}

// FilterByDomain возвращает пользователей, чей email заканчивается на указанный домен.
// Сравнение регистронезависимое. Функция чистая: не меняет входной слайс.
func FilterByDomain(users []User, domain string) []User {
	// Приводим домен к нижнему регистру один раз — это чуть эффективнее, чем делать это в цикле
	normalizedDomain := strings.ToLower(domain)

	var result []User

	for _, u := range users {
		// Приводим email к нижнему регистру и проверяем суффикс
		if strings.HasSuffix(strings.ToLower(u.Email), normalizedDomain) {
			result = append(result, u)
		}
	}

	return result
}

func main() {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@yandex.ru"},
		{ID: 2, Name: "Bob",   Email: "bob@gmail.com"},
		{ID: 3, Name: "Cara",  Email: "cara@Yandex.RU"},
		{ID: 4, Name: "Dan",   Email: "dan@mail.ru"},
	}

	domain := "yandex.ru"
	filtered := FilterByDomain(users, domain)

	fmt.Printf("Найдено пользователей с доменом %q: %d\n", domain, len(filtered))
	for _, u := range filtered {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
	}
}