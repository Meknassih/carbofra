FROM golang:latest
ENV PORT=${PORT}
ENV MONGO_URI=${MONGO_URI}
WORKDIR /go/carbofra
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["carbofra"]