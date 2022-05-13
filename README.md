# Golang 101

This is the video tutorial that I'll follow [https://www.youtube.com/watch?v=yyUHQIec83I]. The goal of this repo is to purely learn go. 

## Context 

Created by Google 2009. 

### Why?

Leverage changes in infrastructure improvements (paralell tasking) --> Multithreading. 
Booking systems (two users trying to book ticktes at the same time --> Can't happen) Prevent double booking --> concurrency 

Go was desgned to code with Multithreading and handeling concurrency easeier so the current updates in infra could be used. 

### Best Use case

* Performant applications 
* Run in scaled, distributed systems

### Characteristics

Go = python's syntax symplicity + C++ efficacy
Used in the server side: Microservices or web apps or db services
Simplicity and Speed --> popular in DevOps and SRE

K8s and Docker are written down on GO

## Install

Go to the official site (https://go.dev/dl/)

Add go path to ~/.zshrc:

```
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
```