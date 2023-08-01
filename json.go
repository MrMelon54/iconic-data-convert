package main

type IconicData struct {
	Underscore string         `json:"_"`
	Modules    []IconicModule `json:"modules"`
}

type IconicModule struct {
	Key   string   `json:"key"`
	Raw   string   `json:"raw"`
	Parts []string `json:"parts"`
}
