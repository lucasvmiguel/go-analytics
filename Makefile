#http://learnxinyminutes.com/docs/make/

env ?= develop

main: debug
	@echo 'Task:' $@

#não esqueça de setar o gvm
dependencies:
	@echo 'Task:' $@
	@go get github.com/gin-gonic/gin
	@go get github.com/spf13/viper
	@go get github.com/gorilla/websocket
	@go get github.com/influxdb/influxdb/client/v2
	@go get github.com/Sirupsen/logrus
	@go get gopkg.in/redis.v3

build: dependencies
	@echo 'Task:' $@ '('${env}')'
	@go build main.go

debug: dependencies build
	@echo 'Task:' $@
	@./main develop

deploy: dependencies up_containers
	@echo 'Task:' $@

containers_pull:
	@sudo docker pull tutum/influxdb
	@sudo docker pull bitnami/redis
	@sudo docker pull grafana/grafana
	@sudo docker pull tutum/grafana
	@sudo docker run --name docker_data --volume /datafolder ubuntu true

containers_run:
	@sudo docker run -d -p 8083:8083 -p 8086:8086 -e PRE_CREATE_DB="go-analytics" --volumes-from docker_data tutum/influxdb:latest
	#http://docs.grafana.org/installation/configuration/
	@sudo docker run -d -p 3000:3000 --volumes-from docker_data -e GF_SECURITY_ADMIN_PASSWORD=admin -e GF_SECURITY_ADMIN_USER=admin grafana/grafana
	#@sudo docker run -d bitnami/redis
