package main

import (
	"Markdown"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("Markdown", func(w *unison.Window) {
		w.Content().AddChild(markdown.New().Layout())
	})
}
