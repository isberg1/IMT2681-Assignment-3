FROM golang:1.11.1-stretch as build

LABEL maintainer=Celebrian

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o API .



FROM scratch

WORKDIR /root/

COPY --from=build /app/API .

COPY README.md /root/

CMD ["./API"]

EXPOSE 8080
