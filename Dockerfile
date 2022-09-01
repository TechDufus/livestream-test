FROM golang:latest as BUILDER

WORKDIR /usr/flux/app/
COPY . .

# RUN go mod init catfact
RUN go build catfact.go

FROM debian:latest as RUNNER

WORKDIR /usr/flux/app/
COPY --from=BUILDER /usr/flux/app/catfact /usr/flux/app

ENTRYPOINT [ "./catfact" ]

