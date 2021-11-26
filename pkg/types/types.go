package types

type Circle struct {
	Center Vector
	Radius float64
}

type Vector struct {
	X float64
	Y float64
}

var Delta float64

const (
	ScreenWidth  = 600
	ScreenHeight = 800

	TargetTicksPerSecond = 60
)
