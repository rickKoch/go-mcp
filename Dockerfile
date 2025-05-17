FROM golang:1.24.3-alpine AS build
WORKDIR /code
COPY . .
RUN  CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/simplemcp /code/cmd/main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /server
COPY --from=build /bin/simplemcp .
CMD ["./simplemcp"]
