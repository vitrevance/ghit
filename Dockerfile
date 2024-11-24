FROM golang:1.22-alpine3.19 as build

WORKDIR /ghit

COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./

ENV CGO_ENABLED=0
RUN go build -o ./ghit ./main.go

FROM scratch

WORKDIR /

COPY --from=build /ghit/ghit /ghit

ENTRYPOINT [ "/ghit" ]