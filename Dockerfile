# builder stage
FROM golang:alpine as builder
WORKDIR /cryptobot
ENV FOLDER_PATH service-7_diff-binance-cryptomkt
ADD . .
RUN go mod download
RUN go mod verify
RUN go build -o bin/$FOLDER_PATH-amd64 cmd/$FOLDER_PATH/main.go

# final stage
FROM alpin
WORKDIR /uluru-api
ENV SVC uluru-api

COPY bin/${SVC} /usr/bin/${SVC}
COPY misc/seeds/goals.json goals.json
COPY misc/seeds/institutions.json institutions.json
COPY misc/seeds/subscriptions.json subscriptions.json
COPY --from=builder /cryptobot/bin/$SVC-amd64 /uluru-api

ENTRYPOINT ./$SVC-amd64