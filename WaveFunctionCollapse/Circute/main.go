package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	DIMENSION int = 25
)

var (
	tiles      []Tile
	tileImages []rl.Texture2D
	grid       []Cell
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		Input()
		Update()
		rl.BeginDrawing()
		DrawScene()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func Input() {

}

func Update() {

}

func DrawScene() {
	rl.ClearBackground(rl.RayWhite)
}
