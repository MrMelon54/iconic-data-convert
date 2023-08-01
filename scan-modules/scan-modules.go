package scan_modules

import (
	"bufio"
	"log"
)

func ScanIconicModules(l *log.Logger, s *bufio.Scanner) map[string]int {
	order := make(map[string]int)
	n := 0
	for s.Scan() {
		a := s.Text()
		if a == "" {
			continue
		}
		order[a] = n
		n++
	}
	if err := s.Err(); err != nil {
		l.Fatal("[ERROR] Failed to parse iconic modules order: ", err)
	}
	return order
}
