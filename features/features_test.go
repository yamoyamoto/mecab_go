package features

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeatures(t *testing.T) {
	str := "品詞,品詞細分類1,品詞細分類2,品詞細分類3,活用型,活用形,原型,読み,発音"
	features := NewFeatures(str)

	assert.Equal(t, "品詞", features.品詞)
}
