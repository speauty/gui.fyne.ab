package main

import (
	"github.com/flopp/go-findfont"
	"gui.fyne.ab/src/app"
	"gui.fyne.ab/src/common/cfg"
	"os"
	"strings"
)

func init() {
	fontPaths := findfont.List()
	fontName := "simkai.ttf"
	for _, path := range fontPaths {
		if strings.Contains(path, fontName) {
			_ = os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	cfg.Api().LoadConfig(".")

	new(app.App).Init().Run()
}
