задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Реализуйте функцию EncodeStudentsToWriter(w io.Writer, students []Student) error,
которая принимает io.Writer и слайс структур Student, и записывает этот слайс в формате JSON в writer с помощью json.Encoder.
Функция должна возвращать ошибку, если сериализация или запись не удалась.

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}