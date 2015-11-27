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
	#@go get github.com/Sirupsen/logrus
	#@go get gopkg.in/mgo.v2
	#@go get gopkg.in/redis.v3

build: dependencies
	@echo 'Task:' $@ '('${env}')'
	@go build main.go

debug: dependencies build
	@echo 'Task:' $@
	@./main develop

deploy: dependencies up_containers
	@echo 'Task:' $@

up_containers:
	@sudo docker pull mongo
	@sudo docker pull frodenas/mongodb
	@sudo docker run -d -p 8083:8083 -p 8086:8086 -e PRE_CREATE_DB="go-analytics" tutum/influxdb:latest
	@sudo docker run -d mongo
