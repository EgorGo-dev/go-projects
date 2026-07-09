// Решение задачи 
package main

import "fmt"

// Animal — интерфейс, который описывает поведение животного
type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

// animal — приватная структура, поля скрыты от внешнего доступа
type animal struct {
	name    string
	species string
	age     int
	sound   string
}

// NewAnimal — конструктор, возвращает интерфейс Animal
func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{
		name:    name,
		species: species,
		age:     age,
		sound:   sound,
	}
}

// Реализация методов интерфейса для структуры animal

func (a *animal) MakeSound() string {
	return a.sound
}

func (a *animal) GetName() string {
	return a.name
}

func (a *animal) GetInfo() string {
	// Форматируем строку с информацией о животном
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

// ZooShow — функция, которая принимает список животных и выводит информацию о каждом
func ZooShow(animals []Animal) {
	for _, a := range animals {
		fmt.Println(a.GetInfo())
		fmt.Println(a.MakeSound())
	}
}

// ZooKeeper — структура смотрителя зоопарка
type ZooKeeper struct{}

// Feed — метод смотрителя, который кормит животное
func (z ZooKeeper) Feed(animal Animal) {
	// animal.GetName() — имя, animal.MakeSound() — звук
	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", animal.GetName(), animal.MakeSound())
}
