package gui

import (
	b64 "encoding/base64"
	"math/rand/v2"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MartinJindra/terry/quotes"
	"github.com/MartinJindra/terry/res"
)

func Execute() {
	imageBytes, _ := b64.StdEncoding.DecodeString(res.GetImageEncoded())
	icon := (fyne.NewStaticResource("Icon.png", imageBytes))

	a := app.New()
	a.SetIcon(icon)
	w := a.NewWindow("Terry A. Davis")

	quoteList := quotes.GetQuotes()

	index := rand.Int() % len(quoteList)
	label := widget.NewLabel(quoteList[index].Text)
	label.Wrapping = fyne.TextWrapBreak
	label.TextStyle = fyne.TextStyle{Italic: true}

	randomButton := widget.NewButton("Random", func() {
		index := rand.Int() % len(quoteList)
		label.SetText(quoteList[index].Text)
	})

	w.SetContent(
		container.NewAdaptiveGrid(1,
			canvas.NewImageFromResource(icon),
			container.NewBorder(
				label,
				randomButton,
				nil,
				nil,
			),
		),
	)
	w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.ShowAndRun()
}
