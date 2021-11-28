package common

import (
	"math"
)

// CheckCollisions checks if there are any collisions
func CheckCollisions() error {
	for i := 0; i < len(Elements)-1; i++ {
		for j := i + 1; j < len(Elements); j++ {
			for _, c1 := range Elements[i].Collisions {
				for _, c2 := range Elements[j].Collisions {
					if collides(c1, c2) && Elements[i].Active && Elements[j].Active {
						err := Elements[i].Collision(Elements[j])
						if err != nil {
							return err
						}
						err = Elements[j].Collision(Elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) +
		math.Pow(c2.Center.Y-c1.Center.Y, 2))

	return dist <= c1.Radius+c2.Radius
}