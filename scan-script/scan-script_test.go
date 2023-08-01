package scan_script

import (
	"bufio"
	"bytes"
	_ "embed"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	scriptTestString = "{ \"Wires\", iconicData._Wires },"

	//go:embed validIconicScript.txt
	validIconicScript []byte
	//go:embed invalid/invalidLine.txt
	invalidLine []byte
)

func makeLog(buf *bytes.Buffer) *log.Logger {
	// flag 0 must be used so the date isn't outputted in the logging lines
	return log.New(buf, "", 0)
}

func TestScriptRegexp(t *testing.T) {
	assert.Equal(t, []string{scriptTestString, "Wires", "Wires"}, reScript.FindStringSubmatch(scriptTestString))
	assert.Nil(t, reScript.FindStringSubmatch(" invalid string "))
}

func TestScanIconicScript(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := ScanIconicScript(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(validIconicScript)))
	assert.Equal(t, map[string]string{
		"Wires":              "Wires",
		"TheButton":          "The Button",
		"Keypad":             "Keypad",
		"SimonSays":          "Simon Says",
		"WhosOnFirst":        "Who's on First",
		"Memory":             "Memory",
		"MorseCode":          "Morse Code",
		"ComplicatedWires":   "Complicated Wires",
		"WireSequence":       "Wire Sequence",
		"Maze":               "Maze",
		"Password":           "Password",
		"VentingGas":         "Needy Vent Gas",
		"CapacitorDischarge": "Needy Capacitor",
		"Knob":               "Needy Knob",
	}, moduleMap)
}

func TestScanIconicScript_InvalidLine(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := ScanIconicScript(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(invalidLine)))
	assert.Equal(t, map[string]string{
		"Wires":              "Wires",
		"TheButton":          "The Button",
		"Keypad":             "Keypad",
		"SimonSays":          "Simon Says",
		"WhosOnFirst":        "Who's on First",
		"Memory":             "Memory",
		"ComplicatedWires":   "Complicated Wires",
		"WireSequence":       "Wire Sequence",
		"Maze":               "Maze",
		"Password":           "Password",
		"VentingGas":         "Needy Vent Gas",
		"CapacitorDischarge": "Needy Capacitor",
		"Knob":               "Needy Knob",
	}, moduleMap)
	assert.Equal(t, `[WARNING] Line didn't match: "{ " invalid line ", Code },"
[WARNING] Line didn't match: "};"
`, fakeLog.String())
}
