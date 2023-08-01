package scan_data

import (
	"bufio"
	"github.com/MrMelon54/iconic-data-convert/json"
	"log"
	"regexp"
	"strings"
)

var (
	reData  = regexp.MustCompile(`^public static string\[] _([^ ]+) = {\s*"([^"]+)", ([^}]+)};$`)
	rePart  = regexp.MustCompile(`"([^"]+)",?\s*`)
	reClean = strings.NewReplacer("░", " ", "▒", " ", "▓", " ", "█", " ", "═", " ", "║", " ")
)

func ScanIconicData(l *log.Logger, s *bufio.Scanner, moduleMap map[string]string, order map[string]int) []json.IconicModule {
	modules := make([]json.IconicModule, 0, len(order))
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
			if strings.Contains(a, "public static string[] BlankModule") {
				running = true
			}
			continue
		}
		if a == "}" {
			break
		}
		if a == "" {
			continue
		}
		if a == "/*" {
			ignoreLines = true
			continue
		}

		b := reData.FindStringSubmatch(a)
		if b == nil {
			l.Printf("[WARNING] Line didn't match: \"%s\"\n", a)
			continue
		}
		if len(b) < 3 {
			l.Fatal("[ERROR] Regexp result is missing a value")
		}
		modVar, modParts := b[1], b[3]
		modRaw := reClean.Replace(b[2])

		partMatchSlice := rePart.FindAllStringSubmatch(modParts, -1)
		if partMatchSlice == nil {
			l.Printf("[WARNING] Failed to parse parts '%s'\n", modParts)
		}
		partSlice := make([]string, len(partMatchSlice))
		for i := range partSlice {
			partSlice[i] = partMatchSlice[i][1]
		}

		if modKey, ok := moduleMap[modVar]; ok {
			if index, ok := order[modKey]; ok {
				modules = append(modules, json.IconicModule{
					Key:   modKey,
					Raw:   modRaw,
					Parts: partSlice,
					Order: index,
				})
			} else {
				l.Printf("[WARNING] Failed to find module '%s' in modules.txt\n", modKey)
			}
		} else {
			l.Printf("[WARNING] Failed to find variable '_%s' in iconicScript.cs\n", modVar)
		}
	}
	if err := s.Err(); err != nil {
		l.Fatal("[ERROR] Failed to parse iconic script: ", err)
	}
	return modules
}
