package logs
import "unicode/utf8"

var applicationMap = map[rune]string {
    '❗':  "recommendation",
    '🔍': "search",
    '☀': "weather",
}
// Application identifies the application emitting the given log.
func Application(log string) string {
    for _, char := range log {
        if app, ok := applicationMap[char]; ok {
            return app
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	output := []rune(log)
    for i, char := range output {
        if (char == oldRune) {
            output[i] = newRune
        }
    }
    return string(output);
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log);
}
