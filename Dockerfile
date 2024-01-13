FROM golang:1.21-alpine as golang

RUN apk add -U tzdata
RUN apk --update add ca-certificates

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY layout.html  ./
COPY techniques.txt ./

RUN go mod download
RUN go mod verify


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /aiki_random main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /healthcheck healthcheck.go 

FROM scratch

COPY --from=golang /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=golang /etc/passwd /etc/passwd
COPY --from=golang /etc/group /etc/group

COPY --from=golang /aiki_random .
COPY --from=golang /healthcheck .
COPY --from=golang /app/layout.html  .
COPY --from=golang /app/techniques.txt .


HEALTHCHECK --interval=1s --timeout=1s --start-period=2s --retries=3 CMD [ "/healthcheck" ]

EXPOSE 8000

CMD [ "/aiki_random" ]
