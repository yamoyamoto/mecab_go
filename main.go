package main

import (
	"fmt"
	"github.com/bluele/mecab-golang"
	"github.com/yamoyamoto/mecab_go/features"
)

func parseToNode(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice("かわぐちともやは大阪出身です")
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := features.NewFeatures(node.Feature())
		fmt.Println(node.Surface(), features)

		if node.Next() != nil {
			break
		}
	}
}

func main() {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	parseToNode(m)
}
