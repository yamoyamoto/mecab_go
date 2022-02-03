package main

import (
	"fmt"
	"github.com/bluele/mecab-golang"
	"strings"
)

func parseToNode(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice("すもももももももものうち")
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			fmt.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
		}
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
