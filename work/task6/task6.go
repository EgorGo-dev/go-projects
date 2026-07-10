package main

import (
    "fmt"
    "slices"
)

type worker struct {
    Name 		string 
    Position   	string
    Salary     	uint
    Experience 	uint
    
}

type Company struct {
    workers []worker 
}

// Методы структуры Company
func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
    w := worker{
        Name:       name,
        Position:   position,
        Salary:     salary,
        Experience: experience,
    }
    c.workers = append(c.workers, w)
    return nil
}

func (c *Company) SortWorkers() ([]string, error) {
    // Определяем ранги должностей
    rank := map[string]int{
        "директор":          5,
        "зам. директора":    4,
        "начальник цеха":    3,
        "мастер":            2,
        "рабочий":           1,
    }

    // Сортируем по доходу (зарплата × стаж) по убыванию, потом по должности
    slices.SortFunc(c.workers, func(a, b worker) int {
        incomeA := int(a.Salary) * int(a.Experience)
        incomeB := int(b.Salary) * int(b.Experience)

        if incomeA != incomeB {
            return incomeB - incomeA // по убыванию дохода
        }
        // если доход одинаковый, сравниваем должность (чем выше ранг, тем выше)
        return rank[b.Position] - rank[a.Position]
    })

    // Формируем результат
    result := make([]string, len(c.workers))
    for i, w := range c.workers {
        income := int(w.Salary) * int(w.Experience)
        result[i] = fmt.Sprintf("%s — %d — %s", w.Name, income, w.Position)
    }

    return result, nil
}