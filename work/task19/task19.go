package main

func Filter[T any](arr []T, predicate func(T) bool) []T {

    result := []T{}

    for _, v := range arr {
        if predicate(v) {
            result = append(result, v)
        }
    }

    return result
}
