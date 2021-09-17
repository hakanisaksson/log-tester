# log-tester

This is a simple hellow-world type application that
periodically prints put messages on stdout,
it is ment to run as a pod in a kubernetes cluster
to verify that the logs are sent correctly to a logserver (for example ELK)

It also has a custom metric available on /metric for verifying prometheus metrics also work

## How to build and test locally with docker and go and kubectl

You can build and test this application locally with the Makefile.

### How to setup a go dev environment
download and install golang in /usr/local/go
```
export GOROOT=/usr/local/go
export GOPATH=~/go
export PATH=$GOROOT/bin/:$PATH:$GOPATH/bin
```

### How to build locally
make go_build
make go_run

### How to build and test local docker container
make docker_build
make docker_test

### How to manually deploy to kubernetes
```
. ./env.sh
make build
make push
make deploy
```

### Metrics
This application also has some promethus metrics that should be avaiable on path /metrics:

Example:
```
make run
# curl http://172.17.0.2:8080/metrics
# HELP log_tester_messages_sent_counter_total
# TYPE log_tester_messages_sent_counter_total counter
log_tester_messages_sent_counter_total 3
make clean
```
Currently is only has a simple counter that measures the total number of messages sent since start
