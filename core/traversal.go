package core

import (
	"math"

	"github.com/martinlaursen97/sand/maths"
	"github.com/martinlaursen97/sand/utils"
)

type FunctionInput struct {
	Func func(...any) any
	Args []any
}

// Functions using this function must contain,
// as the first two arguments, x and y coordinates.
// The rest of the arguments are optional.
func TraversePathAndApplyFunc(curr, next maths.Vector, fi FunctionInput) any {
	if curr.X == next.X && curr.Y == next.Y {
		return nil
	}

	vx := next.X - curr.X
	vy := next.Y - curr.Y

	length := math.Sqrt(vx*vx + vy*vy)

	xIncrement, yIncrement := vx/length, vy/length
	numPoints := int(length)

	for i := 1; i <= numPoints; i++ {
		dx := curr.X + xIncrement*float64(i)
		dy := curr.Y + yIncrement*float64(i)

		dy = utils.RoundYCoordinate(dy)

		if utils.WithinBounds(dx, dy) {
			if len(fi.Args) == 0 {
				result := fi.Func(dx, dy)
				if result != nil {
					return result
				}
			} else {
				args := append([]any{dx, dy}, fi.Args...)
				result := fi.Func(args...)
				if result != nil {
					return result
				}
			}
		}
	}

	return nil
}
