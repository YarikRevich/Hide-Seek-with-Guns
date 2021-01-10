.PHONY: install

NAME := $(shell uname -s)

standarts:
ifeq ($(NAME), Darwin) 
	@echo "\n --- Installs golang via brew --- \n"
	@(brew install golang)
endif
	
ifeq ($(NAME), Linux) 
	@echo "\n --- Installs golang via apt --- \n"
	@(sudo apt install golang)
endif

	@echo "\n --- if your os is Windows you are left without game ㋡ hahahhah --- \n"

	@echo "\n --- Creates main folder for project and pulls a project in it --- \n"
	@(cd ${GOPATH}/src && mkdir HideSeek
	git init && git pull https://github.com/YarikRevich/Hide-Seek-with-Guns master)

packages: standarts
	@echo "\n --- Installs all the important packages for project ---\n"
	@(go get github.com/galsondor/go-ascii
	go get github.com/go-ping/ping
	go get github.com/faiface/beep)

install: packages
	@echo "\n --- Builds project --- \n"
	@(go build main.go)
