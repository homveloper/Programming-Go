package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

type FPerson struct {
	Name    string    `json:"name"`
	Age     uint8     `json:"age"`
	Address *FAddress `json:"address"`
}

type FAddress struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "ChartViewer",
		Width:     1024,
		Height:    768,
		Assets:    assets,
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			&FPerson{},
			&FAddress{},
		},
	})

	if err != nil {
		println("Error:", err)
	}
}
