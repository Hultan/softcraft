package components

import (
	"math"

	"softcraft/pkg/types"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *Element
	speed     float64
}

func newBulletMover(container *Element) *bulletMover {
	return &bulletMover{container: container}
}

func (mover *bulletMover) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) OnUpdate() error {
	c := mover.container

	c.Position.X += bulletSpeed * math.Cos(c.Rotation) * types.Delta
	c.Position.Y += bulletSpeed * math.Sin(c.Rotation) * types.Delta

	if c.Position.X > types.ScreenWidth || c.Position.X < 0 ||
		c.Position.Y > types.ScreenHeight || c.Position.Y < 0 {
		c.Active = false
	}

	c.Collisions[0].Center = c.Position

	return nil
}

func (mover *bulletMover) OnCollision(_ *Element) error {
	mover.container.Active = false
	return nil
}
