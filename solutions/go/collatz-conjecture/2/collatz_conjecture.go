package collatzconjecture
import "errors"


func CollatzConjecture(n int) (int, error) {
	var step int = 0;
    if n < 1 {
        return 0, errors.New("input must be a positive integer")
    }
    for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = 3 * n + 1
        }
        step += 1;
    }
    return step, nil
}
