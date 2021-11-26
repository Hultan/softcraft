package components

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardShooter struct {
	container *Element
	coolDown  time.Duration
	lastShot  time.Time
}

func NewKeyboardShooter(container *Element, coolDown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		coolDown:  coolDown}
}

func (shooter *keyboardShooter) OnDraw(_ *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) OnUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := shooter.container.Position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.coolDown {
			shooter.shoot(pos.X+25, pos.Y-20)
			shooter.shoot(pos.X-25, pos.Y-20)

			shooter.lastShot = time.Now()
		}
	}

	return nil
}

func (shooter *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.Active = true
		bul.Position.X = x
		bul.Position.Y = y
		bul.Rotation = 270 * (math.Pi / 180)
	}
}

func (shooter *keyboardShooter) OnCollision(_ *Element) error {
	return nil
}
