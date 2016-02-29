Marcel: One docker CLI to rule them all

# Installation

    go get -u github.com/dgageot/marcel

# Change current docker host

## Point to default local daemon

    marcel use local

## Point to any docker machine host

    marcel use default
    marcel use another-docker-machine

## Point to any docker host with out without TLS

    marcel use tcp://50.134.234.20:2376 ~/certs
    marcel use tcp://50.134.234.20:2376

# Run any docker, compose or machine

    marcel run hello-world
    marcel compose up
    marcel machine ls
