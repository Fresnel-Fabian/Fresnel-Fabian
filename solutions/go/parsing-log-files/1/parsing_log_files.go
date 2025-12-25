package parsinglogfiles
import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`);
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`);
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    var count int = 0;
	re := regexp.MustCompile(`"(?i:.*password.*)"`);
    for _, line := range(lines) {
        if (re.MatchString(line)) {
            count += 1;
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`);
    return re.ReplaceAllString(text, "");
}

func TagWithUserName(lines []string) []string {
	var output []string;
    re := regexp.MustCompile(`User\s+([\w]+)`)
    for _, line := range(lines) {
        sl := re.FindStringSubmatch(line);
        if (sl != nil) {
            taggedLine := fmt.Sprintf("[USR] %s %s", sl[1], line);
            output = append(output, taggedLine);
        } else {
            output = append(output, line)
        }
    }
    return output;
}
