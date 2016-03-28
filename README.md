**Table of content**

1. About

2. Development

	2.1. Preparing the environment

	2.2. Running
	


# 1. About

My clone of Mou.app written with Go & QML.

No [Moa](https://en.wikipedia.org/wiki/Moa) has suffered during the development.

# 2. Development

## 2.1. Preparing the environment

In order to compile the program you need Go (1.5.2) and qt5. In addition you'll need these Go modules: 

- go-qml ([http://gopkg.in/qml.v1](http://gopkg.in/qml.v1)) 
- blackfriday ([https://github.com/russross/blackfriday](https://github.com/russross/blackfriday))

To get environement ready on OSX you can simply type the following commands:

1. Install Golang

	```
	brew install go
	```

* Install qt5

	```
	brew install qt5 pkg-config
	```

* Install go-qml

	```
	go get gopkg.in/qml.v1
	```

* Install blackfriday - markdown compiler

	```
	go get https://github.com/russross/blackfriday
	```
	
## 2.2 Running

To run the app simply type:

```
go run main.go
```