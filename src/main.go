package main

import (
	"devtools/internal/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dvtls")

	mainView := views.MainView{Width: 800, Height: 500}
	mainView.Initialize(w)

	w.SetMaster()
	w.ShowAndRun()
}
