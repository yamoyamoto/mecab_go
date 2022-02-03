package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bluele/mecab-golang"
	"net/http"
	"sort"
)

type MorphologicalAnalysisResponse struct {
	Morphemes []Morpheme `json:"morphemes"`
}

type Morpheme struct {
	Surface string `json:"surface"`
	Count   int    `json:"count"`
}

func MorphologicalAnalysis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Print("Start alalysis...")

	s := r.FormValue("s")
	fmt.Println(s)
	morphemes := Parse(s)

	morphologicalAnalysisResponse := MorphologicalAnalysisResponse{
		Morphemes: morphemes,
	}

	output, _ := json.Marshal(morphologicalAnalysisResponse)
	w.Write(output)
}

func Parse(s string) []Morpheme {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(s)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	var strSlice []string
	for {
		strSlice = append(strSlice, node.Surface())
		if node.Next() != nil {
			break
		}
	}
	morphemes := groupBy(strSlice)

	return morphemes
}

func groupBy(s []string) []Morpheme {
	count_map := make(map[string]int)
	for _, v := range s {
		if count_map[v] == 0 {
			count_map[v] = 1
		} else {
			count_map[v]++
		}
	}

	var morphemes []Morpheme
	for k, v := range count_map {
		morphemes = append(morphemes, Morpheme{
			Surface: k,
			Count:   v,
		})
	}

	return sortMorpheme(morphemes)
}

func sortMorpheme(m []Morpheme) []Morpheme {
	sort.Slice(m, func(i, j int) bool {
		return m[i].Count > m[j].Count
	})
	return m
}
