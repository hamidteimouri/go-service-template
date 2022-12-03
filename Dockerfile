FROM golang:1.16.4-alpine as build

RUN apk add --no-cache git

WORKDIR /app/ht

COPY /src/go.mod .
COPY /src/go.sum .

RUN ls

RUN go mod download

COPY . .

WORKDIR /app/ht/src/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./out/ht


FROM alpine AS final
RUN apk add --no-cache tzdata

# Copy the compiled file to final light weight image
COPY --from=build /app/ht/src/out/ht /ht

# gRPC port
EXPOSE 50051

## Run the binary
ENTRYPOINT ["/ht"]
