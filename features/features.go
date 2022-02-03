package features

import (
	"strings"
)

const SEPALATOR = ","

// 後で英語に
type Features struct {
	品詞     string
	品詞細分類1 string
	品詞細分類2 string
	品詞細分類3 string
	活用型    string
	活用形    string
	原形     string
	読み     string
	発音     string
}

func NewFeatures(s string) *Features {
	strSlice := strings.Split(s, SEPALATOR)

	features := &Features{
		品詞:     strSlice[0],
		品詞細分類1: strSlice[1],
		品詞細分類2: strSlice[2],
		品詞細分類3: strSlice[3],
		活用型:    strSlice[4],
		活用形:    strSlice[5],
		原形:     strSlice[6],
		読み:     strSlice[7],
		発音:     strSlice[8],
	}

	return features
}
