package main

import (
	"math"

	"softcraft/pkg/components"
	"softcraft/pkg/types"
)

func collides(c1, c2 types.Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) +
		math.Pow(c2.Center.Y-c1.Center.Y, 2))

	return dist <= c1.Radius+c2.Radius
}

func checkCollisions() error {
	for i := 0; i < len(components.Elements)-1; i++ {
		for j := i + 1; j < len(components.Elements); j++ {
			for _, c1 := range components.Elements[i].Collisions {
				for _, c2 := range components.Elements[j].Collisions {
					if collides(c1, c2) && components.Elements[i].Active && components.Elements[j].Active {
						err := components.Elements[i].Collision(components.Elements[j])
						if err != nil {
							return err
						}
						err = components.Elements[j].Collision(components.Elements[i])
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
