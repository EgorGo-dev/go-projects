package main

import (
    "context"
    "errors"
    "sync"
)

var ErrInvalidWorkers = errors.New("invalid workers")

type Item struct {
    ID   int
    Data string
}

type Stats struct {
    Processed int
    Failed    int
    Results   []int
}

func ProcessBatch(
    ctx context.Context,
    items []Item,
    maxWorkers int,
    fn func(ctx context.Context, item Item) (int, error),
) (*Stats, error) {
    if maxWorkers <= 0 {
        stats := &Stats{
            Results: make([]int, len(items)),
        }
        return stats, ErrInvalidWorkers
    }

    if len(items) == 0 {
        return &Stats{Results: make([]int, 0)}, nil
    }

    results := make([]int, len(items))
    processed := make([]bool, len(items))
    failed := make([]bool, len(items))
    errs := make([]error, len(items))

    sem := make(chan struct{}, maxWorkers)
    var wg sync.WaitGroup

    for i := range items {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()

            if ctx.Err() != nil {
                return
            }

            select {
            case sem <- struct{}{}:
                defer func() { <-sem }()
            case <-ctx.Done():
                return
            }

            res, err := fn(ctx, items[i])
            if err != nil {
                failed[i] = true
                errs[i] = err
            } else {
                results[i] = res
                processed[i] = true
            }
        }(i)
    }

    wg.Wait()

    stats := &Stats{
        Results: results,
    }
    for _, p := range processed {
        if p {
            stats.Processed++
        }
    }
    for _, f := range failed {
        if f {
            stats.Failed++
        }
    }

    for i, f := range failed {
        if f {
            return stats, errs[i]
        }
    }

    if ctx.Err() != nil {
        return stats, ctx.Err()
    }

    return stats, nil
}