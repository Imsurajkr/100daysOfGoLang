History of Go
   1.Go was an interal project of Google and was started with an experiment.
   2.Go is inspired by many other programming languages like c,pascal,Alef and oberon.
   3.Robert Griesemer,Ken Thomson  and Rob pike are the developers of Go.

Aim of Go 
  To build reliable,robust and efficient software.

GoDoc utility
    You can read existing go docs function and packages without any internet connection.
    godoc utility can be executed normally on command line arguments that displays its output on terminal.
    this can be done using two ways:----------------------------------------------------------------
    first way:---
    suppose you want to see documentation of fmt package offline,then use following command
    go doc fmt

    second way:----------------------------------------------------------------
    godoc -http=:8000 
    ##8000 is the port number where the docs will open in offline mode.
    you can choose any port number of your own choice but not in between 0-1023 because they are restricted to be used by root user only.
    ##after that,point your browser to http://localhost:8001/pkg/  url to browse the documentation.
    ##there may cause an error that"godoc is not recognised as an internal and external command
    then you should run below command in your terminal
    go get -v  golang.org/x/tools/cmd/godoc
    which will install godoc and after that run godoc in your terminal with a port number.
    
