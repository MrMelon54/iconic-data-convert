package manual

var manualConvert = map[string]string{
	"Needy Quiz":            "NeedyVentV2",
	"Needy Rotary Phone":    "NeedyKnobV2",
	"Needy Button Masher":   "buttonMasherNeedy",
	"Needy Beer Refill Mod": "NeedyBeer",
	"Needy Shape Memory":    "needyShapeMemory",
	"Needy Wingdings":       "needyWingdings",
	"Needy Pong":            "NeedyPong",
	"Needy Crafting Table":  "needycrafting",
	"3x3 Grid":              "3x3Grid",
}

// ManuallyConvertDisplayNameToID converts modules missing the DisplayName field
// on the repo raw json but with weird name mismatches.
func ManuallyConvertDisplayNameToID(displayName string) string {
	return manualConvert[displayName]
}
