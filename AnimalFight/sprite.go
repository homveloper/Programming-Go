package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type FSpriteAnimation struct {
	Image       float32 // used as timer
	ImageCount  int     // number of images in the animation
	FPS         int     // time in seconds for each frame
	ImageIDList []int   // array of Sprite ID
}

type FSprite struct {
	X                   int
	Y                   int
	Width               int
	Height              int
	XOrigin             int
	YOrigin             int
	XScale              float32
	YScale              float32
	ImageCount          int
	AnimationID         int
	AnimationCount      int
	Texture             *rl.Texture2D
	ImagePosList        []rl.Vector2
	SpriteAnimationList []FSpriteAnimation
}

func NewSpriteFromTexture(Texture *rl.Texture2D, Width, Height, XOrigin, YOrigin int) FSprite {
	Sprite := FSprite{
		Width:       Width,
		Height:      Height,
		XOrigin:     XOrigin,
		YOrigin:     YOrigin,
		XScale:      1.0,
		YScale:      1.0,
		Texture:     Texture,
		AnimationID: -1,
	}

	XFrames := int(Texture.Width) / Width
	YFrames := int(Texture.Height) / Height

	Sprite.ImageCount = XFrames * YFrames
	Sprite.ImagePosList = make([]rl.Vector2, XFrames*YFrames)

	for y := 0; y < YFrames; y++ {
		for x := 0; x < XFrames; x++ {
			Sprite.ImagePosList[y*XFrames+x] = rl.NewVector2(float32(x*Width), float32(y*Height))
		}
	}

	return Sprite
}

func (Sprite *FSprite) AddAnimation(Animation FSpriteAnimation) {
	Sprite.SpriteAnimationList = append(Sprite.SpriteAnimationList, Animation)
	Sprite.AnimationCount++
}

func (Sprite *FSprite) Play(AnimationID int, Delta float32) {
	if Sprite.AnimationID != AnimationID {
		Sprite.AnimationID = AnimationID
		Sprite.SpriteAnimationList[AnimationID].Image = 0.0
	} else {
		Sprite.SpriteAnimationList[AnimationID].Image += Delta * float32(Sprite.SpriteAnimationList[AnimationID].FPS)
		if Sprite.SpriteAnimationList[AnimationID].Image > float32(Sprite.SpriteAnimationList[AnimationID].ImageCount) {
			Sprite.SpriteAnimationList[AnimationID].Image -= float32(int(Sprite.SpriteAnimationList[AnimationID].Image))
		}
	}
}

func (Sprite *FSprite) Draw() {
	Animation := &Sprite.SpriteAnimationList[Sprite.AnimationID]
	ImageID := Animation.ImageIDList[int(Animation.Image)]

	SourcePos := Sprite.ImagePosList[ImageID]
	Offset := rl.NewVector2(float32(Sprite.XOrigin)*Sprite.XScale, float32(Sprite.YOrigin)*Sprite.YScale)

	X := float32(Sprite.X) - Offset.X
	Y := float32(Sprite.Y) - Offset.Y
	W := float32(Sprite.Width) * float32(math.Abs(float64(Sprite.XScale)))
	H := float32(Sprite.Height) * float32(math.Abs(float64(Sprite.YScale)))

	Source := rl.NewRectangle(SourcePos.X, SourcePos.Y, float32(Sprite.Width), float32(Sprite.Height))

	rl.DrawTextureRec(*Sprite.Texture, Source)
}
