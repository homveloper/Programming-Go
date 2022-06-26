package main

func main() {
	game := NewGame(600, 600, "Classy Clash")

	game.Init()
	for game.m_bIsRunning {
		game.Input()
		game.Update()
		game.Render()
	}
	game.Quit()
}
