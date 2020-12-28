# <center>Instalation</center>

!["preview"](preview.png)

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
```

Then compile project due to the command beneath
```
go build main.go
```
So, you can run this game!
```
./main
```
