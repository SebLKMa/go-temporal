staging=../staging/surl
bin=gosurlws
LD_VERSION=V1
LD_BUILDTIME=`date +%FT%T%z`
dock_tag=$(DOCKTAG)
dock_file=../Dockerfile-gosurlws
dock_file_arm64=../Dockerfile-gosurlws-arm64
dock_name=gosurlws
dock_image_arm64=gosurlws-arm64
dock_hub=local#pdntu.azurecr.io
dock_env=.env
dock_internal_port=8181
dock_external_port=8181
dock_local_port=8282

# defines newline for printing messages on new line
define newline


endef

check_args:
ifeq ($(dock_tag),)
dock_tag := 0.1
endif
# Using default above
#ifndef DOCKTAG
#	$(error ${newline}Usage:${newline}      make -f Makefile <command> DOCKTAG=<version>${newline}Please provide DOCKTAG argument)
#else
#	@echo "docker image tag version ${dock_tag} used in make"
#endif

stage-amd:
	@echo "staging local as amd64: " ${staging}
	mkdir -p ${staging}
#	@echo "Building for local machine..."
#	go build -ldflags="-X 'main.Version=${LD_VERSION}' -X 'main.Build=${LD_BUILDTIME}'" -o ${staging}/${bin} main.go
	@echo "Building for Docker Alpine"
# How do I build a static Go binary for the Docker Alpine image?
# https://stackoverflow.com/questions/48366173/how-do-i-build-a-static-go-binary-for-the-docker-alpine-image
# https://github.com/golang/go/issues/9344#issuecomment-69944514
# Ensure cross-compile to alpine linux intel
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo \
	-ldflags="-X 'main.Version=${LD_VERSION}' -X 'main.Build=${LD_BUILDTIME}'" \
	-o ${staging}/${bin} main.go

stage-local:
	@echo "staging local arch"
	mkdir -p ${staging}
	@echo "Building for Local Docker Alpine"
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-X 'main.Version=${LD_VERSION}' -X 'main.Build=${LD_BUILDTIME}'" \
	-o ${staging}/${bin} main.go

# DOCKER
#
# Removes image from local docker
docker-clean-arm64: check_args
	docker rmi ${dock_hub}/${dock_image_arm64}:${dock_tag}

# Builds image for local docker
docker-build-arm64: check_args stage-local
	docker login
	docker build --no-cache -t ${dock_hub}/${dock_image_arm64}:${dock_tag} -f ${dock_file_arm64} ../

# Runs local docker container
# etc/postgresql/14/main/postgresql.conf has to allow specific ip or all, e.g. listen_addresses = '*'
docker-runlocal-arm64: check_args
	docker run -p ${dock_local_port}:${dock_internal_port} --env DBHOST=172.17.0.1 --env DBPORT=5432 --env-file ${dock_env} -it --rm -d --name ${dock_name} ${dock_hub}/${dock_image_arm64}:${dock_tag}

docker-stop:
	sudo docker stop ${dock_name}

docker-attach:
	sudo docker exec -i -t ${dock_name} /bin/sh

docker-rm:
	sudo docker rm ${dock_name}

docker-ls:
	docker images | grep ${dock_name}
	docker ps -a | grep ${dock_name}

