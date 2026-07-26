package main

import (
    "testing"
)

func TestReverseString(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {"hello", "olleh"},
        {"", ""},
        {"a", "a"},
        {"Привет", "тевирП"},
        {"ab", "ba"},
    }

    for _, tt := range tests {
        got := ReverseString(tt.input) 
        if got != tt.want {
            t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.want)
        }
    } 
}