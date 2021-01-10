# <center>Instalation</center>

!["preview"](preview.png)


**_You can install game using commands below_**

```
make    #it works if you are in a dir with Makefile
make install
```

---

**_If you are a hard core individual you may try to install it by yourself using instruction below._**



To start playing 'Hide&Seek' you have to do some things
Firstly you have to install golang

- linux 
    ```
    sudo apt install golang
    ```
- macOS
    ```
    brew install golang
    ```

Then you should create the dir in the GOPATH and move yourself to it.
```
cd $GOPATH/src && mkdir HideSeek
```

If you don't have 'src' dir yet create it and then repeat previously said instruction.

Afterwards you should pull this repo to made dir
```
git init && git pull **this repo** master
```

Download important dependences 
```
go get github.com/galsondor/go-ascii
go get github.com/go-ping/ping
go get github.com/faiface/beep
```

Then compile project due to the command infra
```
go build main.go
```
So, you can run this game!
```
./main
```
