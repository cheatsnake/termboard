FROM golang:alpine AS build
WORKDIR /app
ADD . /app
RUN cd /app && go build ./cmd/main.go

FROM scratch
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]