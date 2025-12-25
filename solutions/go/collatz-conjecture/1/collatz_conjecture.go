package collatzconjecture
import "fmt"


func CollatzConjecture(n int) (int, error) {
	var step int = 0;
    if n < 1 {
        return 0, fmt.Errorf("input must be a positive integer")
    }
    for {
        if n == 1 {
            break;
        }
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = 3 * n + 1
        }
        step += 1;
    }
    return step, nil
}
