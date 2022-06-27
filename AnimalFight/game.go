package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EGameScene uint8

const (
	EGS_Title EGameScene = iota
	EGS_Gameplay
	EGS_Ending
)

const (
	ScreenWidth  = 640
	ScreenHeight = 640
)

var (
	ScreenCenter     rl.Vector2 = rl.NewVector2(ScreenWidth/2, ScreenHeight/2)
	bIsRunning                  = true
	ThemeTextColor              = rl.RayWhite
	BackgroundColor             = color.RGBA{0x89, 0x00, 0xf2, 0xff}
	CurrentGameScene            = EGS_Title
	Player                      = &UPlayer{m_Position: ScreenCenter, m_Speed: 4.0}
	Camera           rl.Camera2D
)

func Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Sprout Land")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	// Load Asset
	Player.m_Sprite = rl.LoadTexture("resources/Player/player.png")
	Camera = rl.NewCamera2D(ScreenCenter,
		rl.NewVector2(Player.m_Position.X+float32(Player.m_Sprite.Width), Player.m_Position.Y+float32(Player.m_Sprite.Height)),
		0.0, 1.0)
}

func Quit() {
	// Close Loaded Asset
	rl.UnloadTexture(Player.m_Sprite)
	rl.CloseWindow()
}

func Update() {
	bIsRunning = !rl.WindowShouldClose()

}

func Input() {

	switch CurrentGameScene {
	case EGS_Title:
		if rl.IsGestureDetected(rl.GestureTap) {
			CurrentGameScene = EGS_Gameplay
		}
	case EGS_Gameplay:
		if rl.IsKeyDown(rl.KeyW) {
			Player.UpdatePosition(rl.NewVector2(0.0, -Player.m_Speed))
		}
		if rl.IsKeyDown(rl.KeyS) {
			Player.UpdatePosition(rl.NewVector2(0.0, Player.m_Speed))
		}
		if rl.IsKeyDown(rl.KeyA) {
			Player.UpdatePosition(rl.NewVector2(-Player.m_Speed, 0.0))
		}
		if rl.IsKeyDown(rl.KeyD) {
			Player.UpdatePosition(rl.NewVector2(Player.m_Speed, 0.0))
		}
	}
}

func DrawScene() {
	switch CurrentGameScene {
	case EGS_Title:
		rl.DrawText("Title Screen", 20, 20, 40, ThemeTextColor)
		rl.DrawText("Press the Screen", int32(ScreenCenter.X-100), int32(ScreenCenter.Y), 25, ThemeTextColor)
		Camera.Target = ScreenCenter
	case EGS_Gameplay:
		Camera.Target = rl.NewVector2(Player.m_Position.X+float32(Player.m_Sprite.Width), Player.m_Position.Y+float32(Player.m_Sprite.Height))

		rl.DrawText("Gameplay Screen", 20, 20, 40, ThemeTextColor)
		rl.DrawTextureEx(Player.m_Sprite, Player.m_Position, 0.0, 2.0, rl.White)
	case EGS_Ending:
		rl.DrawText("Ending Screen", 20, 20, 40, ThemeTextColor)
	}
}

func Render() {
	rl.BeginDrawing()
	rl.ClearBackground(BackgroundColor)
	rl.BeginMode2D(Camera)

	DrawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}
