FROM golang:1.10 AS builder

RUN curl -fsSL -o /usr/local/bin/dep \
    https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \ 
    chmod +x /usr/local/bin/dep

RUN mkdir -p $GOPATH/src/github.com/TerrenceHo/ABFeature
WORKDIR $GOPATH/src/github.com/TerrenceHo/ABFeature

ADD Gopkg.lock Gopkg.toml ./
RUN dep ensure -vendor-only

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abfeature \
    ./cmd/ABFeature/main.go

# Second stage build
FROM scratch

WORKDIR /root

COPY --from=builder /go/src/github.com/TerrenceHo/ABFeature/abfeature .

EXPOSE 13317

ENTRYPOINT ["./abfeature"]
