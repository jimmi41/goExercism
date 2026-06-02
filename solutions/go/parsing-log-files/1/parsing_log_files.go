package parsinglogfiles
import "regexp"

var validLine =
    regexp.MustCompile(
        `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`,
    )
func IsValidLine(text string) bool {
	return validLine.MatchString(text)
}

func SplitLogLine(text string) []string {

     re := regexp.MustCompile(`<[~*=-]*>`)

    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
     re := regexp.MustCompile(`(?i)".*password.*"`)

    count := 0

    for i := 0; i < len(lines); i++ {

        if re.MatchString(lines[i]) {
            count++
        }
    }

    return count
}

func RemoveEndOfLineText(text string) string {
	
    re := regexp.MustCompile(`end-of-line\d+`)

    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)

    for i := 0; i < len(lines); i++ {

        match := re.FindStringSubmatch(lines[i])

        if len(match) > 1 {

            username := match[1]

            lines[i] =
                "[USR] " +
                    username +
                    " " +
                    lines[i]
        }
    }

    return lines
}
