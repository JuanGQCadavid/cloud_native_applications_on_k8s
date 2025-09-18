build:
	buildah build -f Dockerfile -t server:latest .

push:
	buildah push --format docker server:latest docker-archive:my.server.tar
	