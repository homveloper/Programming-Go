package main

import rl "github.com/gen2brain/raylib-go/raylib"

type UPlayer struct {
	m_Position rl.Vector2
	m_Sprite   rl.Texture2D
	m_Speed    float32
}

func (Player *UPlayer) UpdatePosition(Offset rl.Vector2) {
	Player.m_Position = rl.Vector2Add(Player.m_Position, Offset)
}
