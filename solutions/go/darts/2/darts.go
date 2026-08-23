package darts

import (
    "math"
)
func getWithInCircle(x, y float64) func(float64) bool {
    var distance float64 = math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
    return func(circleRadius float64) bool {
        return distance <= circleRadius
    }
}

func Score(x, y float64) int {
	withInCircle := getWithInCircle(x, y)
    if withInCircle(1.0) {
        return 10
    } else if withInCircle(5.0) {
        return 5
    } else if withInCircle(10.0) {
        return 1
    } else {
        return 0
    }
}
