package main

import (
	"bytes"
	"fmt"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	buf := make([]byte, 4094)
	window := make([]byte, 0, len(seq)*2)

	for {
		n, err := r.Read(buf)
		if n > 0 {
			window = append(window, buf[:n]...)

			if bytes.Contains(window, seq) {
				return true, nil
			}

			if len(window) > len(buf) {
				window = window[len(window)-len(seq):]
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return false, err
		}
	}

	return false, nil
}

func main() {
	src := bytes.NewReader([]byte("hello, world!"))
	fmt.Println(Contains(src, []byte("hello")))
}
