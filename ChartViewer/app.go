package main

import (
	"context"
	"fmt"

	"github.com/xuri/excelize/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenFile(filePath string) {
	file, err := excelize.OpenFile(filePath)

	if err != nil {

	}

	fmt.Println(file)
}
