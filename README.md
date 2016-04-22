# Heartack
---
API allowing the doctor to recover cardiac readings of his patients to detect any anomaly !


## Pre-requisites
----
You need Golang installed:

* [Install Golang] [Golang]
* Configure your environments variables (see the advice section)

## Installing
----
Child's play

```sh
$ git clone git@git.tymate.com:jeremy/heartack.git
$ cd heartack
$ go get .
$ go install
$ make run
```
Then run your favorite browser and go to:
 [http://localhost:8080] [localhost]

#### Tech
---
* [Golang]
* [AngularJS]
* [Twitter Bootstrap]

#### Go Lib
---
 - Negroni [![GoDoc](https://godoc.org/github.com/codegangsta/negroni?status.svg)](http://godoc.org/github.com/codegangsta/negroni) 
 - Jwt-Go [![Build Status](https://travis-ci.org/dgrijalva/jwt-go.svg?branch=master)](https://travis-ci.org/dgrijalva/jwt-go)


#### Todos
---
 - Write Tests
 - Add DataBase
 - Add Code Comments
 - ...

#### Advice
---
> I recommend using [Autoenv] or [Dotenv] for 
> managing your environment variables
> if you do not.

Here is my .env file for example:
```sh
export GOPATH=/Users/jeremy/go/heartack
export PATH=$PATH:$GOROOT/bin:$GOBIN
export GOBIN=$GOPATH/bin
```

##### Version
---
1.0

##### License
----

Tymate


   [Autoenv]: <https://github.com/kennethreitz/autoenv>
   [Dotenv]: <https://github.com/bkeepers/dotenv>
   [Golang]: <https://golang.org/doc/install>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [AngularJS]: <http://angularjs.org>
   [localhost]: <http://localhost:8080>


