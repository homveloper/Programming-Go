package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 800
	ScreenHeight = 450
)

var (
	bIsRunning      = true
	BackgroundColor = rl.NewColor(147, 211, 196, 255)

	GrassSprite  rl.Texture2D
	PlayerSprite rl.Texture2D

	PlayerSrc  rl.Rectangle
	PlayerDest rl.Rectangle

	PlayerSpeed float32 = 2.0

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
		PlayerDest.Y -= PlayerSpeed
	}

	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		PlayerDest.X -= PlayerSpeed
	}

	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		PlayerDest.Y += PlayerSpeed
	}

	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		PlayerDest.X += PlayerSpeed
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

	rl.UpdateMusicStream(BackgroundSound)
	if bIsBGMPaused {
		rl.PauseMusicStream(BackgroundSound)
	} else {
		rl.ResumeMusicStream(BackgroundSound)
	}

	Camera.Target = rl.NewVector2(PlayerDest.X /*-ScreenWidth/2*/, PlayerDest.Y /*-ScreenHeight/2*/)
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
