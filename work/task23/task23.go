package main

import (
    "encoding/json"
    "io"
)

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error) {
    var student Student

    // Создаём декодер из ридера
    decoder := json.NewDecoder(r)

    // Декодируем JSON в структуру
    err := decoder.Decode(&student)
    if err != nil {
        return Student{}, err
    }

    return student, nil
}