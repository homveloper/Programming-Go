package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SEED = 500
)

var HASH = []int{208, 34, 231, 213, 32, 248, 233, 56, 161, 78, 24, 140, 71, 48, 140, 254, 245, 255, 247, 247, 40,
	185, 248, 251, 245, 28, 124, 204, 204, 76, 36, 1, 107, 28, 234, 163, 202, 224, 245, 128, 167, 204,
	9, 92, 217, 54, 239, 174, 173, 102, 193, 189, 190, 121, 100, 108, 167, 44, 43, 77, 180, 204, 8, 81,
	70, 223, 11, 38, 24, 254, 210, 210, 177, 32, 81, 195, 243, 125, 8, 169, 112, 32, 97, 53, 195, 13,
	203, 9, 47, 104, 125, 117, 114, 124, 165, 203, 181, 235, 193, 206, 70, 180, 174, 0, 167, 181, 41,
	164, 30, 116, 127, 198, 245, 146, 87, 224, 149, 206, 57, 4, 192, 210, 65, 210, 129, 240, 178, 105,
	228, 108, 245, 148, 140, 40, 35, 195, 38, 58, 65, 207, 215, 253, 65, 85, 208, 76, 62, 3, 237, 55, 89,
	232, 50, 217, 64, 244, 157, 199, 121, 252, 90, 17, 212, 203, 149, 152, 140, 187, 234, 177, 73, 174,
	193, 100, 192, 143, 97, 53, 145, 135, 19, 103, 13, 90, 135, 151, 199, 91, 239, 247, 33, 39, 145,
	101, 120, 99, 3, 186, 86, 99, 41, 237, 203, 111, 79, 220, 135, 158, 42, 30, 154, 120, 67, 87, 167,
	135, 176, 183, 191, 253, 115, 184, 21, 233, 58, 129, 233, 142, 39, 128, 211, 118, 137, 139, 255,
	114, 20, 218, 113, 154, 27, 127, 246, 250, 1, 8, 198, 250, 209, 92, 222, 173, 21, 88, 102, 219}

func Noise2(x, y int) int {
	tmp := HASH[(y+SEED)%256]
	return HASH[(tmp+x)%256]
}

func LinearInterp(x, y, s float32) float32 {
	return x + s*(y-x)
}

func SmoothInterp(x, y, s float32) float32 {
	return LinearInterp(x, y, s*s*(3-2*s))
}

func Noise2D(x, y float32) float32 {
	xInt := int(x)
	yInt := int(y)
	xFrac := x - float32(xInt)
	yFrac := y - float32(yInt)

	s := Noise2(xInt, yInt)
	t := Noise2(xInt+1, yInt)
	u := Noise2(xInt, yInt+1)
	v := Noise2(xInt+1, yInt+1)

	low := SmoothInterp(float32(s), float32(t), xFrac)
	high := SmoothInterp(float32(u), float32(v), xFrac)

	return SmoothInterp(low, high, yFrac)
}

func Perlin2D(x, yy, frequency float32, depth int) float32 {
	xa := x * frequency
	ya := y * frequency
	amplitude := float32(1.0)
	fin := float32(0.0)
	div := float32(0.0)

	for i := 0; i < depth; i++ {
		div += 256 * amplitude
		fin += Noise2D(xa, ya) * amplitude
		amplitude /= 2
		xa *= 2
		ya *= 2
	}

	return fin / div
}

func main() {
	rl.InitWindow(600, 600, "3D Terrain Noise")

	var camera rl.Camera = rl.NewCamera3D(rl.NewVector3(30.0, 30.0, 30.0),
		rl.NewVector3(0.0, 5.0, 0.0),
		rl.NewVector3(0.0, 1.0, 0.0),
		45.0,
		rl.CameraPerspective)

	frequency := 1
	depth := 1
	scale := float32(1.0)
	bUpdate := true
	bIsCube := true
	bInfo := false
	bSwitchCamera := false
	bIsFP := true

	var arrayMap [100][100]float32

	rl.SetCameraMode(camera, rl.CameraFirstPerson)
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera)

		if rl.IsKeyPressed(rl.KeyUp) {
			frequency++
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyDown) {
			frequency--
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyLeft) {
			depth++
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			depth--
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyN) {
			scale += 0.1
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyN) {
			scale -= 0.1
			bUpdate = true
		}

		if rl.IsKeyPressed(rl.KeyE) {
			bIsCube = !bIsCube
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			bSwitchCamera = true
		}

		if rl.IsKeyPressed(rl.KeyI) {
			bInfo = !bInfo
		}

		if bUpdate {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					arrayMap[x][y] = Perlin2D(float32(x)*scale, float32(y)*scale, float32(frequency)/10.0, depth) * 10.0
				}
			}

			bUpdate = false
		}

		if bSwitchCamera {
			if bIsFP {
				rl.SetCameraMode(camera, rl.CameraOrbital)
				bIsFP = false
			} else {
				rl.SetCameraMode(camera, rl.CameraFirstPerson)
				bIsFP = true
			}

			camera.Target = rl.NewVector3(0.0, 5.0, 0.0)
			camera.Position = rl.NewVector3(30.0, 30.0, 30.0)
			bSwitchCamera = false
		}
		camera.Position.Y = 30.0

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(camera)

		if !bIsCube {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					rl.DrawSphere(rl.NewVector3(float32(x-50)/2.0, arrayMap[x][y], float32(y-25)/2.0), 0.05, rl.Red)
				}
			}
		} else {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					rl.DrawCube(rl.NewVector3(float32(x-50)/2.0, arrayMap[x][y], float32(y-25)/2.0), 0.5, arrayMap[x][y], 0.5, rl.NewColor(uint8(arrayMap[x][y]/10.0*255), uint8(arrayMap[x][y]/10.0*255), uint8(arrayMap[x][y]/10.0*255), 255))
				}
			}
		}

		rl.DrawGrid(25, 5.0)
		rl.EndMode3D()

		for x := 0; x < 100; x++ {
			for y := 0; y < 100; y++ {
				rl.DrawPixel(int32(600-100+x), int32(y), rl.NewColor(255, 255, 255, uint8(arrayMap[x][y]/10*255)))
			}
		}

		rl.EndDrawing()
	}
}
