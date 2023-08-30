package scan_modules

import (
	"bufio"
	"github.com/MrMelon54/iconic-data-convert/json"
	"log"
)

func ScanIconicModules(l *log.Logger, s *bufio.Scanner, repoRawJson json.KtaneRawJson) map[string]int {
	order := make(map[string]int)
	n := 0
	for s.Scan() {
		a := s.Text()
		if a == "" {
			continue
		}
		b := repoRawJson.ConvertDisplayNameToID(a)
		if b == "" {
			l.Printf("[WARNING] Failed to find '%s' in module list\n", a)
		}
		order[b] = n
		n++
	}
	if err := s.Err(); err != nil {
		l.Fatal("[ERROR] Failed to parse iconic modules order: ", err)
	}
	return order
}
