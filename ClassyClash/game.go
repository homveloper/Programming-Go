package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	m_ScreenWidth     int32
	m_ScreenHeight    int32
	m_Title           string
	m_BackGroundColor rl.Color
	m_TargetFPS       int32
	m_bIsRunning      bool
}

func NewGame() *Game {
	return &Game{}
}

func (game Game) Init() {
	rl.InitWindow(game.m_ScreenWidth, game.m_ScreenHeight, game.m_Title)
	rl.SetExitKey(0)
	rl.SetTargetFPS(game.m_TargetFPS)

}

func (game Game) Quit() {
	rl.CloseWindow()
}

func (game Game) Update() {
	game.m_bIsRunning = !rl.WindowShouldClose()

}

func (game Game) Input() {

}

func (game Game) RenderScene() {

}

func (game Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(game.m_BackGroundColor)

	game.RenderScene()

	rl.EndDrawing()
}
