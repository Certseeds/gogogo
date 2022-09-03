FROM golang:1.18 AS build_prepare

RUN go env -w  GOPROXY="https://goproxy.io,direct"

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . .

FROM build_prepare AS TEST

WORKDIR /app/pkg/background
ENTRYPOINT ["go", "test","-v", "./..."]

FROM build_prepare AS BUILD

RUN cd /app/app/background \
    && CGO_ENABLED=0  \
        GOOS=linux \
        GOARCH=amd64 \
        go build \
        -o job.exe

FROM alpine AS PRODUCT

WORKDIR /app

COPY --from=BUILD /app/app/background/job.exe .

CMD ["/app/job.exe"]

#FROM scratch AS EXPORT
#
#COPY --from=BUILD /app/app/background/job.exe /