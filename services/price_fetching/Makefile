VERSION := $(if $(VERSION),$(VERSION),)

build:
	buildah build -f Dockerfile -t price-fetcher-service:${VERSION} .

push:
	buildah push --tls-verify=false localhost/price-fetcher-service:${VERSION} docker://10.4.0.6:5000/price-fetcher-service:${VERSION}

run:
	sudo nerdctl run  -v /mnt/price-data:/tmp/price-data/  -d --name price-fetcher-service --insecure-registry 10.4.0.6:5000/price-fetcher-service:${VERSION}

kill:
	sudo nerdctl kill price-fetcher-service && sudo nerdctl rm price-fetcher-service

pull:
	git pull

bitch: pull kill build push run