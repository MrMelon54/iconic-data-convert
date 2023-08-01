package scan_script

import (
	"bufio"
	"log"
	"regexp"
	"strings"
)

var reScript = regexp.MustCompile(`^{ "([^"]+)", iconicData\._([^ ]+) },$`)

func ScanIconicScript(l *log.Logger, s *bufio.Scanner) map[string]string {
	moduleMap := make(map[string]string)
	var running, ignoreLines bool
	for s.Scan() {
		a := strings.TrimSpace(s.Text())
		if ignoreLines {
			if a == "*/" {
				ignoreLines = false
			}
			continue
		}
		if !running {
			if strings.Contains(a, "private OrderedDictionary ModuleList = new OrderedDictionary {") {
				running = true
			}
			continue
		}
		if strings.Contains(a, "{ string.Empty, iconicData.BlankModule }") {
			break
		}
		if a == "" {
			continue
		}
		if a == "/*" {
			ignoreLines = true
			continue
		}

		b := reScript.FindStringSubmatch(a)
		if b == nil {
			l.Printf("[WARNING] Line didn't match: \"%s\"\n", a)
			continue
		}
		// map the variable to the string name
		moduleMap[b[2]] = b[1]
	}
	if err := s.Err(); err != nil {
		l.Fatal("[ERROR] Failed to parse iconic script: ", err)
	}
	return moduleMap
}
