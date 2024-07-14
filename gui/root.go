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
	image := canvas.NewImageFromResource(icon)
	image.SetMinSize(fyne.NewSize(300, 400))

	a := app.New()
	a.SetIcon(icon)
	w := a.NewWindow("Terry A. Davis")

	quoteList := quotes.GetQuotes()

	index := rand.Int() % len(quoteList)
	label := widget.NewLabel(quoteList[index].Text)
	label.Wrapping = fyne.TextWrapBreak
	label.TextStyle = fyne.TextStyle{Italic: true}
	label.Resize(fyne.NewSize(300, 400))

	randomButton := widget.NewButton("Random", func() {
		index := rand.Int() % len(quoteList)
		label.SetText(quoteList[index].Text)
	})

	w.SetContent(
		container.NewBorder(
			image,
			randomButton,
			nil,
			nil,
			label,
		),
	)
	// w.Resize(fyne.NewSize(400, 300))
	w.CenterOnScreen()
	// w.SetFixedSize(true)
	w.ShowAndRun()
}
