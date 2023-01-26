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

var coin = ""

func App(data []Coin, converter func(string, string) float64) {
	myApp := app.New()
	myWindow := myApp.NewWindow("BTC -> Moedas")

	myWindow.Resize(fyne.NewSize(680, 480))

	label := widget.NewLabel("Selecione uma moeda.")

	label.Alignment = fyne.TextAlignCenter
	input_value := widget.NewEntry()
	input_value.SetPlaceHolder("Valor")
	label_result := widget.NewLabel("Resultado")
	new_window := container.NewGridWithRows(3, input_value, label_result, widget.NewButton("Converter", func() {
		label_result.SetText(fmt.Sprintf("%f", converter(coin, input_value.Text)))
	}))
	//icon := widget.NewIcon(nil)
	button := widget.NewButtonWithIcon("", nil, func() {
		w := fyne.CurrentApp().NewWindow("Info")
		w.Resize(fyne.NewSize(380, 280))
		w.SetTitle(fmt.Sprintf("BTC -> %v", coin))
		w.SetContent(new_window)
		w.Show()
	})
	button.Hidden = true
	hbox := container.NewHBox(button, label)

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
		//icon.SetResource(theme.InfoIcon())
		button.Hidden = false
		button.SetIcon(theme.InfoIcon())
		coin = data[id].Symbol
	}
	list.OnUnselected = func(id widget.ListItemID) {
		label.SetText("Selecione uma moeda.")
		//icon.SetResource(nil)
		button.SetIcon(nil)
		button.Hidden = true
		coin = ""
	}

	listView := container.NewHSplit(list, container.NewCenter(hbox))

	myWindow.SetContent(listView)
	myWindow.ShowAndRun()
}
