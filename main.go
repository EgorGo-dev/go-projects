package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	cursorHide  = "\033[?25l"
	cursorShow  = "\033[?25h"
	clearScreen = "\033[2J\033[H"

	// ANSI цвета (256-color palette)
	colorBlack      = "\033[38;5;0m"
	colorDarkGray   = "\033[38;5;238m"
	colorMidGray    = "\033[38;5;240m"
	colorOrange     = "\033[38;5;166m"
	colorBrightOrange = "\033[38;5;208m"
	colorYellow     = "\033[38;5;220m"
	resetColor      = "\033[0m"
)

func main() {
	fmt.Print(cursorHide)
	defer fmt.Print(cursorShow)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		os.Exit(0)
	}()

	// Базовые размеры: можно увеличить, если терминал позволяет
	const baseWidth = 80
	const baseHeight = 40

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	var t float64

	for range ticker.C {
		fmt.Print(clearScreen)
		t += 0.12

		for y := 0; y < baseHeight; y++ {
			for x := 0; x < baseWidth; x++ {
				// Центрируем координаты
				cx := float64(x) - float64(baseWidth)/2.0
				cy := float64(y) - float64(baseHeight)/2.0

				// Масштабируем, чтобы дыра была крупнее
				scale := 1.3
				rx := cx * scale
				ry := cy * scale

				r := math.Sqrt(rx*rx + ry*ry)
				angle := math.Atan2(ry, rx)

				// Искажение пространства (гравитационная линза)
				dist := 1.0 / (1.0 + 0.06*r)
				rr := r * dist

				// Радиус чёрной дыры и аккреционного диска
				const holeRadius = 12.0
				const diskWidth = 6.0

				var ch rune
				var color string

				if rr < holeRadius*0.95 {
					// Внутри чёрной дыры — почти полностью чёрный
					ch = ' '
					color = colorBlack
				} else if rr < holeRadius+diskWidth {
					// Аккреционный диск: плавные переходы
					// Вычисляем «яркость» по синусу от угла и времени
					rot := angle + t
					brightness := (math.Sin(rot*3.0) + 1.0) / 2.0

					// Плавная граница диска: интерполяция от тёмного к яркому
					edgeFade := math.Min(1.0, (rr-(holeRadius-1.0))/diskWidth)
					finalBrightness := brightness*edgeFade + (1.0-edgeFade)*0.2

					// Выбор цвета по яркости
					switch {
					case finalBrightness < 0.25:
						color = colorDarkGray
					case finalBrightness < 0.5:
						color = colorMidGray
					case finalBrightness < 0.75:
						color = colorOrange
					default:
						color = colorBrightOrange
					}

					ch = '█'
				} else {
					// Фон: очень тёмный, почти чёрный, чтобы дыра выделялась
					ch = ' '
					color = colorDarkGray
				}

				fmt.Printf("%s%c%s", color, ch, resetColor)
			}
			fmt.Println()
		}

		// Подсказка внизу
		fmt.Printf("\nЧёрная дыра (наблюдение). Ctrl+C для выхода. Время: %.1f\n", t)
	}
}