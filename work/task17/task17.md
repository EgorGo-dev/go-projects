задача из Яндекс LMS - Яндекс Лицея

Ограничение времени	60 секунд
Ограничение памяти	300 Мб
Ввод	стандартный ввод или input.txt
Вывод	стандартный вывод или output.txt
Напишите тест для функции:

    func ReverseString(input string) string {
        runes := []rune(input)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        return string(runes)
    }