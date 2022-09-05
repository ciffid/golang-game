package main

import (
	"game/assets"
	"game/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	assets.Init()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
