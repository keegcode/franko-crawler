FROM golang:1.21.6

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd/cli

CMD ["sh", "-c", "/app/cli --api ${API_KEY} --id ${CHANNEL_ID}"]