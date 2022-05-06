# ddd-boilerplate-go
A Golang web service boilerplate follows Domain Driven Design  and Clean Architecture that includes BDD and Table Driven Testing

Heavily inspired by https://github.com/resotto/goilerplate

## Folder Structure
```sh
.
├── cmd
│   └── app  # main entry
├── internal # internal files
│   |── app
│   |   ├── adapter # adapter layer
│   |   │   ├── ctrl # controller impl
│   |   │   ├── repository # repository impl
│   |   │   │   └── redis # redis repository impl
│   |   │   └── service # service impl
│   |   ├── application # application layer
│   |   │   ├── service  # service interfaces
│   |   │   └── usecase  # core biz logics
│   |   └── domain  # domain entities
│   |       ├── repository  # repository interfaces
│   |       └── vo  # value object - domain entity segments
│   └── cfg
├── testdata  # testing setups and mockups
.
├── cmd
│   └── app # main entry
├── internal # internal files
│   ├── app # application
│   │   ├── adapter # adapter layer
│   │   │   ├── ctrl # controller impl
│   │   │   │   ├── index # index controller impl
│   │   │   │   └── session # session controller impl
│   │   │   ├── repository # repository impl
│   │   │   └── service # service impl
│   │   ├── application # application layer
│   │   │   ├── service # service interface
│   │   │   └── usecase # core biz logic
│   │   ├── domain # domain layer
│   │   │   ├── repository # repository interface
│   │   │   └── vo # value objects - domain entity segments
│   │   ├── infrastructure # infrastructure layer
│   │   |   ├── mysql # mysql
│   │   |   ├── rabbitmq # rabbitmq
│   │   |   └── redis # redis
│   │   └── server.go # http server
│   └── cfg # configs
└─── testdata # testing setups and mockups
```

## Quick Start

```sh
> make install
> make watch
> open localhost:8080/healthz
```

## Run Unit Test

```sh
> make test
```
