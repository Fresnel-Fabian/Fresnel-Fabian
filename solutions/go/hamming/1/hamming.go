package hamming
import "errors"

func Distance(a, b string) (int, error) {
    var hammingDistance int = 0;
    if len(a) != len(b) {
        return 0, errors.New("len of strand is different")
    }
	for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            hammingDistance += 1
        }
    }
    return hammingDistance, nil
}
