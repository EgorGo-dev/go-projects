package main

import (
    "strings"
    "time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
    correct := 0

    for i := range questions {
        select {
        case userAnswer := <-answerCh:
            // Убираем пробелы и сравниваем без учёта регистра
            userAnswer = strings.TrimSpace(userAnswer)
            if strings.EqualFold(userAnswer, answers[i]) {
                correct++
            }
        case <-time.After(1 * time.Second):
            // таймаут — просто переходим к следующему вопросу
        }
    }

    return correct
}