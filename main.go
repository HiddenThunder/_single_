package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateClock(clock *widget.Label, start time.Time) {
	diff := time.Now().Sub(start)
	formatted := time.Time{}.Add(diff).Format("15:04:05")
	clock.SetText(formatted)
}

func main() {
	start := time.Now()

	a := app.New()
	w := a.NewWindow("_maidenless_")

	// banner := widget.NewLabel("no bitches")
	w.Resize(fyne.NewSize(300, 100))
	clock := widget.NewLabel("")
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateClock(clock, start)
		}
	}()

	w.ShowAndRun()
}
