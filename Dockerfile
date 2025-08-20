FROM --platform=$BUILDPLATFORM golang:1.23.2-alpine AS builder
LABEL authors="gucooing"
RUN apk add --no-cache bash protoc protobuf-dev curl
ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

#RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
WORKDIR /app
COPY . .

RUN --mount=type=secret,id=sha,env=SHA \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build \
    -ldflags="-s -w -X github.com/gucooing/BaPs/protocol/mx.Docker=1 -X github.com/gucooing/BaPs/pkg.Commit=$SHA" \
    -o /usr/ba/BaPs \
    ./cmd/BaPs/BaPs.go

COPY ./data/ /usr/ba/data/
COPY ./resources/ /usr/ba/resources/

RUN --mount=type=secret,id=sha,env=SHA \
    GOOS=$TARGETOS GOARCH=$TARGETARCH go build "-tags=debug" \
    -ldflags="-s -w -X github.com/gucooing/BaPs/protocol/mx.Docker=1 -X github.com/gucooing/BaPs/pkg.Commit=$SHA" \
    -o /usr/ba/GenExcelBin \
    ./cmd/GenExcelBin/main.go

RUN cd /usr/ba/ && chmod 777 ./GenExcelBin && ./GenExcelBin

RUN cd /usr/ba/ && ls

RUN cd /usr/ba/data && ls

# 最终镜像
FROM --platform=${TARGETPLATFORM} alpine:latest
RUN apk add --no-cache bash tzdata
WORKDIR /usr/ba
COPY --from=builder /usr/ba/BaPs .
#COPY --from=builder /usr/ba/data/ ./data/
RUN chmod +x BaPs
EXPOSE 5000/tcp
ENTRYPOINT ["./BaPs"]
CMD ["-c", "./config/config.json"]