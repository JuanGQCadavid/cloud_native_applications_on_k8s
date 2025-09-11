build:
	buildah build -f Dockerfile -t price-fetcher-service:latest .

push:
	buildah push --tls-verify=false localhost/price-fetcher-service:latest docker://10.4.0.6:5000/price-fetcher-service:latest

run:
	sudo nerdctl run -v /mnt/price-data:/tmp/price-data/  -d --name price-fetcher-service 10.4.0.28:5000/price-fetcher-service:latest