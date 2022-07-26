FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN go build -o /sberlabAlbumsApp

ENV LISTEN_ON=$LISTEN_ON

EXPOSE 7331

CMD [ "/sberlabAlbumsApp" ]