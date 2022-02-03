module github.com/yamoyamoto/mecab_go

go 1.16

replace github.com/yamoyamoto/mecab_go/models => ./models

replace github.com/yamoyamoto/mecab_go/handlers => ./handlers

require (
	github.com/bluele/mecab-golang v0.0.0-20180831023624-c8cfe04e87f9 // indirect
	github.com/yamoyamoto/mecab_go/handlers v0.0.0-00010101000000-000000000000
)
