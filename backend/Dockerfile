FROM golang:alpine AS build

RUN apk update && apk upgrade && apk add --no-cache gcc libc-dev

WORKDIR /app

COPY go.mod go.sum ./
COPY api/ api/
COPY cmd/ cmd/
COPY internal/ internal/

RUN CGO_ENABLED=1 go build -a -o foodiee-backend cmd/foodiee_backend/main.go



FROM alpine:latest

ENV DB_FILE_PATH=/app/foodiee.db

RUN apk update && apk upgrade && apk add --no-cache sqlite

USER 1000
WORKDIR /app

COPY --from=build /app/foodiee-backend .
COPY entrypoint.sh schema.sql ./

CMD [ "./entrypoint.sh" ]
