package player

import (
	"softcraft/pkg/assetManager"
	"softcraft/pkg/common"
	"softcraft/pkg/components"
)

// NewPlayer returns a new player
func NewPlayer(am *assetManager.AssetManager) *common.Element {
	player := &common.Element{}

	player.Position = common.Vector{
		X: common.ScreenWidth / 2.0,
		Y: common.ScreenHeight /2.0}

	player.Active = true
	player.Tag = "player"

	// Get player asset and create a sprite renderer component
	tex := am.GetAsset(assetManager.AssetStringPlayer)
	sr := components.NewSpriteRenderer(player, tex)
	player.AddComponent(sr)

	// Create a quitter component
	quitter := components.NewKeyboardQuitter(player)
	player.AddComponent(quitter)

	return player
}
