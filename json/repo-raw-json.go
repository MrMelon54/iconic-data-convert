package json

import (
	"github.com/MrMelon54/iconic-data-convert/manual"
	"strings"
)

type KtaneRawJson struct {
	KtaneModules []KtaneModule
}

func (k *KtaneRawJson) ConvertDisplayNameToID(displayName string) string {
	displayName = strings.ReplaceAll(displayName, "\u2019", "'")
	for _, i := range k.KtaneModules {
		if i.DisplayName != "" && strings.EqualFold(i.DisplayName, displayName) {
			return i.ModuleID
		}
		if i.Name != "" && strings.EqualFold(i.Name, displayName) {
			return i.ModuleID
		}
	}
	return manual.ManuallyConvertDisplayNameToID(displayName)
}

func (k *KtaneRawJson) ConvertIdToIconName(id string) string {
	for _, i := range k.KtaneModules {
		if i.ModuleID == id {
			if i.FileName != "" {
				return i.FileName
			}
			if i.Name != "" {
				return i.Name
			}
			return i.ModuleID
		}
	}
	return ""
}

type KtaneModule struct {
	DisplayName string
	ModuleID    string
	Name        string
	SteamID     string
	Symbol      string
	FileName    string
}
