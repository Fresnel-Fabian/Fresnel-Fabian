package luhn
import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "");
    if len(id) <= 1 {
        return false
    }
    var checksum int = 0
    double := false
    for i := len(id) - 1; i >= 0; i-- {
        c := id[i]
        if c < '0' || c > '9' {
            return false
        }
        d := int(c - '0')
        if double {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }
        checksum += d
        double = !double
    }
    return checksum % 10 == 0
}
