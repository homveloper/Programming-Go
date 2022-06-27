package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gameMap    rl.Texture2D
	gameMapPos rl.Vector2 = rl.NewVector2(0, 0)

	playerDir rl.Vector2 = rl.NewVector2(0, 0)
	moveSpeed float32    = 3.0
)

type Game struct {
	m_ScreenWidth     int32
	m_ScreenHeight    int32
	m_Title           string
	m_BackGroundColor rl.Color
	m_TargetFPS       int32
	m_bIsRunning      bool
}

func NewGame(width, height int32, title string) *Game {
	game := &Game{}
	game.m_ScreenWidth = width
	game.m_ScreenHeight = height
	game.m_Title = title
	game.m_BackGroundColor = rl.NewColor(0x89, 0x00, 0xf2, 0xff)
	game.m_TargetFPS = 60
	game.m_bIsRunning = true

	return game
}

func (game Game) Init() {
	rl.InitWindow(game.m_ScreenWidth, game.m_ScreenHeight, game.m_Title)
	rl.SetExitKey(0)
	rl.SetTargetFPS(game.m_TargetFPS)

	gameMap = rl.LoadTexture("resources/TileMap/WolrdMap.bmp")
}

func (game Game) Quit() {
	rl.CloseWindow()
}

func (game Game) Update() {
	game.m_bIsRunning = !rl.WindowShouldClose()

}

func (game Game) Input() {

	dir := rl.NewVector2(0, 0)

	if rl.IsKeyDown(rl.KeyA) {
		dir.X -= 1.0
	}
	if rl.IsKeyDown(rl.KeyD) {
		dir.X += 1.0
	}
	if rl.IsKeyDown(rl.KeyW) {
		dir.Y -= 1.0
	}
	if rl.IsKeyDown(rl.KeyS) {
		dir.Y += 1.0
	}
	if rl.Vector2Length(dir) != 0.0 {
		gameMapPos = rl.Vector2Subtract(gameMapPos, rl.Vector2Scale(rl.Vector2Normalize(dir), moveSpeed))
	}
}

func (game Game) RenderScene() {
	rl.DrawTextureEx(gameMap, gameMapPos, 0.0, 1.0, rl.White)
}

func (game Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(game.m_BackGroundColor)
	game.RenderScene()
	rl.EndDrawing()
}
