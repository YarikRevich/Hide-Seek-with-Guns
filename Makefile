.PHONY: all exit

.ONESHELL:

.SILENT: all exit

CURRDIRBASE := $(shell basename "${PWD}")

CURRDIRFULL = $(CURDIR)

NAME := $(shell uname -s)

all:
	clear > $(shell tty)
ifeq ($(NAME), Darwin) 
	printf "\n --- Installs golang via brew --- \n"
ifeq ($(shell ${USER}), root)
	printf "\n --- Switch user to non root --- \n"
	$(MAKE) exit
endif
	brew install golang > /dev/null
endif
	
ifeq ($(NAME), Linux) 
	printf "\n --- Installs golang via apt --- \n"
	sudo apt install golang > /dev/null
endif

ifeq ($(NAME), Windows)
	printf "\n --- if your os is Windows you are left without a game ㋡ hahahhah --- \n"
	$(MAKE) exit
endif
	clear > $(shell tty);\
	printf "\n --- Creates the environment for the game --- \n";\
	mkdir -p $$GOPATH/src;\
	mv $(CURRDIRFULL) $$GOPATH/src;\
	clear > $(shell tty);\
    printf "\n --- Installs all the important packages for project ---\n";\
	go get github.com/galsondor/go-ascii;\
	go get github.com/go-ping/ping;\
	go get github.com/faiface/beep;\
	clear > $(shell tty);\
	printf "\n --- Builds project --- \n";\
	go build main.go > /dev/null

exit: 
	$(error 1)
