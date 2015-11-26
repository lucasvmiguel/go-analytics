#http://learnxinyminutes.com/docs/make/

env ?= develop

main: debug
	@echo 'Task:' $@

dependencies:
	@echo 'Task:' $@
	@go get github.com/gin-gonic/gin
	@go get github.com/spf13/viper
	@go get github.com/gorilla/websocket

build: dependencies
	@echo 'Task:' $@ '('${env}')'
	@go build main.go

debug: dependencies build up_containers
	@echo 'Task:' $@
	@./main develop

deploy: dependencies up_containers
	@echo 'Task:' $@

up_containers:
	#run docker
