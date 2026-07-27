package main

import (
    "context"
    "sync"
)

type task struct {
    idx int
    val int
}

type result struct {
    idx int
    val int
}

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
    tasks := make(chan task, len(inputs))
    results := make(chan result, len(inputs))

    var wg sync.WaitGroup

    // Запускаем workers горутин
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for t := range tasks {
                // Проверяем контекст перед обработкой
                select {
                case <-ctx.Done():
                    return
                default:
                }

                // Обрабатываем задачу
                res := fn(t.val)

                // Проверяем контекст после обработки (если fn долгая)
                select {
                case <-ctx.Done():
                    return
                default:
                }

                results <- result{idx: t.idx, val: res}
            }
        }()
    }

    // Отправляем задачи
    for i, val := range inputs {
        select {
        case <-ctx.Done():
            return nil, ctx.Err()
        default:
        }
        tasks <- task{idx: i, val: val}
    }
    close(tasks)

    // Ждём завершения горутин в отдельной горутине
    done := make(chan struct{})
    go func() {
        wg.Wait()
        close(results)
        close(done)
    }()

    // Ожидаем либо завершения, либо отмены контекста
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    case <-done:
        // Все горутины завершились, собираем результаты
        resSlice := make([]int, len(inputs))
        for r := range results {
            resSlice[r.idx] = r.val
        }
        return resSlice, nil
    }
}