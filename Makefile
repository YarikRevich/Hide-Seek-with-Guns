.PHONY: standarts packages install exit

.SILENT: exit

NAME := $(shell uname -s)

standarts:
	clear > $(shell tty)
ifeq ($(NAME), Darwin) 
	@echo "\n --- Installs golang via brew --- \n"
	@(brew install golang 1> /dev/null)
endif
	
ifeq ($(NAME), Linux) 
	@echo "\n --- Installs golang via apt --- \n"
	@(sudo apt install golang > /dev/null)
endif

ifeq ($(NAME), Windows)
	@printf "\n --- if your os is Windows you are left without game ㋡ hahahhah --- \n"
	@$(MAKE) exit
endif
	clear > $(shell tty)
	@echo "\n --- Creates main folder for project and pulls a project in it --- \n"
	@(cd ${GOPATH}/src && mkdir -p HideSeek)
	@cd ${GOPATH}/src/HideSeek && git init > /dev/null && git pull --quiet https://github.com/YarikRevich/Hide-Seek-with-Guns super-alpha > /dev/null
	$(MAKE) packages

packages:
	clear > $(shell tty)
	@echo "\n --- Installs all the important packages for project ---\n"
	@(go get github.com/galsondor/go-ascii)
	@(go get github.com/go-ping/ping)
	@(go get github.com/faiface/beep)

install:
	clear > $(shell tty)
	@cd ${GOPATH}/src/HideSeek && echo "\n --- Builds project --- \n" && go build main.go > /dev/null

exit: 
	$(error 1)
