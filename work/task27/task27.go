package main

import (
    "strconv"
    "strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string) {
    parts := strings.SplitN(statusLine, " ", 3)
    if len(parts) != 3 {
        return 0, ""
    }
    code, _ = strconv.Atoi(parts[1])
    reason = parts[2]

    return code, reason
}