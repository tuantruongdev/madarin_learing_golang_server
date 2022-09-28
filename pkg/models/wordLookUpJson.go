package models

type EntryJson struct {
	Traditional string   `json:"1"`
	Pinyin      string   `json:"2"`
	Defs        []string `json:"4"`
}

type CharacterJson struct {
	Simple  string      `json:"0"`
	Entries []EntryJson `json:"3"`
	Rank    int         `json:"5"`
	Hsk     int         `json:"6"`
}
