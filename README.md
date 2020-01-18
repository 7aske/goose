# goose

## Description

Custom web-server tooled to easily deploy other web based projects (apps) using Git, inspired by Heroku. Deployment-server is designed to be a lightweight and simple to use.
Rewrite of `deployment-server-go` project

## Installation

Project is being managed by `make`. In order to install dependencies (NOTE: GOROOT must be properly set up but the project doesn't have to be in it) run:

```
make dep
```

After installing dependencies build the server using:

```
make build
```

Or just do this if you want a dynamic and faster compilation (and interactive mode):

```
make run
```

Install to /usr/bin

```
make install
```
