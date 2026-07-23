package main 

import (

	"context"

	"bytes"

	"io"

)

func Contains(ctx context.Context,r io.Reader, seq []byte) (bool, error) {

    buff := make([]byte, len(seq))

    // Начальное чтение данных в буфер

    n, err := r.Read(buff)

    if err != nil && err != io.EOF {

        return false, err

    }

    if n != len(seq) {

        return false, nil

    }

    // Процесс поиска в потоке данных

    for {

        select {

		case <-ctx.Done():

			// Если контекст отменен, возвращаем ошибку отмены

			return false, ctx.Err()

		default:

        }

        if bytes.Equal(seq, buff) {

            return true, nil

        }

        buff = append(buff[1:], 0) // Сдвиг в буфере

        _, err := r.Read(buff[len(buff)-1:])

        if err != nil {

            return false, nil

        }

    }

}