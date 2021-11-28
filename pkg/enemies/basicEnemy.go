package enemies

import (
	"fmt"

	"softcraft/pkg/common"
	"softcraft/pkg/components"

	"github.com/veandco/go-sdl2/sdl"
)

// NewBasicEnemy returns a new Enemy (not implemented yet)
func NewBasicEnemy(renderer *sdl.Renderer, x, y float64) *common.Element {
	basicEnemy := &common.Element{}

	basicEnemy.Position = common.Vector{X: x, Y: y}
	basicEnemy.Rotation = 180

	idleSequence, err := components.NewSequence("assets/sprites/basic_enemy/idle", 5, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence failed: %v", err))
	}
	destroySequence, err := components.NewSequence("assets/sprites/basic_enemy/destroy", 15, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence failed: %v", err))
	}
	sequences := map[string]*components.Sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := components.NewAnimator(basicEnemy, sequences, "idle")
	basicEnemy.AddComponent(animator)

	vtb := components.NewVulnerableToBullets(basicEnemy)
	basicEnemy.AddComponent(vtb)

	col := common.Circle{
		Center: basicEnemy.Position,
		Radius: 38,
	}
	basicEnemy.Collisions = append(basicEnemy.Collisions, col)

	basicEnemy.Active = true

	return basicEnemy
}
