package main

import (
	"time"

	"softcraft/pkg/components"
	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSize = 105

	playerShotCoolDown = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) *components.Element {
	player := &components.Element{}

	player.Position = types.Vector{
		X: types.ScreenWidth / 2.0,
		Y: types.ScreenHeight - playerSize/2.0}

	sr := components.NewSpriteRenderer(player, renderer, "assets/sprites/player.bmp")
	player.AddComponent(sr)

	mover := components.NewKeyboardMover(player, 5)
	player.AddComponent(mover)

	shooter := components.NewKeyboardShooter(player, playerShotCoolDown)
	player.AddComponent(shooter)

	player.Active = true

	return player
}
