package main

import (
    "encoding/json"
    "io"
)

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
    // Создаём JSON-энкодер, который пишет прямо в w
    encoder := json.NewEncoder(w)

    err := encoder.Encode(&students)
    if err != nil {
        return err
    }

    return nil
}