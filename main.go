package main

import (
	"Markdown"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("Markdown", func(w *unison.Window) {
		Markdown.New().Layout(w.Content())
	})
}
