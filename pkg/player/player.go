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
		Y: common.ScreenHeight /2.0}

	sr := components.NewSpriteRenderer(player, renderer, "assets/player.bmp")
	player.AddComponent(sr)

	shooter := components.NewKeyboardShooter(player, common.PlayerShotCoolDown)
	player.AddComponent(shooter)
	quitter := components.NewKeyboardQuitter(player)
	player.AddComponent(quitter)

	player.Active = true
	player.Tag = "player"

	return player
}
