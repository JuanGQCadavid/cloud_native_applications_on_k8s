build:
	cd services/price_fetching/cmd/cron \
	&& CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o price_fetching
	mv services/price_fetching/cmd/cron/price_fetching builds/