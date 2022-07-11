package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	m_Image rl.Texture2D
	m_Edges []int
	m_Idx   int
	m_Up    []int
	m_Right []int
	m_Down  []int
	m_Left  []int
}

func CompareEdge(a, b int) bool {
	return a == b
}

func (t *Tile) Analyze(tiles []Tile) {
	for i := 0; i < len(tiles); i++ {
		tile := tiles[i]

		if CompareEdge(t.m_Edges[0], tile.m_Edges[2]) {
			t.m_Up = append(t.m_Up, i)
		}

		if CompareEdge(t.m_Edges[1], tile.m_Edges[3]) {
			t.m_Right = append(t.m_Right, i)
		}

		if CompareEdge(t.m_Edges[2], tile.m_Edges[0]) {
			t.m_Down = append(t.m_Down, i)
		}

		if CompareEdge(t.m_Edges[3], tile.m_Edges[1]) {
			t.m_Left = append(t.m_Left, i)
		}
	}
}

func NewTile(image rl.Texture2D, edges []int, idx int) *Tile {
	pTile := new(Tile)
	pTile.m_Image = image
	pTile.m_Edges = edges
	pTile.m_Idx = idx
	return pTile
}
