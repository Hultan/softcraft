package main

import (
	"fmt"

	"softcraft/pkg/components"
	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) *components.Element {
	basicEnemy := &components.Element{}

	basicEnemy.Position = types.Vector{X: x, Y: y}
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

	col := types.Circle{
		Center: basicEnemy.Position,
		Radius: 38,
	}
	basicEnemy.Collisions = append(basicEnemy.Collisions, col)

	basicEnemy.Active = true

	return basicEnemy
}
