FROM golang:latest

RUN apt update && \
    apt -y upgrade && \
    apt install -y build-essential && \
    apt install -y software-properties-common && \
    apt install -y curl git man unzip vim wget sudo && \
    apt install -y mecab libmecab-dev mecab-ipadic-utf8

WORKDIR /go/src/github.com/yamoyamoto/ir_api
# RUN git clone https://github.com/neologd/mecab-ipadic-neologd.git && \
#     cd mecab-ipadic-neologd && \
#     sudo bin/install-mecab-ipadic-neologd


# /usr/lib/x86_64-linux-gnu/mecab/dic/mecab-ipadic-neologd
ENV CGO_LDFLAGS="-L/usr/lib/x86_64-linux-gnu -lmecab -lstdc++"
ENV CGO_CFLAGS="-I/usr/include"

ENV GOOS=linux
ENV GOARCH=amd64
COPY . .
RUN go build main.go

CMD ./main