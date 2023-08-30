package json

type IconicData struct {
	Underscore string         `json:"_"`
	Modules    []IconicModule `json:"modules"`
}

type IconicModule struct {
	Key   string   `json:"key"`
	Icon  string   `json:"icon"`
	Raw   string   `json:"raw"`
	Parts []string `json:"parts"`
	Order int      `json:"-"`
}
