FROM golang:1.18 AS BUILD

RUN go env -w  GOPROXY="https://goproxy.io,direct"

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

RUN cd /app/app/background \
    && CGO_ENABLED=0  \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o job.exe

FROM alpine AS PRODUCT

WORKDIR /app

COPY --from=BUILD /app/app/background/job.exe .

CMD ["/app/job.exe"]

FROM scratch AS EXPORT

COPY --from=BUILD /app/app/background/job.exe /