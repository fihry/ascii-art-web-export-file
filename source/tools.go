package ascii

import (
	"bufio"
	"os"
	"strings"
)

func IsEmpty(slice []string) bool {
	for _, v := range slice {
		if v != "" {
			return false
		}
	}
	return true
}

func IsBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy" || banner == "enigma" || banner == "nirvana"
}

func IsPrintable(s string) bool {
	for _, char := range s {
		if (char < 32 || char > 126) && char != '\n' && char != '\r' && char != '\t' && char != ' ' && char != '\v' && char != '\f' {
			return false
		}
	}
	return true
}

func LoadAscii(Banner string) (slice []string) {
	f, _ := os.Open("banners/" + Banner + ".txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		slice = append(slice, line)
	}
	return
}

// ka t7ayed spaces li 9bel newline
func RemoveTrailingSpaces(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}
	return strings.Join(lines, "\n")
}
