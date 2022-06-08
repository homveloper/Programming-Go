package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth          = 600
	ScreenHeight         = 800
	GridWidth            = 40
	GridHeight           = 40
	ScreenScale  float32 = ScreenHeight / GridWidth
	MaxParticles         = GridWidth * GridHeight
)

var (
	ScreenCenter rl.Vector2 = rl.NewVector2((ScreenWidth-GridWidth*ScreenScale)/2, (ScreenHeight-GridHeight*ScreenScale)/2)
)

type EParticleType uint8

const (
	EPT_None EParticleType = iota
	EPT_Stone
	EPT_Dust
	EPT_Water
)

var (
	ColorDark  = rl.DarkGray
	ColorStone = rl.NewColor(150, 150, 150, 255)
	ColorDust  = rl.Brown
)

type FParticle struct {
	m_Position rl.Vector2
	m_Type     EParticleType
}

func InitWindow() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "SandBox")
	rl.SetTargetFPS(60)

	InitGame()
}

func CloseWindow() {
	CloseGame()
}

func InitGame() {

	Render()
}

func UpdateGame() {

}

func UpdateScene() {

}

func Render() {
	for rl.WindowShouldClose() {
		UpdateGame()
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		UpdateScene()

		rl.EndDrawing()
	}

	CloseWindow()
}

func CloseGame() {

}

func main() {
	InitWindow()
}
