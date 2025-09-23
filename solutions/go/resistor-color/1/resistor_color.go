package resistorcolor

var colorMapResistance = map[string]int{
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9,
}
var colorMapList = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
// Colors returns the list of all colors.
func Colors() []string {
	return colorMapList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colorMapResistance[color]
}
