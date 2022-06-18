package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 450
)

var (
	bIsRunning      = true
	BackgroundColor = rl.Color{0x89, 0x00, 0xf2, 0xff}
	FrameCount      int

	GrassSprite rl.Texture2D

	PlayerSprite    rl.Texture2D
	PlayerSrc       rl.Rectangle
	PlayerDest      rl.Rectangle
	PlayerSpeed     float32 = 2.0
	bIsMoving       bool
	PlayerDirection int
	PlayerFrame     int

	bIsBGMPaused    = false
	BackgroundSound rl.Music
	Camera          rl.Camera2D
)

func Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Sprout Land")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	GrassSprite = rl.LoadTexture("resources/tilesets/Grass.png")
	PlayerSprite = rl.LoadTexture("resources/characters/Basic Charakter Spritesheet.png")

	PlayerSrc = rl.NewRectangle(0, 0, 48, 48)
	PlayerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	BackgroundSound = rl.LoadMusicStream("resources/Monolith OST 320/03 Before the Dawn.mp3")

	Camera = rl.NewCamera2D(
		rl.NewVector2(ScreenWidth/2, ScreenHeight/2),
		rl.NewVector2(PlayerDest.X-ScreenWidth/2, PlayerDest.Y-ScreenHeight/2),
		0.0, 1.0)
}

func Quit() {
	rl.UnloadTexture(GrassSprite)
	rl.UnloadTexture(PlayerSprite)
	rl.UnloadMusicStream(BackgroundSound)

	rl.CloseAudioDevice()
	rl.CloseWindow()

}

func Input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		bIsMoving = true
		PlayerDirection = 1
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		bIsMoving = true
		PlayerDirection = 2
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		bIsMoving = true
		PlayerDirection = 0
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		bIsMoving = true
		PlayerDirection = 3
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		bIsBGMPaused = !bIsBGMPaused
	}
}

func DrawScene() {
	rl.DrawTexture(GrassSprite, 0, 0, rl.White)
	rl.DrawTexturePro(PlayerSprite, PlayerSrc, PlayerDest, rl.NewVector2(PlayerDest.Width, PlayerDest.Height), 0, rl.White)
}

func Update() {
	bIsRunning = !rl.WindowShouldClose()

	if bIsMoving {
		if PlayerDirection == 0 {
			PlayerDest.Y += PlayerSpeed
		}
		if PlayerDirection == 1 {
			PlayerDest.Y -= PlayerSpeed
		}
		if PlayerDirection == 2 {
			PlayerDest.X -= PlayerSpeed
		}
		if PlayerDirection == 3 {
			PlayerDest.X += PlayerSpeed
		}

		if FrameCount%10 == 1 {
			PlayerFrame++
		}
	}

	FrameCount++
	PlayerFrame = PlayerFrame % 3

	PlayerSrc.X = PlayerSrc.Width * float32(PlayerFrame)
	PlayerSrc.Y = PlayerSrc.Height * float32(PlayerDirection)

	rl.UpdateMusicStream(BackgroundSound)
	if bIsBGMPaused {
		rl.PauseMusicStream(BackgroundSound)
	} else {
		rl.ResumeMusicStream(BackgroundSound)
	}

	Camera.Target = rl.NewVector2(PlayerDest.X /*-ScreenWidth/2*/, PlayerDest.Y /*-ScreenHeight/2*/)
	bIsMoving = false
	// PlayerDirection = 0
}

func Render() {
	rl.BeginDrawing()
	rl.ClearBackground(BackgroundColor)
	rl.BeginMode2D(Camera)

	DrawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func main() {
	Init()
	for bIsRunning {
		Input()
		Update()
		Render()
	}
	Quit()
}
