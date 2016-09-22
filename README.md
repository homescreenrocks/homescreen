# Homescreen
This is a service for your personal touchscreen at home. You can display and control all your home automatization functions. We will provide some interesting modules but feel free to contribute and submit your own plugins! The homescreen provides an easily usabled plugin interface.

## Requirements for compiling
* [Go](https://golang.org/) to compile the backend
* [Bower](https://bower.io/) to load the frontend dependencies

## Setup
* go get this project and the dependencies
  + go get github.com/homescreenrocks/homescreen
  + ...
  + bower install the frontend dependencies
* get the plugins you want
  +  go get github.com/homescreenrocks/homescreen-plugin-example
* cd to homescreen backend folder
  + compile and start the backend
* cd to a plugin
  + compile and start the plugin

> we should provide a more easy way of usage :)

## Plugin development
* find the API requirements for cumstom plugins in our [docs](/docs)

