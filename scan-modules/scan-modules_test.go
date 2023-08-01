package scan_modules

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"testing"
)

func TestScanIconicModules(t *testing.T) {
	order := ScanIconicModules(log.New(io.Discard, "", 0), bufio.NewScanner(bytes.NewBufferString("Wires\nThe Button\n")))
	assert.Equal(t, map[string]int{"Wires": 0, "The Button": 1}, order)
}
