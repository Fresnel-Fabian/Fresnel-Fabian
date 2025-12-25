package raindrops
import "strconv"

func Convert(number int) string {
	var output string = "";
    var divisible bool = false;
    if (number % 3 == 0) {
        output += "Pling";
        divisible = true
    }
    if (number % 5 == 0) {
        output += "Plang";
        divisible = true
    }
    if (number % 7 == 0) {
        output += "Plong";
        divisible = true
    }
    if divisible {
        return output
    }
    return strconv.Itoa(number)
}
