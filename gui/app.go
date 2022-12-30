package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Coin struct {
	Buy    float32 // `json:"buy"`
	Sell   float32 // `json:"sell"`
	Symbol string  // `json:"symbol"`
}

// var data = []string{"a", "string", "list"}

func App(data []Coin) {
	myApp := app.New()
	myWindow := myApp.NewWindow("BTC -> Moedas")

	myWindow.Resize(fyne.NewSize(680, 480))

	label := widget.NewLabel("Selecione uma moeda.")

	label.Alignment = fyne.TextAlignCenter
	icon := widget.NewIcon(nil)
	hbox := container.NewHBox(icon, label)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Test")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i].Symbol)
		})

	// selectedItem := widget.NewLabel()

	list.OnSelected = func(id widget.ListItemID) {
		label.SetText(fmt.Sprintf("Moeda: %v\nCompra: %v\nVenda: %v\n", data[id].Symbol, data[id].Buy, data[id].Sell))
		icon.SetResource(theme.InfoIcon())
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Selecione uma moeda.")
		icon.SetResource(nil)
	}

	listView := container.NewHSplit(list, container.NewCenter(hbox))

	myWindow.SetContent(listView)
	myWindow.ShowAndRun()
}
