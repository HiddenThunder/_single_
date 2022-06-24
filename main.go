package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateClock(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	clock := widget.NewLabel("")
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateClock(clock)
		}
	}()

	w.ShowAndRun()
}
