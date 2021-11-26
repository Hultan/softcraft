package player

import (
	"softcraft/pkg/common"
	"softcraft/pkg/components"

	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayer(renderer *sdl.Renderer) *common.Element {
	player := &common.Element{}

	player.Position = common.Vector{
		X: common.ScreenWidth / 2.0,
		Y: common.ScreenHeight - common.PlayerSize/2.0}

	sr := components.NewSpriteRenderer(player, renderer, "assets/sprites/player.bmp")
	player.AddComponent(sr)

	mover := components.NewKeyboardMover(player, 5)
	player.AddComponent(mover)

	shooter := components.NewKeyboardShooter(player, common.PlayerShotCoolDown)
	player.AddComponent(shooter)

	player.Active = true

	return player
}
