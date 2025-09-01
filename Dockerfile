FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && \
  apt-get install -y ca-certificates golang-go


WORKDIR /app

COPY . .

RUN go build -o notely . && mv notely /usr/bin/notely

CMD ["notely"]
