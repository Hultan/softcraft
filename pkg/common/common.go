package common

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth  = 900
	ScreenHeight = 600

	CanvasWidth = 30
	CanvasHeight = 20

	BlockWidth = 32
	BlockHeight = 32

	TargetTicksPerSecond = 60

	BulletSpeed = 10
)


type Component interface {
	OnUpdate() error
	OnDraw(renderer *sdl.Renderer) error
	OnCollision(other *Element) error
}

type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	Tag        string
	Collisions []Circle
	Components []Component
}

type Circle struct {
	Center Vector
	Radius float64
}

type Vector struct {
	X float64
	Y float64
}

var Elements []*Element
var Delta float64
