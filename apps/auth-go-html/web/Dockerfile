FROM golang:1.21.0-alpine AS build
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=8080
RUN apk add ca-certificates

WORKDIR /src
COPY . .
WORKDIR /src/cmd
RUN go build -o /out/app .

FROM scratch AS bin
COPY --from=build /out/app /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./tmpl /tmpl

ENTRYPOINT [ "/app" ]

EXPOSE ${PORT}