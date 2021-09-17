
all:


### go stuff
go_build:
	go build main.go

go_run:
	go run main.go

go_test:
	curl -s http://localhost:8080


### Docker stuff
docker_build:
	docker build -t log-tester .

docker_run:
	$(eval MY_C=$(shell docker run -d log-tester|cut -c 1-12))
	$(eval MY_IP=$(shell docker inspect -f "{{ .NetworkSettings.IPAddress }}" $(MY_C)))
	@echo "docker container $(MY_C) $(MY_IP)"
	curl http://$(MY_IP):8080/

docker_clean:
	$(eval MY_C=$(shell docker ps |grep log-tester| awk '{print $$1}'))
	docker rm -f $(MY_C)

docker_test: docker_run docker_clean

