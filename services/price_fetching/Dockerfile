FROM registry.hpc.ut.ee/mirror/library/golang:1.25.1-alpine3.22 as builder

ENV CGO_ENABLED="0"

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY core core
COPY repository repository

RUN go build -o main ./cmd/cron/main.go


FROM registry.hpc.ut.ee/mirror/library/alpine as prod
RUN mkdir -p /tmp/price-data/ 

COPY --from=builder /app/main /bin/
COPY cronjobs /etc/crontabs/root

RUN chown root:root /etc/crontabs/root

CMD ["crond", "-f", "-d", "8"]
