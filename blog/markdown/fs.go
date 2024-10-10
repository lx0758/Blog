package markdown

import (
	"regexp"
	"strings"
)

type FragmentSource interface {
	GetFragment(string) string
}

var fragmentRegexp = regexp.MustCompile(`<fragment include="(.+?)"/>`)

func getFinalSource(source string, fs FragmentSource) []byte {
	builder := strings.Builder{}

	lines := strings.Split(source, "\n")
	for _, line := range lines {
		line := strings.TrimRight(line, " \t\r\n")
		if !fragmentRegexp.MatchString(line) {
			builder.WriteString(line)
			builder.WriteString("\n")
			continue
		}

		matches := fragmentRegexp.FindStringSubmatch(line)
		fragmentKey := matches[1]
		fragmentContent := fs.GetFragment(fragmentKey)
		if fragmentContent == "" {
			continue
		}
		builder.WriteString(fragmentContent)
		builder.WriteString("\n")
	}

	return []byte(builder.String())
}
