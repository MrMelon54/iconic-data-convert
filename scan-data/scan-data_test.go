package scan_data

import (
	"bufio"
	"bytes"
	_ "embed"
	"github.com/MrMelon54/iconic-data-convert/json"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

const (
	dataTestStringRaw = "   00000000 00000000 000    00   000000000000000000000000  0000  000000000000000000000000 0000000000000000000000000000000 00000000000000000000000000000000 0000 000000000000000000000000000 00  0000000000000000000000000000    0000000000000000000000000000000 000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000 0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 000000000000000000000000000000  000000000000000000000000000000    00000000 00000000 00000000   "
	dataTestString    = "public static string[] _Wires = {\"" + dataTestStringRaw + "\", \"Somewhere on the module\"};"
)

var (
	//go:embed validIconicData.txt
	validIconicData []byte
	//go:embed invalid/invalidLine.txt
	invalidLine []byte
	//go:embed invalid/invalidParts.txt
	invalidParts []byte
)

func makeLog(buf *bytes.Buffer) *log.Logger {
	// flag 0 must be used so the date isn't outputted in the logging lines
	return log.New(buf, "", 0)
}

func TestDataRegexp(t *testing.T) {
	assert.Equal(t, []string{dataTestString, "Wires", dataTestStringRaw, "\"Somewhere on the module\""}, reData.FindStringSubmatch(dataTestString))
	assert.Nil(t, reData.FindStringSubmatch(" invalid string "))
}

func TestScanIconicData(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := map[string]string{"Wires": "Wires", "WhosOnFirst": "Who's on First"}
	order := map[string]int{"Wires": 0, "Who's on First": 1}
	modules := ScanIconicData(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(validIconicData)), moduleMap, order)
	assert.Equal(t, []json.IconicModule{
		{Key: "Wires", Raw: "00", Parts: []string{"First wire", "Second wire"}, Order: 0},
		{Key: "Who's on First", Raw: "11", Parts: []string{"Display", "Top-left button"}, Order: 1},
	}, modules)
	assert.Equal(t, "", fakeLog.String())
}

func TestScanIconicData_InvalidLine(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := map[string]string{"Wires": "Wires", "WhosOnFirst": "Who's on First"}
	order := map[string]int{"Wires": 0, "Who's on First": 1}
	modules := ScanIconicData(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(invalidLine)), moduleMap, order)
	assert.Equal(t, []json.IconicModule{
		{Key: "Who's on First", Raw: "11", Parts: []string{"Display", "Top-left button"}, Order: 1},
	}, modules)
	assert.Equal(t, "[WARNING] Line didn't match: \"public static _Wires = {\"00\", \"First wire\", \"Second wire\" };\"\n", fakeLog.String())
}

func TestScanIconicData_InvalidOrder(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := map[string]string{"Wires": "Wires", "WhosOnFirst": "Who's on First"}
	order := map[string]int{"Wires": 0}
	modules := ScanIconicData(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(validIconicData)), moduleMap, order)
	assert.Equal(t, []json.IconicModule{
		{Key: "Wires", Raw: "00", Parts: []string{"First wire", "Second wire"}, Order: 0},
	}, modules)
	assert.Equal(t, "[WARNING] Failed to find module 'Who's on First' in modules.txt\n", fakeLog.String())
}

func TestScanIconicData_InvalidVariable(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := map[string]string{"Wires": "Wires"}
	order := map[string]int{"Wires": 0, "Who's on First": 1}
	modules := ScanIconicData(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(validIconicData)), moduleMap, order)
	assert.Equal(t, []json.IconicModule{
		{Key: "Wires", Raw: "00", Parts: []string{"First wire", "Second wire"}, Order: 0},
	}, modules)
	assert.Equal(t, "[WARNING] Failed to find variable '_WhosOnFirst' in iconicScript.cs\n", fakeLog.String())
}

func TestScanIconicData_InvalidParts(t *testing.T) {
	fakeLog := new(bytes.Buffer)
	moduleMap := map[string]string{"Wires": "Wires", "WhosOnFirst": "Who's on First"}
	order := map[string]int{"Wires": 0, "Who's on First": 1}
	modules := ScanIconicData(makeLog(fakeLog), bufio.NewScanner(bytes.NewBuffer(invalidParts)), moduleMap, order)
	assert.Equal(t, []json.IconicModule{
		{Key: "Wires", Raw: "00", Parts: []string{"First wire", "Second wire"}, Order: 0},
		{Key: "Who's on First", Raw: "11", Parts: []string{}, Order: 1},
	}, modules)
	assert.Equal(t, "[WARNING] Failed to parse parts '\"Dis '\n", fakeLog.String())
}
