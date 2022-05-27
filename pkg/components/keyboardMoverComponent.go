package components

import (
	"softcraft/pkg/assetManager"
	"softcraft/pkg/common"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container   *common.Element
	speed       float64
	movingSpeed float64
	world       *World
}

func NewKeyboardMover(container *common.Element, speed float64, world *World) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		world:     world,
	}
}

func (m *keyboardMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (m *keyboardMover) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	if m.movingSpeed == 0 {
		m.movingSpeed = m.speed * common.Delta
	}

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if m.canMove(-1, 0) {
			m.world.Position.X -= m.movingSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if m.canMove(1, 0) {
			m.world.Position.X += m.movingSpeed
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		if m.canMove(0, -1) {
			m.world.Position.Y -= m.movingSpeed
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if m.canMove(0, 1) {
			m.world.Position.Y += m.movingSpeed
		}
	}

	return nil
}

func (m *keyboardMover) OnCollision(_ *common.Element) error {
	return nil
}

func (m *keyboardMover) canMove(dx, dy float64) bool {
	x := int64((m.world.Position.X + dx*m.movingSpeed) / common.BlockWidth)
	y := int64((m.world.Position.Y + dy*m.movingSpeed) / common.BlockHeight)
	if m.world.data[y][x] == assetManager.AssetMapWater ||
		m.world.data[y][x] == assetManager.AssetMapGround {
		return false
	}
	return true
}
