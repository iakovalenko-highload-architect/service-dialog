FROM golang:1.21 as service-dialog

WORKDIR /project

COPY go.mod .
RUN go mod download

COPY . /project
RUN go build -o /bin/service-dialog -v ./cmd/service

RUN rm -rf /project

CMD ["/bin/service-dialog"]
