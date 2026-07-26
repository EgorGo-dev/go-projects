package main

import (
    "testing"
)

func TestAreAnagrams(t *testing.T) {
    tests := []struct {
        name string
        str1 string
        str2 string
        want bool
    }{
        {
            name: "одинаковые слова",
            str1: "listen",
            str2: "silent",
            want: true,
        },
        {
            name: "с разным регистром",
            str1: "Listen",
            str2: "SILENT",
            want: true,
        },
        {
            name: "не анаграммы",
            str1: "hello",
            str2: "world",
            want: false,
        },
        {
            name: "разная длина",
            str1: "abc",
            str2: "abcd",
            want: false,
        },
        {
            name: "пустые строки",
            str1: "",
            str2: "",
            want: true,
        },
        {
            name: "одна пустая, другая нет",
            str1: "",
            str2: "a",
            want: false,
        },
        {
            name: "с пробелами",
            str1: "a b",
            str2: "b a",
            want: true,
        },
        {
            name: "с Unicode (кириллица)",
            str1: "каша",
            str2: "шака",
            want: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := AreAnagrams(tt.str1, tt.str2)
            if got != tt.want {
                t.Errorf("AreAnagrams(%q, %q) = %v, want %v", tt.str1, tt.str2, got, tt.want)
            }
        })
    }
}