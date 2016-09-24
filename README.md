# Homescreen
This is a service for your personal touchscreen at home. You can display and control all your home automatization functions. We will provide some interesting modules but feel free to contribute and submit your own plugins! The homescreen provides an easily usabled plugin interface.

## Requirements for compiling
* [Go](https://golang.org/) to compile the backend
* [Glide](https://golang.org/) to load backend dependencies
* [Bower](https://bower.io/) to load the frontend dependencies

## Compile it yourself
* go get this project
  + go get github.com/homescreenrocks/homescreen
* load dependencies
  + `glide install` the backend dependencies from main directory
  + `bower install` the frontend dependencies from core/frontend folder
* get the plugins you want
  +  go get github.com/homescreenrocks/homescreen-plugin-example
* compile and start the homescreen core
  + execute in the main directory: `go build core\backend\app.go && app`
* compile and start a plugin
  + cd to the plugin: `go build example\main.go && main http://localhost:3000`

> we should provide a more easy way of usage :)

## Use the binaries
* ...

## Frontend requirements
* any modern browser (Chrome, Edge, Firefox)
