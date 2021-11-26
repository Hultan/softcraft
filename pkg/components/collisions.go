package components

import (
	"math"

	"softcraft/pkg/common"
)

func collides(c1, c2 common.Circle) bool {
	dist := math.Sqrt(math.Pow(c2.Center.X-c1.Center.X, 2) +
		math.Pow(c2.Center.Y-c1.Center.Y, 2))

	return dist <= c1.Radius+c2.Radius
}

func CheckCollisions() error {
	for i := 0; i < len(common.Elements)-1; i++ {
		for j := i + 1; j < len(common.Elements); j++ {
			for _, c1 := range common.Elements[i].Collisions {
				for _, c2 := range common.Elements[j].Collisions {
					if collides(c1, c2) && common.Elements[i].Active && common.Elements[j].Active {
						err := common.Elements[i].Collision(common.Elements[j])
						if err != nil {
							return err
						}
						err = common.Elements[j].Collision(common.Elements[i])
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
